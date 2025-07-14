package main

import (
	"os"
	"testing"
)

// Tests all standard linux-bench defintion files
func TestGetDefinitionFilePath(t *testing.T) {
	d, err := os.Open("./cfg")
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	vers, err := d.Readdir(-1)
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	for _, ver := range vers {
		if !ver.IsDir() {
			t.Logf("Skipping non-directory: %v", ver.Name())
			continue
		}
		t.Logf("%v", ver)
		_, err := getDefinitionFilePath(ver.Name())
		if err != nil {
			t.Errorf("unexpected error: %s\n", err)
		}
	}
}

func TestRunControls(t *testing.T) {
	cfgDir = "./hack"
	path, err := getDefinitionFilePath("test-definitions")
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	control, err := getControls(path, nil)
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	// Run all checks
	_ = runControls(control, "")

	// Run only specified checks
	checkList := "1.2, 2.1"
	_ = runControls(control, checkList)
}
