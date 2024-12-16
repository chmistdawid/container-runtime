package main

import (
	"fmt"

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
			runContainer(args)
		},
	}
	stopCMD := &cobra.Command{
		Use:   "stop",
		Short: "Stop a container",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cmd.Printf("Usage: stop <container-id>")
				return
			}
			fmt.Println("stop container")
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(stopCMD)
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}

func runContainer(args []string) {
	// TODO: run container
	if len(args) != 1 {
		fmt.Printf("Usage: start <container-id>")
		return
	}
	fmt.Println("run container")
}
