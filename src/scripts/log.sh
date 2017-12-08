#!/bin/sh

# Make info file about this build
mkdir -p /etc/BUILDS/

# Get current time and assign it to env variable DATE
DATE=`date -u +"%d-%m-%YT%H:%M:%SZ"`

# Generating log
echo "Build of alpine-golang: $GO_VERSION, date: $DATE" > /etc/BUILDS/alpine-golang-log.txt