#!/usr/bin/env bash

ProPath ="$HOME/go/data-transfer-chaincode/"
LibrayPath ="$ProPath/data-transfer-share-libray"
SdkApiPATH ="$ProPath/fabricsdk-api"
CCPath ="$ProPath/transfer-chaincode"

# settings this pro gpath
export GOPATH="$LibrayPath:$CCPath:$SdkApiPATH"

# build your pro app
cd $SdkApiPATH && ./build.sh && cd $CCPath && ./build.sh

# run your pro app
cd $SdkApiPATH && ./farbic-api && cd $CCPath && ./data-transfer-cc

