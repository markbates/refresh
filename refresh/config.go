package refresh

import (
	"io/ioutil"
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
