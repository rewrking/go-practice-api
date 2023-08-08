#!/usr/bin/env bash

if [[ $OSTYPE == 'linux-gnu'* || $OSTYPE == 'cygwin'* ]]; then
	PLATFORM=linux
elif [[ $OSTYPE == 'darwin'* ]]; then
	PLATFORM=macos
elif [[ $OSTYPE == 'msys' || $OSTYPE == 'win32' ]]; then
	PLATFORM=windows
fi

if [[ $PLATFORM == "windows" ]]; then
	export CGO_ENABLED=1
	export PATH=/c/mingw64/bin:$PATH
fi


air

echo ''

EXIT_CODE=$?

exit $EXIT_CODE