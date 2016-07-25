package refresh

import (
	"fmt"
	"os/exec"

	"github.com/markbates/going/clam"
)

func (m *Manager) runner() {
	var cmd *exec.Cmd
	for {
		<-m.Restart
		fmt.Println("restart!!!")
		if cmd != nil {
			// kill the preview command
			pid := cmd.Process.Pid
			m.Logger.Printf("=== Killing PID %d ===\n", pid)
			cmd.Process.Kill()
		}
		cmd = exec.Command(m.FullBuildPath(), m.CommandFlags...)
		go func() {
			err := clam.RunAndListen(cmd, func(s string) {
				m.Logger.Println(s)
			})
			m.Logger.Println(err)
		}()
	}
}

// func run(id string) bool {
// 	runnerLog("=== Running (%s) ===", id)
//
// 	cmd := exec.Command(buildPath(), cmdFlags())
//
// 	stderr, err := cmd.StderrPipe()
// 	if err != nil {
// 		fatal(err)
// 	}
//
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		fatal(err)
// 	}
//
// 	err = cmd.Start()
// 	if err != nil {
// 		fatal(err)
// 	}
//
// 	go io.Copy(appLogWriter{}, stderr)
// 	go io.Copy(appLogWriter{}, stdout)
//
// 	go func() {
// 		<-stopChannel
// 		pid := cmd.Process.Pid
// 		runnerLog("=== Killing PID %d ===", pid)
// 		cmd.Process.Kill()
// 	}()
//
// 	return true
// }
