package main

import (
	"bufio"
	"log"
	"os/exec"
)

func runBuildCmd() error {
	cmd := exec.Command("flutter", "build", "web")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		return err
	}
	// read command's stdout line by line
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		log.Printf(in.Text()) // write each line to your log, or anything you need
	}
	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}
	cmd.Wait()
	return nil
}
