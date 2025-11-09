# gofetch

Dead simple system info fetch program written in Go.

## Why?
another skool projekt :(((((((((

## Features
- Simple Tux ASCII art
- Shows: OS, Kernel, CPU, RAM, Uptime
- Fast (compiled Go binary)
- No dependencies
- Works on any Linux with /proc filesystem

## Installation
```bash
# Clone the repo
git clone https://github.com/qzak112/gofetch.git
cd gofetch

# Build
go build -o gofetch gofetch.go

# Install system-wide
sudo cp gofetch /usr/local/bin/
sudo chmod +x /usr/local/bin/gofetch


## Usage

Just type:
```bash
gofetch
```
