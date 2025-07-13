@echo off
echo ========================================
echo Publishing GitDB Go SDK to GitHub
echo ========================================
echo.

echo Checking Go installation...
go version >nul 2>&1
if errorlevel 1 (
    echo ERROR: Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

echo Go found:
go version

echo.
echo Checking Git installation...
git --version >nul 2>&1
if errorlevel 1 (
    echo ERROR: Git is not installed or not in PATH
    echo Please install Git from https://git-scm.com/
    pause
    exit /b 1
)

echo Git found:
git --version

echo.
echo Testing the package...
go test ./...
if errorlevel 1 (
    echo ERROR: Tests failed
    pause
    exit /b 1
)

echo.
echo Building the package...
go build ./...
if errorlevel 1 (
    echo ERROR: Build failed
    pause
    exit /b 1
)

echo.
echo Validating go.mod...
go mod tidy
go mod verify
if errorlevel 1 (
    echo ERROR: go.mod validation failed
    pause
    exit /b 1
)

echo.
echo IMPORTANT: Go SDK publishing requires:
echo 1. GitHub repository is created
echo 2. Code is pushed to GitHub
echo 3. Release tag is created
echo.
echo This script will initialize git and prepare for GitHub push.
echo.

set /p confirm="Do you want to initialize git repository and prepare for GitHub? (y/N): "
if /i not "%confirm%"=="y" (
    echo Publishing cancelled.
    pause
    exit /b 0
)

echo.
echo Initializing git repository...
git init
if errorlevel 1 (
    echo ERROR: Failed to initialize git repository
    pause
    exit /b 1
)

echo.
echo Adding files to git...
git add .
if errorlevel 1 (
    echo ERROR: Failed to add files to git
    pause
    exit /b 1
)

echo.
echo Making initial commit...
git commit -m "Initial commit: GitDB Go SDK"
if errorlevel 1 (
    echo ERROR: Failed to make initial commit
    pause
    exit /b 1
)

echo.
echo Creating git tag v1.0.0...
git tag v1.0.0
if errorlevel 1 (
    echo ERROR: Failed to create git tag
    pause
    exit /b 1
)

echo.
echo ========================================
echo SUCCESS: Go SDK prepared for GitHub
echo ========================================
echo.
echo Next steps:
echo 1. Create GitHub repository: gitdb-go-client
echo 2. Add remote: git remote add origin https://github.com/karthikeyanV2K/gitdb-go-client.git
echo 3. Push code: git push -u origin main
echo 4. Push tag: git push origin v1.0.0
echo 5. Create GitHub release
echo.
echo Package will be available at:
echo https://github.com/karthikeyanV2K/gitdb-go-client
echo.
echo Users can install with:
echo go get github.com/karthikeyanV2K/gitdb-client
echo.
pause 