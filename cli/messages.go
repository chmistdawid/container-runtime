package cli

import (
	"fmt"
	"net"
	"os"
)

func sendStartMessage(cpu string, ram string, env string, volume string, network_bridge string, image string, name string) {
	conn, err := net.Dial("unix", SOCK_PATH)
	if err != nil {
		fmt.Printf("Error connecting to socket: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := fmt.Sprintf("start %s %s %s %s %s %s %s", cpu, ram, env, volume, network_bridge, image, name)

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error writing to socket: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully")
}

func sendStopMessage(name string) {
	conn, err := net.Dial("unix", SOCK_PATH)
	if err != nil {
		fmt.Printf("Error connecting to socket: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := fmt.Sprintf("stop %s", name)

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error writing to socket: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully")
}

func sendRestartMessage(name string) {
	conn, err := net.Dial("unix", SOCK_PATH)
	if err != nil {
		fmt.Printf("Error connecting to socket: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := fmt.Sprintf("restart %s", name)

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error writing to socket: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully")
}

func sendStatusMessage(name string) {
	conn, err := net.Dial("unix", SOCK_PATH)
	if err != nil {
		fmt.Printf("Error connecting to socket: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := fmt.Sprintf("status %s", name)

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error writing to socket: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully")
}

func sendDeleteMessage(name string) {
	conn, err := net.Dial("unix", SOCK_PATH)
	if err != nil {
		fmt.Printf("Error connecting to socket: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := fmt.Sprintf("delete %s", name)

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error writing to socket: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully")
}
