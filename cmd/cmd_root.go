package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	RunE:         rootRunE,
	Short:        fmt.Sprintf("%s is a wireguard GRPC API", appname),
	SilenceUsage: true,
}

func rootRunE(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}
