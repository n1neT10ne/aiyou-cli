# 🤖 aiyou-cli

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.23-blue.svg)](https://golang.org/dl/)

A minimalist and efficient command-line client for interacting with the AI.You API. Designed to be user-friendly while offering maximum flexibility.

## ✨ Features

* 🎯 Simple and intuitive command-line interface
* 🔧 Flexible configuration via YAML file or CLI arguments
* 🤖 Main features:
  * List available models
  * Send messages to models
  * AI.You assistants support
* 🔄 Streaming support for real-time responses
* ⚡ Automatic retry handling
* 🛠️ Flexible configuration options
* 🔍 Debug mode for development
* 🌡️ Response temperature control

## 📝 TODO

* [x] Add standard input support for message input to enable command piping to the binary
* [x] Support multi-line system prompts in the YAML configuration file

## 📦 Installation

### From Source

#### Option 1: Direct Installation
```bash
go install github.com/n1neT10ne/aiyou-cli@latest
```

#### Option 2: Local Build
Clone the repository and build it locally:
```bash
# Clone the repository
git clone https://github.com/n1neT10ne/aiyou-cli.git
cd aiyou-cli

# Method 1: Build for your current platform only
go build

# Method 2: Build for a specific platform
# For Windows (optimized size)
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o aiyou-cli.exe

# For Linux (optimized size)
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o aiyou-cli-linux

# Optional: Further compress with UPX (if installed)
# Windows: upx --best --lzma aiyou-cli.exe
# Linux: upx --best --lzma aiyou-cli-linux

# Method 3: Build for all platforms at once
cd scripts/cmd && go run main.go && cd ../..
```

The build script (scripts/builder/build.go) will create binaries in the dist/ directory:
- dist/aiyou-cli.exe (Windows binary)
- dist/aiyou-cli-linux (Linux binary)

### Pre-built Binaries
You can download pre-built binaries for Windows and Linux from the [releases page](https://github.com/n1neT10ne/aiyou-cli/releases).

## 🚀 Release Process

To create a new release with native binaries for Windows and Linux:

1. Update version in `main.go` (`const version = "vX.Y.Z"`)
2. Commit your changes
3. Create and push a new tag:
```bash
git tag vX.Y.Z
git push origin vX.Y.Z
```

This will trigger the GitHub Actions workflow that:
- Builds native binaries for Windows (.exe) and Linux
- Creates a GitHub release with these binaries
- Generates release notes automatically

## 🚀 Usage

### Configuration

The CLI can be configured in three ways, with the following order of precedence (from highest to lowest):

1. Command line arguments (`--token` or `-k`)
2. Environment variables (`AIYOU_CLI_TOKEN`)
3. YAML configuration file (`token` field)

### Environment Variables

The API token can be set via the environment variable `AIYOU_CLI_TOKEN`. The command varies depending on your shell:

```bash
# PowerShell
$env:AIYOU_CLI_TOKEN="your-token"

# Windows Command Prompt (cmd.exe)
set AIYOU_CLI_TOKEN=your-token

# Linux/macOS
export AIYOU_CLI_TOKEN="your-token"
```

Note: These commands set the variable for the current session only. For permanent configuration:
- Windows: Set through System Properties > Environment Variables
- Linux/macOS: Add to your shell's configuration file (~/.bashrc, ~/.zshrc, etc.)

Alternatively, you can use the configuration file or command line flag for persistent settings.

### YAML Configuration File

```yaml
# aiyou-cli.yaml
debug: false
temperature: 0.7
retry: 3
# Single-line prompt
prompt: "System instructions"
# OR multi-line prompt (using YAML's literal style with |)
prompt: |
  You are an AI assistant.
  Please respond in a concise manner.
  Use examples when relevant.
token: "your-token"
model: "model-name"
stream: true
assistantId: "assistant-id"
```

### Command Line Arguments

```bash
# List available models
aiyou-cli --list-models -k "your-token"

# Send a simple message
aiyou-cli -k "your-token" -m "Your message" -M "model-name"

# Read message from standard input
echo "Your message" | aiyou-cli -k "your-token" -i -M "model-name"
# OR
cat file.txt | aiyou-cli -k "your-token" -i -M "model-name"

# With advanced options
aiyou-cli \
  -k "your-token" \
  -m "Your message" \
  -M "model-name" \
  -t 0.7 \
  -p "System instructions" \
  -s \
  -d
```

### Available Options

```
Flags:
  --config string     Configuration file path (default: "./aiyou-cli.yaml")
  -d, --debug        Enable debug mode (default: false)
  -t, --temp float   Temperature (0.0-2.0) (default: 1.0)
  -r, --retry int    Number of retries (default: 0)
  -p, --prompt       System instructions
  -m, --message      Message to send to the LLM
  -k, --token        API token
  -M, --model        LLM to use
  -s, --stream       Enable streaming mode (default: false)
  -a, --assistant    Assistant ID
  -L, --list-models  List available models
  -i, --stdin        Read message from standard input
  -h, --help         Display help
```

## 🔄 Streaming Mode

Streaming mode allows receiving the response as it's being generated:

```bash
aiyou-cli -k "your-token" -m "Your message" -s
```

## ⚠️ Error Handling

The CLI includes robust error handling with:

* Required parameter validation
* Value checks (e.g., temperature between 0.0 and 2.0)
* Automatic retries for network errors
* Clear and informative error messages

## 📁 Project Structure

The project is organized as follows:

```
aiyou-cli/
├── main.go              # Main CLI entry point
├── config.go            # Configuration handling
├── aiyou-cli.yaml      # Default configuration file
├── dist/               # Multi-platform builds output directory
│   ├── aiyou-cli.exe   # Windows binary (optimized)
│   └── aiyou-cli-linux # Linux binary (optimized)
└── scripts/            # Build tools
    ├── builder/        # Build package with optimization logic
    │   └── build.go    # Multi-platform build implementation
    └── cmd/            # Build command entry point
        └── main.go     # Build script main program
```

### Build Optimization

The project uses several techniques to optimize binary size:

1. Debug symbols stripping via `-ldflags="-s -w"`:
   - `-s`: Remove symbol table
   - `-w`: Remove DWARF debugging information

2. Optional UPX compression:
   - Further reduces binary size by 50-70%
   - Automatically applied in GitHub releases
   - Can be manually applied to local builds

## 🔗 SDK

This CLI uses the official AI.You Go SDK available here: [aiyou-go-sdk](https://github.com/n1neT10ne/aiyou-go-sdk)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**Cyrille BARTHELEMY**

* Github: [@n1neT10ne](https://github.com/n1neT10ne)
