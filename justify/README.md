# ASCII Art Justify

A Go program that generates ASCII art from text with various alignment options, we worked on the alignment for center, left, right and justify.

## Features

- Generate ASCII art from text using different banner styles
- Support for left, center, right, and justify alignments
- Default banner is "standard" if not specified to run as the normal ascii-art

## Usage

```bash
go run . [--align=<type>] <text> [<banner>]
```

### Parameters

- `--align=<type>`: Alignment type (optional, default: left)
  - `left`: Left align (default)
  - `center`: Center align
  - `right`: Right align
  - `justify`: Justify align (distributes spaces between words)
- `<text>`: The text to convert to ASCII art
- `<banner>`: The banner file name (optional, default: standard)

### Available Banners

- `standard`
- `shadow`
- `thinkertoy`

## Examples

```bash
# Basic usage with default banner and alignment
go run . hello

# Specify alignment
go run . --align=center "Hello World"

# Specify banner
go run . hello shadow

# Full command
go run . --align=justify "Hello World" standard
```

## Installation

Clone the repository and run from the project directory.

```bash
git clone <repository-url>
cd ascii-art-justify
go run . hello
```

## Testing

Run the tests with:

```bash
go test ./alignment
```