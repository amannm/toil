package cmd

import (
	"context"
	"os"
	"testing"
)

func TestBasic(t *testing.T) {
	os.Args = []string{"toil", "stop"}
	err := Execute(context.Background())
	if err != nil {
		t.Error(err)
	}
}
