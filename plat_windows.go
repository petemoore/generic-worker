package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func startup() error {
	fmt.Println("Detected Windows platform...")
	fmt.Println("Looking for existing task users...")
	out, err := exec.Command("wmic", "useraccount", "get", "name").Output()
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	for _, line := range strings.Split(string(out), "\r\n") {
		if strings.HasPrefix(line, "Task_") {
			removeOSUser(line)
		}
	}
	return nil
}

func removeOSUser(user string) {
	out, err := exec.Command("net", "user", user, "/delete").Output()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(string(out))
	out, err = exec.Command("wmic", "path", "win32_userprofile", "where", "name=\""+user+"\"", "delete").Output()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(string(out))
}

func (task *TaskRun) generateCommand() *exec.Cmd {
	// TODO: below is the *nix implementation, Windows needs to generate a .bat file
	// and run it using PsExec instead...
	cmd := exec.Command(task.Payload.Command[0], task.Payload.Command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	task.prepEnvVars(cmd)
	return cmd
}