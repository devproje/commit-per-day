package main

import (
	"os"
	"os/exec"
	"time"
)

var (
	first = true
)

func runCode(name string, args ...string) {
	exec.Command(name, args...).Run()
}

func editFile() {
	var str string
	if !first {
		str += "\n"
	}
	str += "- "
	today := time.Now().Format("1/2/2006")

	str += today
	fs, err := os.ReadFile("README.md")
	if err != nil {
		return
	}

	converted := string(fs[:])
	converted += str

	convert := []byte(converted)

	os.WriteFile("README.md", convert, os.FileMode(775))
}

func push() {
	runCode("git", "add", ".")
	runCode("git", "commit", "-m", "\"Update README.md\"")
	runCode("git", "push", "origin", "master")
	time.Sleep(time.Second * 10)
}

func main() {
	for {
		if time.Now().Local().Hour() == 0 && time.Now().Local().Minute() == 0 && time.Now().Local().Second() == 0 {
			editFile()
			push()
			if first {
				first = false
			}
		}
	}
}
