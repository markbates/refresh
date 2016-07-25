package refresh

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func (m *Manager) runner() {
	var cmd *exec.Cmd
	for {
		<-m.Restart
		if cmd != nil {
			// kill the previous command
			pid := cmd.Process.Pid
			m.Logger.Success("Stopping: PID %d", pid)
			cmd.Process.Kill()
		}
		cmd = exec.Command(m.FullBuildPath(), m.CommandFlags...)
		go func() {
			err := m.runAndListen(cmd, func(s string) {
				fmt.Println(s)
			})
			m.Logger.Error(err)
		}()
	}
}

func (m *Manager) runAndListen(cmd *exec.Cmd, fn func(s string)) error {
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	r, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("%s\n%s", err, stderr.String())
	}

	scanner := bufio.NewScanner(r)
	go func() {
		for scanner.Scan() {
			fn(scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("%s\n%s", err, stderr.String())
	}

	m.Logger.Success("Running: %s (PID: %d)", strings.Join(cmd.Args, " "), cmd.Process.Pid)
	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("%s\n%s", err, stderr.String())
	}
	return nil
}
