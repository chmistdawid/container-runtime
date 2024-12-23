package cli

import (
	"log"

	"github.com/spf13/cobra"
)

const SOCK_PATH = "/var/run/feather.sock"

func NewRunCmd() *cobra.Command {
	var cpu string
	var ram string
	var env string
	var volume string
	var network_bridge string
	var image string
	var name string
	runCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "run",
		Short: "Run a container",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	runCmd.PersistentFlags().StringVarP(&cpu, "cpu", "c", "1", "CPUs")
	runCmd.PersistentFlags().StringVarP(&ram, "ram", "r", "1", "RAM")
	runCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "Environment variables")
	runCmd.PersistentFlags().StringVarP(&volume, "volume", "v", "", "Volumes")
	runCmd.PersistentFlags().StringVarP(&image, "image", "i", "", "Image")
	runCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	runCmd.PersistentFlags().StringVarP(&network_bridge, "network-bridge", "n", "", "Network bridge")
	return runCmd
}

func NewStartCmd() *cobra.Command {
	var name string
	startCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "start",
		Short: "Start a container",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("start")
		},
	}
	startCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return startCmd
}

func NewStopCmd() *cobra.Command {
	var name string
	stopCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "stop",
		Short: "Stop a container",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("stop")
		},
	}
	stopCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return stopCmd
}

func NewRestartCmd() *cobra.Command {
	var name string
	restartCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "restart",
		Short: "Restart a container",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("restart")
		},
	}
	restartCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return restartCmd
}

func NewStatusCmd() *cobra.Command {
	var name string
	statusCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "status",
		Short: "Status a container",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("status")
		},
	}
	statusCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return statusCmd
}

func NewDeleteCmd() *cobra.Command {
	var name string
	deleteCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "delete",
		Short: "Delete a container",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("delete")
		},
	}
	deleteCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return deleteCmd
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
