# Technical Context: aiyou-cli

## Technologies Used

### Core Technologies
- **Go (Golang)**: Version 1.23.4 or higher
- **AI.You API**: For accessing AI models and capabilities
- **AI.You Go SDK**: Official SDK for API interaction

### Key Libraries
- **Cobra**: Command-line interface framework
- **YAML v3**: Configuration file parsing
- **Standard Library**: For core functionality

## Development Setup

### Requirements
- Go 1.23.4 or higher
- Git for version control
- AI.You API token for testing

### Development Environment
```bash
# Clone the repository
git clone https://github.com/n1neT10ne/aiyou-cli.git
cd aiyou-cli

# Build for development
go build

# Run tests
go test ./...
```

### Build Process
The project includes a build script for creating optimized binaries:
```bash
# Build for current platform
go build

# Build for specific platforms with optimization
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o aiyou-cli.exe
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o aiyou-cli-linux

# Build for all platforms at once
cd scripts/cmd && go run main.go
```

## Technical Constraints

### API Limitations
- Requires valid AI.You API token
- Subject to API rate limits
- Dependent on API availability

### Performance Considerations
- Streaming mode increases responsiveness but requires stable connection
- Large responses may require pagination or streaming
- Token validation adds minimal overhead

### Security Considerations
- API tokens should be kept secure
- Environment variables or config files should have appropriate permissions
- No sensitive data should be logged in debug mode

## Dependencies

### Direct Dependencies
- github.com/n1neT10ne/aiyou-go-sdk v0.1.0
- github.com/spf13/cobra v1.8.1
- gopkg.in/yaml.v3 v3.0.1

### Indirect Dependencies
- github.com/inconshreveable/mousetrap v1.1.0
- github.com/spf13/pflag v1.0.5

## Deployment

### Distribution Methods
- Pre-built binaries for Windows and Linux
- Source installation via `go install`
- Manual build from source

### Release Process
1. Update version in `main.go`
2. Commit changes
3. Create and push a new tag
4. GitHub Actions workflow builds and publishes releases

## Configuration

### Configuration Sources (in order of precedence)
1. Command-line arguments
2. Environment variables
3. YAML configuration file

### Environment Variables
- `AIYOU_CLI_TOKEN`: API token

### Configuration File
Default location: `./aiyou-cli.yaml`
Format: YAML with support for multi-line strings
