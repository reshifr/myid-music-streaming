package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

type Cli struct {
	binMemo map[string]string
}

func (cli *Cli) Exec(name string, arg ...string) (string, int) {
	bin, err := cli.locateName(name)
	if err != nil {
		return "", FAILURE
	}
	os.Clearenv()
	cmd := exec.Command(bin, arg...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(bin)
		log.Fatalf("Exec: can not execute '%v'.", cmd.String())
		return "", FAILURE
	}
	cleanOutput := getCleanOutput(output)
	return cleanOutput, SUCCESS
}

func (cli *Cli) locateName(name string) (string, error) {
	// Looking up command path
	bin, exists := cli.binMemo[name]
	if exists {
		return bin, nil
	}
	// Find command path
	cmd := exec.Command("which", name)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Exec: can not locate command path.")
		return "", err
	}
	bin = getCleanOutput(output)
	if cli.binMemo == nil {
		cli.binMemo = make(map[string]string)
	}
	cli.binMemo[name] = bin
	return bin, nil
}

func getCleanOutput(output []byte) string {
	outputLength := len(output)
	if outputLength == 0 {
		return ""
	}
	return string(output[:outputLength-1])
}
