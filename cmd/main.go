package main

import (
	"github.com/chmistdawid/container-runtime/cli"
	"github.com/spf13/cobra"
)

func main() {

	//imagesDir := "./images"

	var rootCmd = &cobra.Command{Use: "app"}

	stopCMD := cli.NewStopCmd()
	runCmd := cli.NewRunCmd()
	startCmd := cli.NewStartCmd()
	restartCmd := cli.NewRestartCmd()
	versionCmd := cli.NewVersionCmd()
	statusCmd := cli.NewStatusCmd()
	deleteCmd := cli.NewDeleteCmd()

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(stopCMD)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}
