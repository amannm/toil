package main

import (
	"context"
	"github.com/amannm/toil/cmd"
	"os"
)

func main() {
	ctx := context.Background()
	err := cmd.Execute(ctx)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
