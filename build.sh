#!/usr/bin/env bash

if [[ $OSTYPE == 'linux-gnu'* || $OSTYPE == 'cygwin'* ]]; then
	PLATFORM=linux
elif [[ $OSTYPE == 'darwin'* ]]; then
	PLATFORM=macos
elif [[ $OSTYPE == 'msys' || $OSTYPE == 'win32' ]]; then
	PLATFORM=windows
fi

TARGET_EXT=
if [[ $PLATFORM == "windows" ]]; then
	export CGO_ENABLED=1
	export PATH=/c/mingw64/bin:$PATH
	TARGET_EXT=.exe
fi

TARGET=go-practice-api$TARGET_EXT

go build -o ./build/$TARGET ./src/main.go

if [[ $? == 0 ]]; then
	./build/$TARGET
fi

# air

# go run ./src/*.go

echo ''

EXIT_CODE=$?

exit $EXIT_CODE