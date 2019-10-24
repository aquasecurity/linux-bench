[![Build Status](https://travis-ci.org/aquasecurity/linux-bench.svg?branch=master)](https://travis-ci.org/aquasecurity/linux-bench)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Coverage Status][cov-img]][cov]

[cov-img]: https://codecov.io/github/aquasecurity/linux-bench/branch/master/graph/badge.svg
[cov]: https://codecov.io/github/aquasecurity/linux-bench

Linux-bench is a Go application that checks whether the Linux operating system is configured securely by running the checks documented in the CIS Distribution Independent Linux Benchmark.

Tests are configured with YAML files, making this tool easy to update as test specifications evolve.

## CIS Linux Benchmark support

linux-bench currently supports tests for benchmark version 1.1.0 only.

linux-bench will determine the test set to run on the host machine based on the following:
- **Operating system platform - ubuntu/debian/rhel/coreos**
- **Boot loader - grub/grub2**
- **System logging tool - rsyslog/syslog-ng**
- **Lsm - selinux/apparmor**

## Installation

### Installing from sources

Install [Go](https://golang.org/doc/install), then
clone this repository and run as follows (assuming your [\$GOPATH is set](https://github.com/golang/go/wiki/GOPATH)):


```shell
go get github.com/aquasecurity/linux-bench
cd $GOPATH/src/github.com/aquasecurity/linux-bench
go build -o linux-bench .

# See all supported options
./linux-bench --help

# Run checks
./linux-bench

# Run checks for specified linux cis version
./linux-bench --version <version>
```

# Tests

Tests are specified in definition files `cfg/<version>/definitions.yaml.`

Where `<version>` is the version of linux cis for which the test applies.

# Contributing

We welcome PRs and issue reports.
