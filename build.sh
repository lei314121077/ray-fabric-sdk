#!/usr/bin/env bash
#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o data-transfer-chaincode main.go


