#!/bin/sh

INPUT_FILE_NAME="main.go"

OUTPUT_FILE_NAME="main"

if [ "$MODE" == "production" ]; then
	
    go build $INPUT_FILE_NAME

    if [ $? -eq 0 ] && [ -f $OUTPUT_FILE_NAME ]; then
        ./main
        echo "Start production server"
    else
        echo "Build failed"
    fi
        
else
    echo "start development server"
    go run $INPUT_FILE_NAME
fi