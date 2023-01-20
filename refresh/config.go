package refresh

import (
	"crypto/md5"
	"fmt"
	"io"

	"os"
	"path"
	"runtime"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	AppRoot            string        `yaml:"app_root"`
	BinaryName         string        `yaml:"binary_name"`
	BuildDelay         time.Duration `yaml:"build_delay"`
	BuildFlags         []string      `yaml:"build_flags"`
	BuildPath          string        `yaml:"build_path"`
	BuildTargetPath    string        `yaml:"build_target_path"`
	CommandEnv         []string      `yaml:"command_env"`
	CommandFlags       []string      `yaml:"command_flags"`
	EnableColors       bool          `yaml:"enable_colors"`
	ForcePolling       bool          `yaml:"force_polling,omitempty"`
	IgnoredFolders     []string      `yaml:"ignored_folders"`
	IncludedExtensions []string      `yaml:"included_extensions"`
	LogName            string        `yaml:"log_name"`
	EnableLivereload   bool          `yaml:"enable_livereload"`
	Debug              bool          `yaml:"-"`
	Path               string        `yaml:"-"`
	Stderr             io.Writer     `yaml:"-"`
	Stdin              io.Reader     `yaml:"-"`
	Stdout             io.Writer     `yaml:"-"`
}

func (c *Configuration) FullBuildPath() string {
	buildPath := path.Join(c.BuildPath, c.BinaryName)
	if runtime.GOOS == "windows" {
		if !strings.HasSuffix(strings.ToLower(buildPath), ".exe") {
			buildPath += ".exe"
		}
	}
	return buildPath
}

func (c *Configuration) Load(path string) error {
	// "io/ioutil" has been deprecated since Go 1.16: As of Go 1.16, the same functionality is now provided by package io or package os, and those implementations should be preferred in new code. See the specific function documentation for details.
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	c.Path = path
	return yaml.Unmarshal(data, c)
}

func (c *Configuration) Dump(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	// "io/ioutil" has been deprecated since Go 1.16: As of Go 1.16, the same functionality is now provided by package io or package os, and those implementations should be preferred in new code. See the specific function documentation for details.
	return nil
}

func ID() string {
	d, _ := os.Getwd()
	return fmt.Sprintf("%x", md5.Sum([]byte(d)))
}
