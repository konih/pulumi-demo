package parser

import (
	"io/ioutil"
)

type Developer struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

type Namespace struct {
	Name     string            `yaml:"name"`
	Labels   map[string]string `yaml:"additional_labels"`
	Selector []map[string]map[string]string
}

type RancherProject struct {
	ProjectName string      `yaml:"project_name"`
	Namespaces  []Namespace `yaml:"namespaces"`
}

type Configuration struct {
	Key      string `yaml:"key"`
	Metadata struct {
		EAMID      string      `yaml:"eam_id"`
		Developers []Developer `yaml:"developers"`
	} `yaml:"metadata"`
	MailContact     []string         `yaml:"mail_contact"`
	RancherProjects []RancherProject `yaml:"rancher_projects"`
}

func ParseYAML(filePath string) (Configuration, error) {
	yamlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Configuration{}, err
	}

	var config Configuration
	err = yaml.Unmarshal(yamlContent, &config)
	if err != nil {
		return Configuration{}, err
	}

	return config, nil
}
