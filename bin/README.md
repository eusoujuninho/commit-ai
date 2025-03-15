# Pre-compiled Binaries for Commit-AI

This directory contains pre-compiled binaries of Commit-AI for different operating systems, making it easier to use without having to build from source.

## Available Binaries

- **Linux**: `commit-ai-linux-amd64`
- **macOS**: `commit-ai-darwin-amd64`
- **Windows**: `commit-ai-windows-amd64.exe`

## Verifying Checksums

For security, you should verify the integrity of the binaries using the provided checksums file:

```bash
# On Linux/macOS
cd bin/
sha256sum -c checksums.txt

# On Windows (PowerShell)
Get-FileHash commit-ai-windows-amd64.exe -Algorithm SHA256 | Format-List
# Then manually compare with the value in checksums.txt
```

## Installation Instructions

### Linux

```bash
# Make the binary executable
chmod +x commit-ai-linux-amd64

# Option 1: Move to a directory in your PATH (requires admin privileges)
sudo mv commit-ai-linux-amd64 /usr/local/bin/commit-ai

# Option 2: Move to a user bin directory (no admin privileges required)
mkdir -p ~/bin
mv commit-ai-linux-amd64 ~/bin/commit-ai
# Make sure ~/bin is in your PATH
```

### macOS

```bash
# Make the binary executable
chmod +x commit-ai-darwin-amd64

# Option 1: Move to a directory in your PATH (requires admin privileges)
sudo mv commit-ai-darwin-amd64 /usr/local/bin/commit-ai

# Option 2: Move to a user bin directory (no admin privileges required)
mkdir -p ~/bin
mv commit-ai-darwin-amd64 ~/bin/commit-ai
# Make sure ~/bin is in your PATH
```

### Windows

```
# Simply move the .exe file to a directory in your PATH
# Or create a shortcut to it on your desktop
```

## Usage

Once installed, you can use Commit-AI from any Git repository:

```bash
# First-time setup
commit-ai --configure

# Regular usage
commit-ai
```

For more detailed instructions, see the main [README.md](../README.md) and [DOCUMENTATION.md](../DOCUMENTATION.md) files. 