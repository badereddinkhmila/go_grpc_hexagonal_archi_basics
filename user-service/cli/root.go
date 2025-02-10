package cli

import (
	"user-service/cli/cmd"

	"github.com/spf13/cobra"
)

func Execute() error {
	root := &cobra.Command{
		Use:   "microservice",
		Short: "microservice",
		Long:  `Microservice cli to execute different commands related to staring the service`,
	}

	root.AddCommand(cmd.Serve())
	return root.Execute()
}
