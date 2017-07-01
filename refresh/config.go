package refresh

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	AppRoot            string        `yaml:"app_root"`
	IgnoredFolders     []string      `yaml:"ignored_folders"`
	IncludedExtensions []string      `yaml:"included_extensions"`
	BuildTargetPath    string        `yaml:"build_target_path"`
	BuildPath          string        `yaml:"build_path"`
	BuildDelay         time.Duration `yaml:"build_delay"`
	BuildFlags         []string      `yaml:"build_flags,flow"`
	BinaryName         string        `yaml:"binary_name"`
	CommandFlags       []string      `yaml:"command_flags"`
	CommandEnv         []string      `yaml:"command_env"`
	EnableColors       bool          `yaml:"enable_colors"`
	LogName            string        `yaml:"log_name"`
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
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}
	// when build_flags is empty set to [-v, -i]
	// for compatibility with previous versions
	if len(c.BuildFlags) == 0 {
		c.BuildFlags = []string{"-v", "-i"}
	}
	return nil
}

func (c *Configuration) Dump(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0666)
}

func ID() string {
	d, _ := os.Getwd()
	return fmt.Sprintf("%x", md5.Sum([]byte(d)))
}
