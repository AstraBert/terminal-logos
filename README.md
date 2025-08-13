# terminal-logos

CLI tool to print beautiful and colorful ASCII-art logos to your terminal.

## Installation

```bash
go install github.com/AstraBert/terminal-logos
```

## Usage

**terminal-logos create**

```text
Create a logo with the specified phrase (-p, --phrase flag), color (-c, --color flag) and font (-f, --font flag). Example usage: `terminal-logos create --phrase 'Hello Logos' --color 'cyan' --font 'roman'`

Usage:
  terminal-logos create [flags]

Aliases:
  create, c

Flags:
  -c, --color string    The color to print the logo with (default "white")
  -f, --font string     The font to print the logo with (default "alphabet")
  -h, --help            help for create
  -p, --phrase string   The phrase to be printed as logo
```