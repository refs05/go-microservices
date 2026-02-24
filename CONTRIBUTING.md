# Contributing to Go Microservices

Thank you for your interest in contributing! Please follow these guidelines.

## Code of Conduct

- Be respectful and professional
- Provide constructive feedback
- Report issues responsibly

## Getting Started

### 1. Fork & Clone

```bash
git clone https://github.com/your-username/go-microservices.git
cd go-microservices
```

### 2. Create a Feature Branch

```bash
git checkout -b feature/your-feature-name
```

### 3. Set Up Development Environment

```bash
# Install dependencies
go mod download
go mod tidy

# Install protoc plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Compile proto files
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=common/model common/model/garage.proto common/model/user.proto
```

## Development Workflow

### Modifying Proto Files

1. Edit `.proto` files in `common/model/`
2. Recompile:
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=common/model common/model/*.proto
   ```
3. Test your changes

### Code Style

- Follow Go conventions ([Effective Go](https://golang.org/doc/effective_go))
- Use `gofmt` for formatting
- Use `golint` for style checks

```bash
gofmt -w .
go vet ./...
```

### Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...
```

## Submission Process

### 1. Commit Your Changes

```bash
git add .
git commit -m "feat: add amazing feature"
```

**Commit message format:**

- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation
- `refactor:` Code refactor
- `test:` Adding tests
- `chore:` Maintenance

### 2. Push to Your Fork

```bash
git push origin feature/your-feature-name
```

### 3. Create a Pull Request

- Provide a clear description of changes
- Reference any related issues (#123)
- Include screenshots/examples if applicable
- Ensure all tests pass

## Pull Request Checklist

- [ ] Code follows Go style guidelines
- [ ] Comments added for complex logic
- [ ] All tests pass (`go test ./...`)
- [ ] Proto files recompiled if modified
- [ ] README updated if needed
- [ ] No unnecessary dependencies added

## Reporting Issues

Use GitHub Issues to report bugs:

1. **Title:** Clear, concise description
2. **Environment:** Go version, OS, architecture
3. **Steps to Reproduce:** Detailed reproduction steps
4. **Expected Behavior:** What should happen
5. **Actual Behavior:** What actually happened
6. **Error Messages:** Include full error output

### Example Issue

```
Title: User service fails to start on Windows with protoc not found

Environment:
- Go 1.21
- Windows 11
- protoc 3.24.0

Steps to Reproduce:
1. Install project
2. Run: protoc --version (works)
3. Run: go run services/user-service/main.go
4. Error occurs

Error Message:
[full error text]
```

## Questions?

Feel free to open a discussion or contact maintainers via GitHub Issues.

Thank you for contributing! 🎉
