FROM golang:1.23.4 AS build
WORKDIR /go/src/github.com/aquasecurity/linux-bench/
COPY makefile makefile
COPY go.mod go.sum ./
COPY app.go main.go root.go utils.go .
RUN CGO_ENABLED=0 go build -o linux-bench && cp linux-bench /go/bin/linux-bench

FROM alpine:3.21.2 AS run
WORKDIR /opt/linux-bench/
# add GNU ps for -C, -o cmd, --no-headers support and add findutils to get GNU xargs
# https://github.com/aquasecurity/kube-bench/issues/109
# https://github.com/aquasecurity/kube-bench/issues/1656
RUN apk --no-cache add procps findutils

# Upgrading apk-tools to remediate CVE-2021-36159 - https://snyk.io/vuln/SNYK-ALPINE314-APKTOOLS-1533752
# https://github.com/aquasecurity/kube-bench/issues/943
RUN apk --no-cache upgrade apk-tools

# Add glibc for running oc command 
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
RUN apk add gcompat
RUN apk add jq

# Add bash for running helper scripts
RUN apk add bash

ENV PATH=$PATH:/usr/local/mount-from-host/bin:/go/bin

COPY --from=build /go/bin/linux-bench /usr/local/bin/linux-bench
COPY entrypoint.sh .
COPY cfg/ cfg/
ENTRYPOINT ["./entrypoint.sh"]
CMD ["install"]
