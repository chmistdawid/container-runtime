package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRunCmd() *cobra.Command {
	var cpu string
	var ram string
	var env string
	var volume string
	var network_bridge string
	runCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "run",
		Short: "Run a container",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run")
			fmt.Println(cpu)
			fmt.Println(ram)
			fmt.Println(env)
			fmt.Println(volume)
			fmt.Println(network_bridge)
		},
	}
	runCmd.PersistentFlags().StringVarP(&cpu, "cpu", "c", "1", "CPUs")
	runCmd.PersistentFlags().StringVarP(&ram, "ram", "r", "1", "RAM")
	runCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "Environment variables")
	runCmd.PersistentFlags().StringVarP(&volume, "volume", "v", "", "Volumes")
	runCmd.PersistentFlags().StringVarP(&network_bridge, "network-bridge", "n", "", "Network bridge")
	return runCmd
}

func NewStartCmd() *cobra.Command {
	startCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "start",
		Short: "Start a container",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("start")
		},
	}
	return startCmd
}

func NewStopCmd() *cobra.Command {
	stopCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "stop",
		Short: "Stop a container",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("stop")
		},
	}
	return stopCmd
}

func NewRestartCmd() *cobra.Command {
	restartCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "restart",
		Short: "Restart a container",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("restart")
		},
	}
	return restartCmd
}

func NewVersionCmd() *cobra.Command {
	var version = "0.0.1"
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of container-runtime",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("container-runtime version %s\n", version)
		},
	}
	return versionCmd
}
