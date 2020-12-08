package main

import (
	"os/exec"
	"regexp"
	"strings"
)

func GetOperatingSystem() (platform string, err error) {
	out, err := exec.Command("bash", "-c", "cat /etc/os-release").Output()

	if err != nil {
		return "", err
	} else {
		output := strings.ToLower(string(out))
		output = strings.Replace(output, `"`, "", -1)
		output = strings.Replace(output, `_id`, "", -1) // version_id kills the regex

		flagRe := regexp.MustCompile("id" + `=([^ \n]*)`)
		vals := flagRe.FindStringSubmatch(output)
		if len(vals) > 1 {
			platform = vals[1]
		}

		platform += getPlatformVersion(output, platform)
	}

	return platform, nil
}

func GetBootLoader() (boot string, err error) {
	out, err := exec.Command("grub-install", "--version").Output()
	if err != nil {
		out, err = exec.Command("bash", "-c", "ls /boot | grep grub").Output()
		if err != nil {
			out, err = exec.Command("bash", "-c", "ls /boot/boot | grep grub").Output()
			if err != nil {
				return "", err
			}
		}
	}

	output := strings.ToLower(string(out))

	if strings.Contains(output, "grub2") {
		boot = "grub2"
	} else if strings.Contains(output, "grub") {
		boot = "grub"
	}

	return boot, nil
}

func GetSystemLogManager() (syslog string, err error) {
	out, err := exec.Command("bash", "-c", "sudo lsof +D /var/log | grep /var/log/syslog | cut -f1 -d' '").Output()
	if err != nil {
		out, err := exec.Command("bash", "-c", "service rsyslog status").Output()
		if err != nil {
			return "", err
		}
		output := strings.ToLower(string(out))
		if strings.Contains(output, "active (running)") {
			syslog = "rsyslog"
		} else {
			syslog = "syslog-ng"

		}

	} else {
		output := strings.ToLower(string(out))
		if strings.Contains(output, "syslog-ng") {
			syslog = "syslog-ng"
		} else {
			syslog = "rsyslog"
		}
	}

	return syslog, nil
}

func GetLSM() (lsm string, err error) {
	out, err := exec.Command("bash", "-c", "sudo apparmor_status").Output()
	if err != nil {
		out, err = exec.Command("bash", "-c", "sestatus").Output()
		if err != nil {
			return "", err
		} else {
			output := strings.ToLower(string(out))
			space := regexp.MustCompile(`\s+`)
			output = space.ReplaceAllString(output, " ")
			if strings.Contains(output, "selinux status: enabled") {
				lsm = "selinux"
			}
		}
	} else {
		output := strings.ToLower(string(out))
		if strings.Contains(output, "apparmor module is loaded") {
			lsm = "apparmor"
		}
	}
	return lsm, nil
}

func getPlatformVersion(output, platform string) string {
	flagRe := regexp.MustCompile(`version[_id]*=([^ \n]*)`)
	vals := flagRe.FindStringSubmatch(output)
	if len(vals) > 1 {
		switch platform {
		case "rhel":
			return vals[1][:1] // Get the major version only, examaple: 7.6 will return 7
		case "ubuntu":
			return vals[1][:2] // Get the major version only, examaple: 18.04 will return 18
		default:
			return ""
		}
	}

	return ""
}
