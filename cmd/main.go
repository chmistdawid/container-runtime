package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var version = "0.0.1"

	var rootCmd = &cobra.Command{Use: "app"}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of container-runtime",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("container-runtime version %s\n", version)
		},
	}
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run a container",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("run container\n")
		},
	}
	stopCMD := &cobra.Command{
		Use:   "stop",
		Short: "Stop a container",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("stop container\n")
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(stopCMD)
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}
