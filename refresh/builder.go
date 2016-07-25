package refresh

import (
	"os/exec"

	"github.com/markbates/going/randx"
)

type Builder struct {
	Manager
	ID string
}

func NewBuilder(r Manager) Builder {
	return Builder{
		Manager: r,
		ID:      randx.String(10),
	}
}

func (b Builder) Build() error {
	cmd := exec.Command("go", "build", "-v", "-i", "-o", b.FullBuildPath())
	err := b.runAndListen(cmd, func(s string) {
		b.Logger.Print(s)
	})

	if err != nil {
		b.Logger.Error("Building Error!")
		b.Logger.Error(err)
		return err
	}
	b.Logger.Success("Building Completed (PID: %d)", cmd.Process.Pid)
	return nil
}
