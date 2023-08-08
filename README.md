## go-practice-api

Go Rest API with generic CRUD

### Instructions

1. Install [air](https://github.com/cosmtrek/air) if you haven't already

```sh
go install github.com/cosmtrek/air@latest
```

2a. On Windows, install a version of MinGW to C:/mingw64 (for go-sqlite3)
2b. On Mac, add the following to .zshrc (or equivalent):

```sh
export GOROOT=/usr/local/go
export GOPATH=$HOME/go

export PATH=$GOPATH/bin:$PATH
```

3. Run the server with `bash ./develop.sh` (or build with `./build.sh`)
4. If using Thunder Client in VS Code, import thunder-client.json and test our the routes
