# Norgannon

Norgannon is a simple file encryption tool written in Go. It allows you to easily encrypt and decrypt files using a password.

## Features

- Encrypt files with AES-256-GCM encryption
- Password-based encryption using SHA-256 for key derivation
- Simple command-line interface
- Cross-platform compatibility

## Installation

### Prerequisites

- Go 1.18 or higher

#### Installing Go

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install golang-go
```

**Fedora/RHEL/CentOS:**
```bash
sudo dnf install golang
```

**macOS (using Homebrew):**
```bash
brew install go
```

**Windows:**
Download and run the MSI installer from [golang.org/dl/](https://golang.org/dl/)

### Building from source

```bash
# Clone the repository
git clone https://github.com/ferris/norgannon.git
cd norgannon

# Build the project
go build -o norgannon .

# (Optional) Install to your PATH
go install
```

## Usage

### Encrypting a file

```bash
norgannon [file path] [password]
```

This will create a new file with the `.encrypted` extension.

### Decrypting a file

```bash
norgannon -d [encrypted file path] [password]
```

If the file has a `.encrypted` extension, it will be removed for the output file. Otherwise, the output file will have a `.decrypted` extension.

### Help

```bash
norgannon -h
```

## Security Notes

- The password is never stored
- Uses industry-standard AES-256-GCM encryption
- Each encryption uses a unique nonce
- Authenticated encryption protects against tampering

## License

See the [LICENSE](LICENSE) file for details. 