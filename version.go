package amugine

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

var ver string
var rev string

type version struct {
}

func (v *version) Name() string {
	return "version"
}

func (v *version) Synopsis() string {
	return "Show the version of this tool."
}

func (v *version) Usage() string {
	return `version:
	Show the version of this tool.
`
}

func (v *version) SetFlags(f *flag.FlagSet) {
}

func (v *version) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	versionJSON, _ := json.Marshal(map[string]string{
		"version":  ver,
		"revision": rev,
	})
	fmt.Printf("%s\n", versionJSON)

	return subcommands.ExitSuccess
}
