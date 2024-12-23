package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/chmistdawid/container-runtime/cli"
	"github.com/spf13/cobra"
)

func main() {

	dbName := "file:./local.db"
	//imagesDir := "./images"

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Printf("failed to open db %s", err)
		os.Exit(1)
	}
	defer db.Close()

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
