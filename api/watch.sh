#/bin/bash

# Watch requires a script because Win and Unix systems require 
# different file extensions for binary when run using the "air" tool
if command -v air > /dev/null; then
    if [[ $1 == win ]]; then
        echo "Running watch for Windows OS"
        # The --build.bin value here is using a different path separator.
        # There's an open PR for an issue here https://github.com/air-verse/air/pull/590
        air --build.cmd "go build -o bin/main.exe cmd/main.go" --build.bin "bin\main.exe"
    else
        air --build.cmd "go build -o bin/main cmd/main.go" --build.bin "bin\main"
    fi
else
    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice
    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then
        go install github.com/air-verse/air@latest
        cd api; air
        echo "Watching..."
    else \
        echo "You chose not to install air. Exiting..."
        exit 1
    fi
fi