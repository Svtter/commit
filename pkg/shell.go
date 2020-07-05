package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Shellout(command string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func ReadFromCommand() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("commit msg: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	// commitArgs := fmt.Sprintf("commit -m \"%s\"", text)
	// fmt.Println(commitArgs)
	return text
}

func OutputError(out string, errout string, err error) {
	log.Printf("Command finished with error: %v", err)
	log.Printf("stdout: %v", out)
	log.Fatalf("stderr: %v", errout)
}

func Output(out, errout string, err error) {
	if err != nil {
		OutputError(out, errout, err)
	} else {
		fmt.Print(out)
	}
}