# GitDB Go SDK Publishing Guide

## Overview
This guide covers publishing the GitDB Go SDK via GitHub. Go packages are published by pushing to GitHub and making them available through Go modules.

## Prerequisites

### 1. GitHub Repository
- Create a new repository on GitHub (e.g., `gitdb-go-client`)
- Repository should be public for free access
- Repository should contain the Go SDK code

### 2. Go Installation
Ensure you have Go installed:
```bash
# Check if Go is installed
go version

# If not installed, download from https://golang.org/dl/
```

### 3. Git Installation
Ensure you have Git installed:
```bash
# Check if Git is installed
git --version

# If not installed, download from https://git-scm.com/
```

## Publishing Steps

### 1. Verify Package Configuration
Check the `go.mod` file:
```go
module github.com/karthikeyanV2K/gitdb-client

go 1.19
```

### 2. Test the Package
```bash
cd gitdb-go-client

# Test the package
go test ./...

# Build the package
go build ./...

# Validate go.mod
go mod tidy
go mod verify
```

### 3. Initialize Git Repository
```bash
# Initialize git repository
git init

# Add all files
git add .

# Make initial commit
git commit -m "Initial commit: GitDB Go SDK"
```

### 4. Create GitHub Repository
1. Go to https://github.com/new
2. Create a new repository named `gitdb-go-client`
3. Make it public
4. Don't initialize with README (we already have one)

### 5. Push to GitHub
```bash
# Add remote origin
git remote add origin https://github.com/karthikeyanV2K/gitdb-go-client.git

# Rename branch to main
git branch -M main

# Push to GitHub
git push -u origin main
```

### 6. Create Release Tag
```bash
# Create a git tag for the version
git tag v1.0.0

# Push the tag
git push origin v1.0.0
```

### 7. Create GitHub Release
1. Go to your repository on GitHub
2. Click "Releases" → "Create a new release"
3. Tag version: `v1.0.0`
4. Title: `GitDB Go SDK v1.0.0`
5. Description: Add release notes
6. Publish release

## Package Information

### Package Name
- **Module**: `github.com/karthikeyanV2K/gitdb-client`
- **Registry**: Go Modules (via GitHub)
- **Current Version**: 1.0.0
- **GitHub**: https://github.com/karthikeyanV2K/gitdb-go-client

### Dependencies
- `net/http` - HTTP client
- `encoding/json` - JSON handling
- `bytes` - Buffer operations
- `io` - I/O operations
- `fmt` - Formatting
- `strings` - String operations

### Features
- Simple Go interface
- Full CRUD operations
- MongoDB-style query operators
- Collection management
- Error handling
- Type safety

## Usage Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/karthikeyanV2K/gitdb-client/gitdb"
)

func main() {
    client := gitdb.NewClient("your-token", "owner", "repo")
    
    document := map[string]interface{}{
        "name": "John Doe",
        "email": "john@example.com",
    }
    
    id, err := client.Insert("users", document)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Inserted: %s\n", id)
}
```

## Directory Structure

```
gitdb-go-client/
├── go.mod              # Go module configuration
├── README.md           # Documentation
├── PUBLISHING.md       # This guide
├── gitdb/
│   └── client.go       # Main client implementation
└── examples/
    └── basic_usage.go  # Usage examples
```

## Troubleshooting

### Common Issues

1. **Module Not Found**
   - Ensure repository is public
   - Check that go.mod is in the correct location
   - Verify module path matches GitHub repository

2. **Import Errors**
   ```bash
   # Update dependencies
   go mod tidy
   go mod download
   ```

3. **Git Tag Issues**
   - Ensure git tags follow semantic versioning (v1.0.0)
   - Push tags to GitHub: `git push origin --tags`

4. **Go Version Issues**
   - Ensure Go version is 1.19 or higher
   - Update Go if needed

### Support

- Go Documentation: https://golang.org/doc/
- Go Modules: https://golang.org/ref/mod
- GitHub: https://github.com/

## Post-Publishing

### 1. Verify Installation
```bash
# Test installation
go get github.com/karthikeyanV2K/gitdb-client

# Use in a new project
mkdir test-project
cd test-project
go mod init test
go get github.com/karthikeyanV2K/gitdb-client
```

### 2. Update Documentation
- Update the main README.md to include Go SDK
- Add installation instructions
- Include usage examples

### 3. Version Management
To update the package version:
1. Update version in go.mod (if needed)
2. Create new git tag: `git tag v1.0.1`
3. Push tag: `git push origin v1.0.1`
4. Create new GitHub release

### 4. Monitoring
- Monitor repository stars and forks
- Respond to issues on GitHub
- Update dependencies as needed

## Quick Publish Script

Create a `publish.bat` file for Windows:
```batch
@echo off
echo Publishing GitDB Go SDK to GitHub...
cd gitdb-go-client
git add .
git commit -m "Update for release"
git tag v1.0.0
git push origin main
git push origin v1.0.0
echo Go SDK published successfully!
echo Visit: https://github.com/karthikeyanV2K/gitdb-go-client
pause
```

## Manual Publishing (Alternative)

If automated script doesn't work:

1. **Create Repository on GitHub**
   - Go to GitHub and create new repository
   - Name it `gitdb-go-client`
   - Make it public

2. **Push Code Manually**
   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin https://github.com/karthikeyanV2K/gitdb-go-client.git
   git branch -M main
   git push -u origin main
   ```

3. **Create Release**
   - Go to GitHub repository
   - Click "Releases" → "Create a new release"
   - Tag version: `v1.0.0`
   - Add release notes
   - Publish release

## Success Criteria

✅ Package is available on GitHub
✅ Package can be installed with `go get github.com/karthikeyanV2K/gitdb-client`
✅ Documentation is accessible
✅ Examples work correctly
✅ All tests pass
✅ Go modules work properly

## Installation Instructions for Users

```bash
# Install via go get
go get github.com/karthikeyanV2K/gitdb-client

# Or add to go.mod
require github.com/karthikeyanV2K/gitdb-client v1.0.0
```

## Package URL
Once published, the package will be available at:
https://github.com/karthikeyanV2K/gitdb-go-client

Users can install it with:
```bash
go get github.com/karthikeyanV2K/gitdb-client
``` 