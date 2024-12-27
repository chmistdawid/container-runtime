package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/chmistdawid/container-runtime/proto"
	"github.com/spf13/cobra"
)

const SOCK_PATH = "/var/run/feather.sock"

func NewRunCmd(c proto.FeatherClient, ctx context.Context) *cobra.Command {
	var cpu string
	var ram string
	var env string
	var volume string
	var network_bridge string
	var image string
	var name string

	runCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(0),
		Use:   "run",
		Short: "Run a container",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := c.Run(ctx, &proto.RunRequest{
				Cpu:           cpu,
				Ram:           ram,
				Env:           env,
				Volume:        volume,
				NetworkBridge: network_bridge,
				Image:         image,
				Name:          name,
			})
			if err != nil {
				log.Fatalf("failed to run container: %v", err)
			}
			fmt.Println(resp.Message)
		},
	}
	runCmd.PersistentFlags().StringVarP(&cpu, "cpu", "c", "", "CPUs")
	runCmd.PersistentFlags().StringVarP(&ram, "ram", "r", "", "RAM")
	runCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "Environment variables")
	runCmd.PersistentFlags().StringVarP(&volume, "volume", "v", "", "Volumes")
	runCmd.PersistentFlags().StringVarP(&image, "image", "i", "", "Image")
	runCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	runCmd.PersistentFlags().StringVarP(&network_bridge, "network-bridge", "b", "", "Network bridge")
	return runCmd
}

func NewStartCmd(c proto.FeatherClient, ctx context.Context) *cobra.Command {
	var name string
	startCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "start",
		Short: "Start a container",
		Run: func(cmd *cobra.Command, args []string) {
			sendStartMessage(name)
		},
	}
	startCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return startCmd
}

func NewStopCmd(c proto.FeatherClient, ctx context.Context) *cobra.Command {
	var name string
	stopCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "stop",
		Short: "Stop a container",
		Run: func(cmd *cobra.Command, args []string) {
			sendStopMessage(name)
		},
	}
	stopCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return stopCmd
}

func NewRestartCmd(c proto.FeatherClient, ctx context.Context) *cobra.Command {
	var name string
	restartCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "restart",
		Short: "Restart a container",
		Run: func(cmd *cobra.Command, args []string) {
			sendRestartMessage(name)
		},
	}
	restartCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return restartCmd
}

func NewStatusCmd(c proto.FeatherClient, ctx context.Context) *cobra.Command {
	var name string
	statusCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "status",
		Short: "Status a container",
		Run: func(cmd *cobra.Command, args []string) {
			sendStatusMessage(name)
		},
	}
	statusCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return statusCmd
}

func NewDeleteCmd(c proto.FeatherClient, ctx context.Context) *cobra.Command {
	var name string
	deleteCmd := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "delete",
		Short: "Delete a container",
		Run: func(cmd *cobra.Command, args []string) {
			sendDeleteMessage(name)
		},
	}
	deleteCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name")
	return deleteCmd
}

func NewVersionCmd(c proto.FeatherClient, ctx context.Context) *cobra.Command {
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
