package main

import (
	"path/filepath"
	"strconv"
	"strings"
)

func helloGoodbye() []string {
	return []string{
		"echo hello world!",
		"echo goodbye world!",
	}
}

func rawHelloGoodbye() string {
	return `"command": [
    "echo hello world!",
    "echo goodbye world!"
  ]`
}

func checkSHASums() []string {
	return []string{
		"PowerShell.exe -NoProfile -ExecutionPolicy Bypass -File preloaded\\check-shasums.ps1",
	}
}

func failCommand() []string {
	return []string{
		"exit 1",
	}
}

func incrementCounterInCache() []string {
	// The `echo | set /p dummyName...` construction is to avoid printing a
	// newline. See answer by xmechanix on:
	// http://stackoverflow.com/questions/7105433/windows-batch-echo-without-new-line/19468559#19468559
	command := `
		setlocal EnableDelayedExpansion
		if exist my-task-caches\test-modifications\counter (
		  set /p counter=<my-task-caches\test-modifications\counter
		  set /a counter=counter+1
		  echo | set /p dummyName="!counter!" > my-task-caches\test-modifications\counter
		) else (
		  echo | set /p dummyName="1" > my-task-caches\test-modifications\counter
		)
`
	return []string{command}
}

func sleep(seconds uint) []string {
	return []string{
		"ping 127.0.0.1 -n " + strconv.Itoa(int(seconds+1)) + " > nul",
	}
}

func goRun(goFile string) []string {
	return []string{
		"go run \"" + goFile + "\"",
	}
}

func copyArtifact(path string) []string {
	targetFile := strings.Replace(path, "/", "\\", -1)
	sourceFile := filepath.Join(testdataDir, targetFile)
	return []string{
		"mkdir \"" + filepath.Dir(targetFile) + "\"",
		"copy \"" + sourceFile + "\" \"" + targetFile + "\"",
	}
}
