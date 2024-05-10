package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	version := os.Getenv("INPUT_VERSION")
	arch := os.Getenv("INPUT_ARCH")
	installDir := os.Getenv("INPUT_INSTALL_DIR")

	if version == "" {
		version = "latest"
	}
	if arch == "" {
		arch = "x86_64"
	}
	if installDir == "" {
		installDir = "/usr/local/bin"

	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("curl -o awscliv2.zip https://awscli.amazonaws.com/awscli-exe-linux-%s-%s.zip && unzip awscliv2.zip && sudo ./aws/install --update --bin-dir %s", arch, version, installDir))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("::set-output name=aws-cli-version::%s\n", version)
}
