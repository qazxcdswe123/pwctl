/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/qazxcdswe123/pwctl/pkgs/util"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

var (
	username string
	password string

	hostname string
)

func NewRootCommand() *cobra.Command {
	var (
		l = slog.Default()
	)

	cmd := &cobra.Command{
		Use:     "pwctl",
		Short:   "pwctl is a CLI tool to manage pgwatch3",
		Version: util.GetVersion().String(),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Validate hostname early
			if err := util.ValidateHostname(hostname); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Set the JWT Token before running any subcommand
			err := util.SetJWTToken(username, password, hostname)
			if err != nil {
				return fmt.Errorf("failed to set JWT Token: %v", err)
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&hostname, "hostname", "u", "http://localhost:8080", "URL of the server")
	cmd.PersistentFlags().StringVarP(&username, "username", "U", "", "Username")
	cmd.PersistentFlags().StringVarP(&password, "password", "P", "", "Password")

	cmd.AddCommand(newListCommand(l))
	cmd.AddCommand(newGetCommand(l))
	//cmd.AddCommand(newVersionCommand(l))

	return cmd
}
