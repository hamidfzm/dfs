package main

import (
	"github.com/spf13/cobra"

	"github.com/hamidfzm/dfs/pkg/file"
)

func newPutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put [src...]",
		Short: "Put file in system",
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {

				if err := file.Split(arg, "temp", 1<<10); err != nil {
					return err
				}
			}
			return nil
		},
		Args: cobra.MinimumNArgs(1),
	}
	return cmd
}
