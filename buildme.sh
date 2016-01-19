#!/bin/bash

# build the aws lambda function and update/upload it
#####################################################################
#
# Config area
#
# NOTE - LFNAME is case sensitive and must match the actual Lambda Function Name exactly
LFNAME="eventDumper"

# standard name of binary that index.js references
OUTFILE="gocode-amd64"

# for eventDumper where to send the results. This is compiled into the binary
SNSTOPIC="arn:aws:sns:us-west-2:123456789012:snsTopic"

# set go build environment to linux 64 bit which AWS use
export GOOS=linux
export GOARCH=amd64

if [ -z "${LFNAME}" ]; then
    echo "No Lambda Function name provided!!!"
    exit 1
fi

if [ -z "${SNSTOPIC}" ]; then
    echo "error - No SNS Topic name provided"
    exit 1
fi

#
# End Config area Do not change code below
#
#####################################################################

# build the binary
go build -o ${OUTFILE} -ldflags="-X main.snsTopic=${SNSTOPIC}"

if [ "x$?" != "x0" ]; then
    echo "Build failed and ${OUTFILE}* removed"
    rm -f ${OUTFILE}*
    exit 1
fi

zip ${OUTFILE}.zip index.js ${OUTFILE}

echo "Do you want to update the lambda function now? CTRL-C to abort"
read crap

echo "Uploading to AWS Lambda"

# upload the zip file to Lambda
aws lambda update-function-code --function-name ${LFNAME} --zip-file fileb://${OUTFILE}.zip

# clean up things. Comment out if you want to keep things
rm -f ${OUTFILE}*

echo "code uploaded to Lambda and build files removed"
echo

