package cra

import (
	"fmt"
	"log"
	"os/exec"
)

func runCMD(dir string, cmdStrings []string) error {
	args := cmdStrings[1:]
	cmd := exec.Command(cmdStrings[0], args...)
	cmd.Dir = dir

	log.Println(cmd.String())
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error executing command: %e", err)
	}

	if err := cmd.Wait(); err != nil {
		exiterr, ok := err.(*exec.ExitError)
		if !ok {
			return fmt.Errorf("error casting to ExitError: %e", err)
		}

		if exiterr.ExitCode() != 0 {
			return fmt.Errorf("process exited with a non 0 code: %e", err)
		}
	}

	return nil
}
