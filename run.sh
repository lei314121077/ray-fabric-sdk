#!/usr/bin/env bash

ProPath ="$HOME/go/data-transfer-chaincode/"
LibrayPath ="$ProPath/data-transfer-share-libray"
SdkApiPATH ="$ProPath/fabricsdk-api"
CCPath ="$ProPath/transfer-chaincode"

# settings this pro gpath
export GOPATH="$LibrayPath:$CCPath:$SdkApiPATH"

# build and run your fabric sdk api
cd $SdkApiPATH && ./build.sh && ./farbic-api

# build and run your fabric chaincode
cd $CCPath && ./build.sh && ./data-transfer-cc

