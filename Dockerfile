FROM golang:1.23.6 AS builder

ENV OUTPUT_DIR=/out

COPY . /linux-bench
WORKDIR /linux-bench

RUN GOSUMDB=off CGO_ENABLED=0 GOOS=linux go build -o ${OUTPUT_DIR}/linux-bench .
RUN cp LICENSE ${OUTPUT_DIR}/LICENSE
RUN cp README.md ${OUTPUT_DIR}/README.md
RUN cp -rf cfg ${OUTPUT_DIR}/cfg

WORKDIR /
RUN rm -rf /linux-bench
