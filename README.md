# System Info Tool

A fast, cross-platform system information display tool written in Go, inspired by neofetch.

```
   .--.     user@hostname
  |o_o |    ---------------
  |:_/ |    OS        : Ubuntu 22.04.3 LTS
 //   \ \   Kernel    : 5.15.0-91-generic
(|     | )  Uptime    : 2d 14h 32m
/'\\_   _/`\ Shell     : zsh
\___)=(___/  CPU       : Intel Core i7-12700K (16 cores)
             GPU       : NVIDIA GeForce RTX 4080
             Memory    : 8.2 GB / 32.0 GB
             Disk      : 45G / 256G (18%)
             Packages  : 2847 (dpkg)

███████████████████████████
```

## Features

- 🚀 **Fast**: Written in Go for optimal performance
- 🔄 **Cross-platform**: Supports Linux, macOS, and Windows
- 🎨 **Beautiful output**: Clean ASCII art with colorful display
- 🏗️ **Modular**: Clean architecture with separated concerns
- 🧪 **Testable**: Interfaces and dependency injection for easy testing
- 📦 **Zero dependencies**: Uses only Go standard library

## Installation

```bash
curl -sSL https://raw.githubusercontent.com/MiraWuka/Mirkafetch/refs/heads/main/install.sh | bash
```

```bash
curl -L https://raw.githubusercontent.com/MiraWuka/Mirkafetch/refs/heads/main/install.sh -o install.sh
chmod +x install.sh
./install.sh
```



### From Source

```bash
git clone https://github.com/username/sysinfo.git
cd sysinfo
make build
sudo cp bin/sysinfo /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/username/sysinfo/cmd/sysinfo@latest
```

## Usage

Simply run the command:

```bash
sysinfo
```

## Supported Information

- **User & Hostname**: Current user and system hostname
- **Operating System**: Detailed OS information
- **Kernel**: Kernel version
- **Uptime**: System uptime in human-readable format
- **Shell**: Current shell
- **CPU**: Processor information with core count
- **GPU**: Graphics card information
- **Memory**: RAM usage (used/total)
- **Disk**: Root disk usage
- **Packages**: Installed package count with package manager

## Supported Systems

### Linux
- All major distributions (Ubuntu, Debian, Fedora, Arch, etc.)
- Package managers: dpkg, rpm, pacman, apk

### macOS
- All modern macOS versions
- Package manager: Homebrew

### Windows
- Windows 10/11 (basic support)

## Development

### Prerequisites

- Go 1.21 or higher

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Development workflow
make dev
```

### Project Structure

```
sysinfo/
├── cmd/sysinfo/           # Application entry point
├── internal/
│   ├── app/              # Application logic
│   ├── collector/        # System info collection
│   ├── display/          # Output formatting
│   └── models/           # Data structures
├── pkg/utils/            # Reusable utilities
├── Makefile             # Build automation
└── README.md
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage
```

### Code Quality

```bash
# Format code
make fmt

# Lint code  
make lint

# Vet code
make vet

# Full check
make check
```

## Architecture

The application follows clean architecture principles:

- **`cmd/`**: Entry points
- **`internal/app/`**: Application orchestration
- **`internal/collector/`**: System information gathering
- **`internal/display/`**: Output formatting
- **`internal/models/`**: Data structures
- **`pkg/utils/`**: Reusable utilities

### Key Interfaces

- `Collector`: Defines system information collection
- `Display`: Defines output formatting

This design enables easy testing, mocking, and extending with new collectors or display formats.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by [neofetch](https://github.com/dylanaraps/neofetch)
- ASCII art adapted from various sources