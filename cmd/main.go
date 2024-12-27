package main

import (
	"context"
	"log"

	"github.com/chmistdawid/container-runtime/cli"
	"github.com/chmistdawid/container-runtime/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//imagesDir := "./images"
	conn, err := grpc.NewClient("unix:///var/run/feather.sock", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer conn.Close()

	c := proto.NewFeatherClient(conn)
	ctx := context.Background()

	var rootCmd = &cobra.Command{Use: "app"}

	stopCMD := cli.NewStopCmd(c, ctx)
	runCmd := cli.NewRunCmd(c, ctx)
	startCmd := cli.NewStartCmd(c, ctx)
	restartCmd := cli.NewRestartCmd(c, ctx)
	versionCmd := cli.NewVersionCmd(c, ctx)
	statusCmd := cli.NewStatusCmd(c, ctx)
	deleteCmd := cli.NewDeleteCmd(c, ctx)

	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCMD)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}
