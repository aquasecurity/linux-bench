package main

import (
	"fmt"
	"github.com/aquasecurity/bench-common/util"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aquasecurity/bench-common/runner"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func app(cmd *cobra.Command, args []string) {
	var version string
	var err error

	if linuxCisVersion != "" {
		version = linuxCisVersion
	} else {
		version = "1.1.0"
	}

	path, err := getDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	constraints, err := getConstraints()
	if err != nil {
		util.ExitWithError(err)
	}

	yamlCfg, err := ioutil.ReadFile(path)
	if err != nil {
		util.ExitWithError(err)
	}
	benchRunner, err := runner.New(yamlCfg).
		WithConstrains(constraints).
		WithCheckList(checkList).Build()

	if err != nil {
		util.ExitWithError(err)
	}

	err = benchRunner.RunTestsWithOutput(jsonFmt, noRemediations, includeTestOutput)
	if err != nil {
		util.ExitWithError(err)
	}
}

func getDefinitionFilePath(version string) (string, error) {
	filename := "definitions.yaml"

	glog.V(2).Info(fmt.Sprintf("Looking for config for version %s", version))

	path := filepath.Join(cfgDir, version)
	file := filepath.Join(path, filename)

	glog.V(2).Info(fmt.Sprintf("Looking for config file: %s\n", file))

	_, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	return file, nil
}

func getConstraints() (constraints []string, err error) {
	platform, err := GetOperatingSystem()
	if err != nil {
		glog.V(1).Info(fmt.Sprintf("Failed to get operating system platform, %s", err))
	}

	boot, err := GetBootLoader()
	if err != nil {
		glog.V(1).Info(fmt.Sprintf("Failed to get boot loader, %s", err))
	}

	syslog, err := GetSystemLogManager()
	if err != nil {
		glog.V(1).Info(fmt.Sprintf("Failed to get syslog tool, %s", err))

	}

	lsm, err := GetLSM()
	if err != nil {
		glog.V(1).Info(fmt.Sprintf("Failed to get lsm, %s", err))
	}

	constraints = append(constraints,
		fmt.Sprintf("platform=%s", platform),
		fmt.Sprintf("boot=%s", boot),
		fmt.Sprintf("syslog=%s", syslog),
		fmt.Sprintf("lsm=%s", lsm),
	)

	return constraints, nil
}
