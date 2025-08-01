# 🧹 WipeOs

<div align="center">
  <img src="assets/logo.png" alt="WipeOs Logo" width="200"/>
  <br/>
  <h3>Professional Secure File Wiping Tool</h3>
  <p>Military-grade file deletion with a beautiful terminal interface</p>
  
  ![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)
  ![License](https://img.shields.io/badge/license-MIT-green)
  ![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey)
  ![Downloads](https://img.shields.io/github/downloads/joao-rrondon/wipeOs/total)
  ![Stars](https://img.shields.io/github/stars/joao-rrondon/wipeOs)
</div>

---

## ✨ Features

- 🔒 **Military-Grade Security**: Uses DoD 5220.22-M standard overwriting patterns
- 🎯 **Smart Targeting**: Wipe specific files, directories, or predefined data types
- 🚀 **Blazing Fast**: Optimized Go implementation with concurrent processing
- 🛡️ **Safety First**: Confirmation prompts and dry-run mode prevent accidents
- 📊 **Detailed Reporting**: Comprehensive logging and progress indicators
- 🎨 **Beautiful Interface**: Modern terminal UI with colors and ASCII art
- 🌐 **Cross-Platform**: Works on Linux, macOS, and Windows
- 🧹 **Predefined Cleaners**: Browser data, system temp files, and more

## 🎬 Demo

<div align="center">
  <img src="assets/demo.gif" alt="WipeOs Demo" width="800"/>
</div>

## 📦 Installation

### Install via Go
```bash
go install github.com/joao-rrondon/wipeOs@latest
```

### Download Binary
Visit the [releases page](https://github.com/joao-rrondon/wipeOs/releases) and download the appropriate binary for your platform.

### Package Managers

#### Homebrew (macOS/Linux)
```bash
brew tap joao-rrondon/tools
brew install wipeOs
```

#### Debian/Ubuntu
```bash
wget https://github.com/joao-rrondon/wipeOs/releases/latest/download/wipeOs_Linux_x86_64.deb
sudo dpkg -i wipeOs_Linux_x86_64.deb
```

## 📋 **Complete Command Reference**

### 🏠 **Main Commands**

#### **Interactive Mode (Default)**
```bash
# Launch interactive terminal (recommended)
wipeOs                    # Default interactive mode
wipeOs start             # Same as above
wipeOs interactive       # Explicit interactive mode
wipeOs retro            # Quick retro theme + interactive
```

#### **Direct CLI Commands**
```bash
wipeOs <command> [options]    # Direct command execution
wipeOs --help                 # Show all available commands
```

---

## 🗑️ **`wipe` - Secure File Deletion**

**Purpose**: Permanently delete files using military-grade overwriting

### **Basic Usage**
```bash
# Single file
wipeOs wipe secret.txt

# Multiple files
wipeOs wipe file1.txt file2.txt file3.txt

# Pattern matching
wipeOs wipe *.log *.tmp

# Directory (recursive)
wipeOs wipe /path/to/directory --recursive
```

### **Safety Options**
```bash
# Preview mode (ALWAYS use first!)
wipeOs wipe *.sensitive --dry-run

# Force mode (skip confirmations)
wipeOs wipe file.txt --force

# Custom overwrite passes (1-35)
wipeOs wipe file.txt --passes 7
```

### **Predefined Targets**
```bash
# Browser data (cache, history, cookies)
wipeOs wipe --browser-data

# System temporary files
wipeOs wipe --system-temp

# Combine with files
wipeOs wipe secret.txt --browser-data --system-temp
```

### **Examples**
```bash
# Safe testing
wipeOs wipe *.log --dry-run

# Quick single file
wipeOs wipe document.pdf

# Thorough directory wipe
wipeOs wipe sensitive_folder/ --recursive --passes 7

# Emergency browser clean
wipeOs wipe --browser-data --force
```

---

## 🧽 **`clean` - Quick Predefined Cleaning**

**Purpose**: Fast cleanup of common targets without specifying files

### **Available Targets**
```bash
# Individual targets
wipeOs clean browser        # Browser data only
wipeOs clean temp          # System temporary files
wipeOs clean cache         # User cache directories
wipeOs clean logs          # Application logs
wipeOs clean downloads     # Downloads folder (with confirmation)

# Combined operations
wipeOs clean browser temp  # Multiple targets
wipeOs clean all          # Everything (most thorough)
wipeOs clean quick        # Essential cleanup only
```

### **Safety Options**
```bash
# Preview any operation
wipeOs clean all --dry-run
wipeOs clean browser --dry-run

# Skip confirmations
wipeOs clean temp --force

# Custom passes
wipeOs clean all --passes 5
```

### **Examples**
```bash
# Quick browser cleanup
wipeOs clean browser

# Safe full preview
wipeOs clean all --dry-run

# Emergency quick clean
wipeOs clean quick --force

# Thorough cleanup
wipeOs clean all --passes 7
```

---

## 🔍 **`forensic` - Anti-Forensic Operations**

**Purpose**: Military-grade trace removal for high-security scenarios

### **⚠️ DANGER WARNING**
```bash
# ALWAYS test first - these operations are IRREVERSIBLE!
wipeOs forensic --all --dry-run
```

### **Quick Operations**
```bash
# Essential anti-forensic cleanup
wipeOs forensic --quick --dry-run
wipeOs forensic --quick

# Complete trace removal
wipeOs forensic --all --dry-run
wipeOs forensic --all
```

### **Specific Operations**
```bash
# System logs and events
wipeOs forensic --logs --eventlogs --dry-run

# Registry traces
wipeOs forensic --registry --prefetch --dry-run

# Memory and swap
wipeOs forensic --memory --swap --dry-run

# Shadow copies and MFT
wipeOs forensic --shadows --mft --dry-run

# Free space wiping
wipeOs forensic --freespace --passes 7 --dry-run
```

### **Individual Flags**
| Flag | Description |
|------|-------------|
| `--logs` | Clean system and application logs |
| `--registry` | Clean Windows Registry traces |
| `--prefetch` | Clean Windows Prefetch files |
| `--thumbnails` | Clean thumbnails and recent files |
| `--eventlogs` | Clear Windows Event Logs |
| `--mft` | Clean Master File Table records |
| `--shadows` | Delete Volume Shadow Copies |
| `--memory` | Remove memory dump files |
| `--swap` | Clean swap/page files |
| `--freespace` | Wipe free disk space |

### **Examples**
```bash
# Post-operation cleanup
wipeOs forensic --quick --dry-run
wipeOs forensic --quick

# Maximum security
wipeOs forensic --all --passes 7 --dry-run
wipeOs forensic --all --passes 7

# Selective cleanup
wipeOs forensic --logs --registry --shadows --dry-run
```

---

## 🎨 **`icons` - Theme Management**

**Purpose**: Customize visual appearance with different icon themes

### **Available Themes**
| Theme | Style | Description |
|-------|-------|-------------|
| `classic` | 🗑️🧽🔍 | Original emoji style |
| `retro` | ⚡★◆ | 80s/90s gaming vibes |
| `cyber` | ⟨⟩⦿◉ | Cyberpunk hacker aesthetic |
| `military` | ⚔⊗⊙ | Military tactical symbols |
| `minimal` | ×○● | Clean minimal design |
| `matrix` | ⊗◯◉ | Matrix digital rain |
| `neon` | ◢◣◤ | Bright electronic style |

### **Commands**
```bash
# View all themes
wipeOs icons list

# Preview a theme
wipeOs icons preview retro
wipeOs icons preview cyber

# Switch theme
wipeOs icons set retro
wipeOs icons set military

# Show current theme
wipeOs icons current
```

### **Examples**
```bash
# Try different themes
wipeOs icons preview retro
wipeOs icons set retro
wipeOs

# Switch to hacker style
wipeOs icons set cyber
wipeOs interactive
```

---

## 🚀 **`start` / `interactive` - Terminal Session**

**Purpose**: Launch persistent interactive terminal with command history

### **Features**
- ✅ Persistent session with command history
- ✅ Navigate history with ↑/↓ arrows  
- ✅ Real-time command execution
- ✅ Beautiful terminal interface
- ✅ All commands available interactively

### **Usage**
```bash
# Launch interactive mode
wipeOs start
wipeOs interactive
wipeOs              # Default behavior

# Quick retro launch
wipeOs retro        # Sets retro theme + interactive
```

### **Interactive Commands**
```
WipeOs► help                           # Show all commands
WipeOs► wipe test.txt --dry-run       # Safe file deletion
WipeOs► clean browser --dry-run       # Quick cleanup
WipeOs► forensic --quick --dry-run    # Anti-forensic preview
WipeOs► icons set cyber               # Change theme
WipeOs► version                       # Show version
WipeOs► status                        # Session info
WipeOs► clear                         # Clear screen
WipeOs► exit                          # Quit
```

---

## ⚡ **`retro` - Quick Retro Mode**

**Purpose**: Instant 80s/90s gaming theme activation

```bash
# One command to rule them all
wipeOs retro

# Equivalent to:
# wipeOs icons set retro
# wipeOs interactive
```

**Perfect for that nostalgic gaming vibe!** 🕹️

---

## ℹ️ **`version` - Version Information**

**Purpose**: Show detailed version and build information

```bash
# Basic version
wipeOs version

# Detailed build info
wipeOs version --verbose
```

---

## 🛡️ **Safety Guidelines**

### **⚠️ Always Use --dry-run First**
```bash
# SAFE: Preview what would happen
wipeOs wipe *.sensitive --dry-run
wipeOs clean all --dry-run  
wipeOs forensic --all --dry-run

# DANGER: Actual operations (irreversible)
wipeOs wipe *.sensitive
wipeOs clean all
wipeOs forensic --all
```

### **🔒 Security Levels**
1. **`--dry-run`**: SAFE - Shows what would happen
2. **Normal**: Asks for confirmation before dangerous operations
3. **`--force`**: DANGEROUS - Skips all confirmations

### **🎯 Recommended Workflow**
```bash
# 1. Always preview first
wipeOs forensic --all --dry-run

# 2. If satisfied, execute
wipeOs forensic --all

# 3. For emergency situations
wipeOs forensic --quick --force
```

## 🔧 Configuration

### Overwrite Patterns
WipeOs uses multiple overwrite patterns based on the DoD 5220.22-M standard:

1. **Pass 1**: Random data
2. **Pass 2**: All zeros (0x00)
3. **Pass 3**: All ones (0xFF)
4. **Pass 4+**: Alternating patterns

### Supported Browsers
- Google Chrome
- Mozilla Firefox
- Microsoft Edge
- Safari (macOS)

### System Compatibility
- **Linux**: Full support for all features
- **macOS**: Full support including Safari data
- **Windows**: Full support including Edge data

## 📊 Security Standards

WipeOs follows industry-standard secure deletion practices:

- **DoD 5220.22-M**: Department of Defense standard for media sanitization
- **Multiple Passes**: Configurable overwrite passes (default: 3)
- **Pattern Variety**: Different bit patterns to ensure complete data destruction
- **Metadata Clearing**: File metadata and directory entries are properly cleared

## 🏗️ Building from Source

### Prerequisites
- Go 1.22 or higher
- Git

### Build Steps
```bash
git clone https://github.com/joao-rrondon/wipeOs.git
cd wipeOs
go build -o wipeOs main.go
```

### Development
```bash
# Install dependencies
go mod download

# Run tests
go test ./...

# Run with hot reload
go run main.go

# Build for all platforms
goreleaser build --snapshot --clean
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run benchmarks
go test -bench=. ./...
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Run the test suite
6. Submit a pull request

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Important Security Notice

**WipeOs permanently destroys data. This action is IRREVERSIBLE.**

- Always backup important data before using WipeOs
- Use `--dry-run` to preview operations
- Be extremely careful with recursive operations
- Test on non-critical data first

## 🌟 Star History

[![Star History Chart](https://api.star-history.com/svg?repos=joao-rrondon/wipeOs&type=Date)](https://star-history.com/#joao-rrondon/wipeOs&Date)

## 💬 Community

- [GitHub Discussions](https://github.com/joao-rrondon/wipeOs/discussions)
- [Issues](https://github.com/joao-rrondon/wipeOs/issues)
- [Reddit Community](https://reddit.com/r/wipeOs)

---

<div align="center">
  <p>Made with ❤️ by <a href="https://github.com/joao-rrondon">João Rondon</a></p>
  <p>If WipeOs helped you, consider giving it a ⭐!</p>
</div> 