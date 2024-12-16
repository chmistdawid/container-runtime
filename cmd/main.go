package main

import (
	"github.com/chmistdawid/container-runtime/cli"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{Use: "app"}

	stopCMD := cli.NewStopCmd()
	runCmd := cli.NewRunCmd()
	startCmd := cli.NewStartCmd()
	restartCmd := cli.NewRestartCmd()
	versionCmd := cli.NewVersionCmd()

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(stopCMD)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(restartCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}
