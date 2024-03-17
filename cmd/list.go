package cmd

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/qazxcdswe123/pwctl/pkgs/resource"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"strconv"
)

func newListCommand(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List monitored database, metrics, or preset",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(newListDatabaseCommand(l))

	return cmd
}

func newListDatabaseCommand(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "List monitored databases",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			db := resource.NewDatabase()

			databases, err := db.List(ctx, &resource.ListOptions{
				Hostname: hostname,
			})
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Connstr"})

			for i, database := range databases {
				table.Append([]string{strconv.Itoa(i), database.Name, database.Connstr})
			}

			table.Render()
			return nil
		},
	}

	return cmd
}
