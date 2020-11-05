package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aquasecurity/bench-common/check"
	"github.com/aquasecurity/bench-common/util"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func app(cmd *cobra.Command, args []string) {
	var version string
	var err error

	if linuxCisVersion != "" {
		version = linuxCisVersion
	} else {
		version = "2.0.0"
	}

	path, err := getDefinitionFilePath(version)
	if err != nil {
		util.ExitWithError(err)
	}

	constraints, err := getConstraints()
	if err != nil {
		util.ExitWithError(err)
	}

	controls, err := getControls(path, constraints)
	if err != nil {
		util.ExitWithError(err)
	}

	summary := runControls(controls, checkList)
	err = outputResults(controls, summary)
	if err != nil {
		util.ExitWithError(err)
	}
}

func outputResults(controls *check.Controls, summary check.Summary) error {
	// if we successfully ran some tests and it's json format, ignore the warnings
	if (summary.Fail > 0 || summary.Warn > 0 || summary.Pass > 0) && jsonFmt {
		out, err := controls.JSON()
		if err != nil {
			return err
		}
		util.PrintOutput(string(out), outputFile)
	} else {
		util.PrettyPrint(controls, summary, noRemediations, includeTestOutput)
	}

	return nil
}

func runControls(controls *check.Controls, checkList string) check.Summary {
	var summary check.Summary

	if checkList != "" {
		ids := util.CleanIDs(checkList)
		summary = controls.RunChecks(ids...)
	} else {
		summary = controls.RunGroup()
	}

	return summary
}

func getControls(path string, constraints []string) (*check.Controls, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	controls, err := check.NewControls([]byte(data), constraints)
	if err != nil {
		return nil, err
	}

	return controls, err
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

	glog.V(1).Info(fmt.Sprintf("The constraints are:, %s", constraints))
	return constraints, nil
}
