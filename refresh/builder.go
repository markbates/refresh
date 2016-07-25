package refresh

import (
	"os/exec"

	"github.com/markbates/going/clam"
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
	b.Logger.Printf("=== Building (ID: %s) ===", b.ID)

	cmd := exec.Command("go", "build", "-v", "-i", "-o", b.FullBuildPath())
	err := clam.RunAndListen(cmd, func(s string) {
		b.Logger.Printf("\t[%s] %s\n", b.ID, s)
	})

	if err != nil {
		return err
	}
	b.Logger.Printf("=== Completed (ID: %s) (PATH: %s) ===\n", b.ID, b.FullBuildPath())
	return nil
}
