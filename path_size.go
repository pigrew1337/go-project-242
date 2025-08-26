package code

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func Cli() {
	(&cli.Command{}).Run(context.Background(), os.Args)
}
