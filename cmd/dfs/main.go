package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dfs",
		Short: "Experimental Distributed File System",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Hello World")
		},
		Version: "0.1.0",
	}
	cmd.AddCommand(newPutCommand())

	return cmd
}

func main() {
	cmd := newRootCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
