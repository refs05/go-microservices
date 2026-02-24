# Todo

1. Apply Microservices Architecture
2. Implement context

# Intro

This is lightweight practice to learn implementing microservices using g-RPC.

## Table of Contents

- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation & Setup](#installation--setup)
- [Building Proto Files](#building-proto-files)
- [Running Services](#running-services)
- [Architecture](#architecture)

## Project Structure

```
go-microservices/
├── main.go                          # Entry point
├── go.mod                           # Go module definition
├── client/                          # gRPC client
│   └── main.go
├── common/                          # Shared code
│   └── model/
│       ├── garage.proto            # Garage service proto
│       ├── garage.pb.go            # Generated proto code
│       ├── garage_grpc.pb.go       # Generated gRPC code
│       ├── user.proto              # User service proto
│       ├── user.pb.go              # Generated proto code
│       ├── user_grpc.pb.go         # Generated gRPC code
│       └── google/
│           └── protobuf/
│               └── empty.proto     # Standard Google proto
└── services/                        # Microservices
    ├── garage-service/
    │   └── main.go
    └── user-service/
        └── main.go
```

## Prerequisites

Before you begin, ensure you have the following installed:

1. **Go** (v1.20 or higher)
   - Download from [golang.org](https://golang.org/dl/)
   - Verify: `go version`

2. **Protocol Buffer Compiler** (protoc)
   - Download from [protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf/releases)
   - Verify: `protoc --version`

3. **gRPC Go Plugins**
   - Generate Go code from proto files
   - Used by protoc during compilation

## Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-microservices.git
cd go-microservices
```

### 2. Install Dependencies

```bash
go mod download
go mod tidy
```

### 3. Install Protocol Buffer Plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Ensure your `$GOPATH/bin` is in your system `PATH`:

**Windows (PowerShell):**

```powershell
[Environment]::GetEnvironmentVariable('GOPATH','User')
# Should return: C:\Users\<YourUsername>\go
# Add C:\Users\<YourUsername>\go\bin to your PATH
```

**Linux/Mac:**

```bash
echo $GOPATH/bin
```

## Building Proto Files

The proto definitions are located in `common/model/`:

### Compile All Proto Files

From the project root, run:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=common/model common/model/*.proto
```

**Explanation of flags:**

- `--go_out ./common/model` - Output directory for generated protobuf code
- `--go_opt=paths=source_relative` - Generate files in same directory as proto
- `--go-grpc_out ./common/model` - Output directory for generated gRPC code
- `--go-grpc_opt=paths=source_relative` - Generate gRPC files relative to proto location
- `--proto_path=common/model` - Where protoc looks for imports (including Google's standard library)

### Output Files

This generates:

- `common/model/garage.pb.go` - Protobuf marshaling code for garage service
- `common/model/garage_grpc.pb.go` - gRPC service stubs for garage service
- `common/model/user.pb.go` - Protobuf marshaling code for user service
- `common/model/user_grpc.pb.go` - gRPC service stubs for user service

## Running Services

### User Service

```bash
go run .\services\user-service\main.go
```

### Garage Service

```bash
go run .\services\garage-service\main.go
```

### Run Client

```bash
go run .\client\main.go
```

### Running with `go build`

For production, build binaries:

```powershell
# Build user service
go build -o bin/user-service.exe .\services\user-service\main.go

# Build garage service
go build -o bin/garage-service.exe .\services\garage-service\main.go

# Build client
go build -o bin/client.exe .\client\main.go

# Run binaries
.\bin\user-service.exe
.\bin\garage-service.exe
.\bin\client.exe
```

## Architecture

### Services

#### User Service

Manages user-related operations via gRPC.

**Proto Definition:** `common/model/user.proto`

#### Garage Service

Manages garage-related operations via gRPC.

**Proto Definition:** `common/model/garage.proto`

Both services use Protocol Buffers for data serialization and gRPC for inter-service communication.

### Communication

- Services communicate using **gRPC** (HTTP/2 based RPC framework)
- Data is serialized using **Protocol Buffers** (efficient binary format)
- Standard Google types (Empty, Timestamp, etc.) are imported from `google/protobuf/`

## Troubleshooting

### Proto Import Errors

If you get errors like `Import "google/protobuf/empty.proto" was not found`:

1. Ensure `google/protobuf/empty.proto` exists in `common/model/google/protobuf/`
2. Use `--proto_path=common/model` when running protoc
3. Check protoc version: `protoc --version` (should be v3+)

### gRPC Module Not Found

If you see `could not import google.golang.org/grpc/status`:

```bash
go get google.golang.org/grpc
go get google.golang.org/protobuf
go mod tidy
```

### GOPATH/bin Not in PATH

If protoc plugins aren't found:

```bash
# Windows PowerShell - Add to profile
$env:PATH += ";$(go env GOPATH)\bin"

# Linux/Mac - Add to .bashrc or .zshrc
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Building & Running

### Quick Start

```bash
# 1. Install dependencies
go mod tidy

# 2. Compile proto files
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=common/model common/model/garage.proto common/model/user.proto

# 3. Run services (in separate terminals)
go run .\services\user-service\main.go
go run .\services\garage-service\main.go

# 4. Test with client
go run .\client\main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues and questions, please open an issue on the GitHub repository.
