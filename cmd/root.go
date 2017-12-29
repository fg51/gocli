package cmd

import "fmt"

import "github.com/spf13/cobra"

type ExitStatus int

const (
	ExitSuccess ExitStatus = iota
	ExitFailure
	ExitUsageError
)

var RootCmd = &cobra.Command{
	Use:   "sample",
	Short: "Root short help",
	Long:  "Root long help",
}

func Execute() ExitStatus {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		return ExitFailure
	}
	return ExitSuccess
}

func init() {
}
