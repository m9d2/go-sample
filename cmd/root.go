package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "sample",
		Short: "A sample Webapp",
		Long:  "A sample Webapp.",
	}
)

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port to listen on")
}
