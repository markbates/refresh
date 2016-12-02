package refresh

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	AppRoot            string        `yaml:"app_root"`
	IgnoredFolders     []string      `yaml:"ignored_folders"`
	IncludedExtensions []string      `yaml:"included_extensions"`
	BuildPath          string        `yaml:"build_path"`
	BuildDelay         time.Duration `yaml:"build_delay"`
	BinaryName         string        `yaml:"binary_name"`
	CommandFlags       []string      `yaml:"command_flags"`
	EnableColors       bool          `yaml:"enable_colors"`
	LogName            string        `yaml:"log_name"`
}

func (c *Configuration) FullBuildPath() string {
	return path.Join(c.BuildPath, c.BinaryName)
}

func (c *Configuration) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
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
