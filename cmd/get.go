/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
)

func newGetCommand(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			l.Error("Get command requires a subcommand")
			return cmd.Help()
		},
	}

	cmd.AddCommand(newGetDatabaseCommand(l))

	return cmd
}

func newGetDatabaseCommand(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database [name]",
		Short: "Get a monitored databases by name",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//dbName := args[0]
			//db := resource.NewDatabase()

			return nil
		},
	}

	return cmd
}
