package main

import (
	"fmt"
	"os"

	"github.com/lateralusd/fcker/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "[-] Error occured: %+v\n", err)
	}
}
