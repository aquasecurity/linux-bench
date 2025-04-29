#!/bin/sh -e
if [ "$1" == "install" ]; then
  if [ -d /host ]; then
    mkdir -p /host/cfg/
    yes | cp -rf cfg/* /host/cfg/
    yes | cp -rf /usr/local/bin/kube-bench /host/
    echo "==============================================="
    echo "linux-bench is now installed on your host       "
    echo "Run ./linux-bench to perform a security check   "
    echo "==============================================="
  else
    echo "Usage:"
    echo "  install: docker run --rm -v \`pwd\`:/host aquasec/linux-bench install"
    echo "  run:     docker run --rm --pid=host aquasec/linux-bench [command]"
    exit
  fi
else
  exec linux-bench "$@"
fi
