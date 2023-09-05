package parser

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseYAML(t *testing.T) {
	testYAML := `
key: cp
metadata:
  eam_id: PROD150
  developers:
    - id: a01331a
      name: Wurst, Hans
  mail_contact:
    - test@test.com
global_roles:
  - global_role_template_name: Admin
rancher_projects:
  - project_name: test
    namespaces:
      - name: testwurst
      - name: fleischwurst
`

	ioutil.WriteFile("temp_test.yaml", []byte(testYAML), 0644)
	defer os.Remove("temp_test.yaml")

	config, err := ParseYAML("temp_test.yaml")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if config.Key != "cp" {
		t.Errorf("Expected key 'cp', got %v", config.Key)
	}

	if config.Metadata.EAMID != "PROD150" {
		t.Errorf("Expected EAMID 'PROD150', got %v", config.Metadata.EAMID)
	}

	if len(config.Metadata.Developers) != 1 {
		t.Errorf("Expected 1 developer, got %v", len(config.Metadata.Developers))
	}
}
