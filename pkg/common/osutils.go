package common

import (
	"io"
	"os"
	"os/exec"
	"sync"
)

func Exec(command string, arg ...string) (string, error) {
	out, err := exec.Command(command, arg...).CombinedOutput()
	outStr := string(out)
	return outStr, err
}

// Function to execute local command
func ExecShell(command string) (string, string, error) {
	return ExecShellWithEnv(command)
}

func ExecShellWithEnv(command string, envVars ...string) (string, string, error) {
	var stout, sterr []byte
	cmd := exec.Command("bash", "-c", command)
	log.Debug("Command %s ", command)
	cmd.Env = append(cmd.Env, envVars...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Debug("Command %s failed to start. Cause: %v", command, err)
		return "", "", err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stout, _ = copyAndCapture(os.Stdout, stdout)
		wg.Done()
	}()

	sterr, _ = copyAndCapture(os.Stderr, stderr)

	wg.Wait()

	err := cmd.Wait()
	return string(stout), string(sterr), err
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}
