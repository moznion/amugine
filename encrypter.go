package amugine

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/subcommands"
	"golang.org/x/crypto/openpgp"
)

type encrypter struct {
}

func (e *encrypter) Name() string {
	return "encrypt"
}

func (e *encrypter) Synopsis() string {
	return "Encrypt s given payload."
}

func (e *encrypter) Usage() string {
	return `encrypt <key> <payload>:
	Encrypt a given payload. It supports to give payload through an argument and STDIN (example: "cat payload.txt | amugine encrypt <key>").
	If a parameter has "@" prefix, it will be handled as the filepath and read that.
`
}

func (e *encrypter) SetFlags(f *flag.FlagSet) {
}

func (e *encrypter) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	args := f.Args()
	if len(args) == 0 {
		fmt.Printf("lacked mandatory parameter(s)\n")
		return subcommands.ExitFailure
	}

	keyParamReader := &paramReader{args[0]}
	key, err := keyParamReader.GetValue()
	if err != nil {
		fmt.Printf("failed to read a key: %s\n", err)
		return subcommands.ExitFailure
	}

	var payload []byte
	if len(args) >= 2 {
		payloadParamReader := &paramReader{args[1]}
		payload, err = payloadParamReader.GetValue()
	} else {
		payload, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		fmt.Printf("failed to read a payload: %s\n", err)
		return subcommands.ExitFailure
	}

	encryptedBuff, err := encrypt(key, payload)
	if err != nil {
		fmt.Printf("failed to encrypt a payload: %s\n", err)
		return subcommands.ExitFailure
	}

	_, err = os.Stdout.Write(encryptedBuff.Bytes())
	if err != nil {
		fmt.Printf("failed to output an encrypted payload to stdout: %s\n", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func encrypt(key []byte, payload []byte) (*bytes.Buffer, error) {
	encryptedBuff := bytes.NewBuffer(nil)

	encrypter, err := openpgp.SymmetricallyEncrypt(encryptedBuff, key, nil, encryptionConfiguration)
	if err != nil {
		return nil, err
	}
	defer encrypter.Close()

	_, err = encrypter.Write(payload)
	if err != nil {
		return nil, err
	}

	return encryptedBuff, nil
}
