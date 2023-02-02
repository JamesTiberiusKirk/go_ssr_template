package cra

import (
	"fmt"
	"os/exec"
)

func runCMD(dir string, cmdStrings []string, verbose bool) error {
	args := cmdStrings[1:]
	cmd := exec.Command(cmdStrings[0], args...)
	cmd.Dir = dir

	fmt.Println(cmd.String())
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error executing command: %e", err)
	}

	if verbose {
		stdout, err := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		if err != nil {
			return err
		}

		for {
			tmp := make([]byte, 1024)
			_, err := stdout.Read(tmp)
			fmt.Print(string(tmp))
			if err != nil {
				break
			}
		}
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
