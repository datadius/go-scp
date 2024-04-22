package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
	"main/scp"
)

func main() {
	clientConfig, err := auth.PasswordKey("claud", "", ssh.InsecureIgnoreHostKey())
	if err != nil {
		fmt.Print("Failed to create config: ", err)
		return
	}

	client := scp.NewClient("localhost:22", &clientConfig)

	err = client.Connect()
	if err != nil {
		fmt.Print("Failed to connect: ", err)
		return
	}

	f, err := os.Create("test.mp4")
	if err != nil {
		fmt.Print("Failed to open: ", err)
		return
	}

	defer client.Close()

	defer f.Close()

	// if the connection requires a PTY, then it will not work
	//err = client.CopyFromRemoteProgressPassThru(context.Background(), f, "hello.txt", nil)

	err = client.CopyFromRemoteProgressPassThru(context.Background(), f, "test.mp4", nil)

	if err != nil {
		fmt.Print("Failed to copy: ", err)
		return
	}
}
