# Setup Instructions

Quick reference for setting up the Go Microservices project.

## Prerequisites

### Required

- **Go** v1.20+ — [Download](https://golang.org/dl/)
- **Protocol Buffer Compiler (protoc)** v3+ — [Download](https://github.com/protocolbuffers/protobuf/releases)
- **Git** — [Download](https://git-scm.com/)

### Verify Installation

```bash
go version
protoc --version
git --version
```

## Windows Setup

### 1. Download & Install Go

1. Download from [golang.org](https://golang.org/dl/)
2. Run installer (accepts all defaults)
3. Verify: `go version`

### 2. Download & Install Protoc

1. Download from [GitHub releases](https://github.com/protocolbuffers/protobuf/releases)
   - Look for `protoc-X.X.X-win64.zip`
2. Extract to `C:\protoc`
3. Add `C:\protoc\bin` to your PATH:
   - Press `Win + X` → System
   - Advanced system settings → Environment Variables
   - Edit `PATH` → Add `C:\protoc\bin`
4. Verify: `protoc --version` (restart terminal if needed)

### 3. Clone Project

```powershell
git clone https://github.com/yourusername/go-microservices.git
cd go-microservices
```

### 4. Install Go Dependencies

```powershell
go mod download
go mod tidy
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 5. Compile Proto Files

```powershell
protoc --go_out ./common/model --go_opt=paths=source_relative --go-grpc_out ./common/model --go-grpc_opt=paths=source_relative --proto_path=common/model common/model/*.proto
```

### 6. Run Services

**Terminal 1 - User Service:**

```powershell
go run .\services\user-service\main.go
```

**Terminal 2 - Garage Service:**

```powershell
go run .\services\garage-service\main.go
```

**Terminal 3 - Client:**

```powershell
go run .\client\main.go
```

## macOS Setup

### 1. Install Go

Using Homebrew:

```bash
brew install go
go version
```

### 2. Install Protoc

```bash
brew install protobuf
protoc --version
```

### 3. Clone Project

```bash
git clone https://github.com/yourusername/go-microservices.git
cd go-microservices
```

### 4. Install Go Dependencies

```bash
go mod download
go mod tidy
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 5. Compile Proto Files

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=common/model common/model/*.proto
```

### 6. Run Services

**Terminal 1:**

```bash
go run ./services/user-service/main.go
```

**Terminal 2:**

```bash
go run ./services/garage-service/main.go
```

**Terminal 3:**

```bash
go run ./client/main.go
```

## Linux Setup

### 1. Install Go

**Ubuntu/Debian:**

```bash
sudo apt-get update
sudo apt-get install golang-go
go version
```

**Fedora/RHEL:**

```bash
sudo dnf install go
go version
```

### 2. Install Protoc

**Ubuntu/Debian:**

```bash
sudo apt-get install protobuf-compiler
protoc --version
```

**Fedora/RHEL:**

```bash
sudo dnf install protobuf-compiler
protoc --version
```

### 3. Clone Project

```bash
git clone https://github.com/yourusername/go-microservices.git
cd go-microservices
```

### 4. Install Go Dependencies

```bash
go mod download
go mod tidy
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 5. Compile Proto Files

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=common/model common/model/garage.proto common/model/user.proto
```

### 6. Run Services

**Terminal 1:**

```bash
go run ./services/user-service/main.go
```

**Terminal 2:**

```bash
go run ./services/garage-service/main.go
```

**Terminal 3:**

```bash
go run ./client/main.go
```

## Docker Setup (Optional)

Build and run services in containers:

```bash
docker-compose up
```

## Troubleshooting

### "protoc: command not found"

- Protoc not installed or not in PATH
- Reinstall protoc and add bin directory to PATH
- Verify with `protoc --version`

### "go: command not found"

- Go not installed
- Download and run installer from golang.org
- Verify with `go version`

### Proto import errors

- Ensure `--proto_path=common/model` is used
- Check that `google/protobuf/empty.proto` exists
- Rebuild with: `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=common/model common/model/*.proto`

### "missing go.sum entry"

```bash
go mod tidy
```

### Port already in use

- Services use default gRPC ports
- Change ports in service main.go files
- Or kill process using the port

## Need Help?

- Check [README.md](../README.md) for general info
- See [CONTRIBUTING.md](../CONTRIBUTING.md) for contribution guidelines
- Open an issue on GitHub for bugs/features
