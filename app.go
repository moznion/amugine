package amugine

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"golang.org/x/crypto/openpgp/packet"
)

var encryptionConfiguration = &packet.Config{
	DefaultCipher: packet.CipherAES256,
}

// Run is the entry point of this tool.
func Run(args []string) {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&encrypter{}, "")
	subcommands.Register(&decrypter{}, "")
	subcommands.Register(&version{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
