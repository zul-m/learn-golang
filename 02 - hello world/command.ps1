# Create a Go module.
go mod init learngo

# Running a Go program:

## 1. Install(copy) the binary to location `~/go/bin`.
go install

## 2. Creates the binary inside the directory.
go build

## 3. Compiles the file to a temporary location and runs the file from that location. Use `--work` to know the location where `go run` compiles the file to.
go run main.go
go run --work main.go
