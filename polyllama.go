package main

import (
    "os"
    "github.com/urfave/cli/v2"
)

func main() {
    app := cli.App{}
    app.Run(os.Args)
}
