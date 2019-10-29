package amugine

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/subcommands"
	"golang.org/x/crypto/openpgp"
)

type decrypter struct {
}

func (d *decrypter) Name() string {
	return "decrypt"
}

func (d *decrypter) Synopsis() string {
	return "Decrypt a given payload."
}

func (d *decrypter) Usage() string {
	return `decrypt <key> <payload>:
	Decrypt a given payload. If a parameter has "@" prefix, it will be handled as the filepath and read that.
`
}

func (d *decrypter) SetFlags(f *flag.FlagSet) {
}

func (d *decrypter) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	args := f.Args()
	if len(args) < 2 {
		fmt.Printf("lacked mandatory parameter(s)\n")
		return subcommands.ExitFailure
	}

	keyParamReader := &paramReader{args[0]}
	key, err := keyParamReader.GetValue()
	if err != nil {
		fmt.Printf("failed to read a key: %s\n", err)
		return subcommands.ExitFailure
	}

	payloadParamReader := &paramReader{args[1]}
	payload, err := payloadParamReader.GetValue()
	if err != nil {
		fmt.Printf("failed to read a payload: %s\n", err)
		return subcommands.ExitFailure
	}

	decrypted, err := decrypt(key, payload)
	if err != nil {
		fmt.Printf("failed to decrypt a payload: %s\n", err)
		return subcommands.ExitFailure
	}

	_, err = os.Stdout.Write(decrypted)
	if err != nil {
		fmt.Printf("failed to output an decrypted payload to stdout: %s\n", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func decrypt(key []byte, payload []byte) ([]byte, error) {
	prompt := func(passPhrase []byte) openpgp.PromptFunction {
		var called bool
		return func([]openpgp.Key, bool) ([]byte, error) {
			if called {
				return nil, errors.New("the passphrase is invalid")
			}
			called = true
			return passPhrase, nil
		}
	}

	msg, err := openpgp.ReadMessage(bytes.NewReader(payload), nil, prompt(key), encryptionConfiguration)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(msg.UnverifiedBody)
}
