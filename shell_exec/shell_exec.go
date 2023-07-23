package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd", "/c", "C:\\Program Files\\Smart Storage Administrator\\ssacli\\bin\\ssacli.exe", "help", "login")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	output, _ := io.ReadAll(stdout)
	fmt.Printf("%s\n", output)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

}
