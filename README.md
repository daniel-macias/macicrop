# macicrop

`macicrop` is a small command-line tool written in Go for trimming transparent borders from PNG images.

It was created as a simple image manipulation experiment and a way to practice building CLI applications in Go. The main use case was cleaning up exported image layers that contain large transparent empty space around the visible artwork.

## What it does

Given a folder of PNG files, `macicrop` will:

- Scan each PNG image
- Detect the visible (non-transparent) pixels
- Crop away unnecessary transparent borders
- Save the trimmed images into an output folder

If an image is fully transparent, it can generate a minimal `1x1` transparent PNG.

## Example Use Case

Some art programs export image layers using the full canvas size, even when the actual drawing only uses a small portion of the canvas.

`macicrop` helps reduce that empty space automatically.

## Usage

Run the tool with:

    macicrop trim <input_folder> <output_folder>

Example:

    macicrop trim "./exports" "./trimmed"

Windows example:

    macicrop trim "C:\Users\YourName\Downloads\exports" "C:\Users\YourName\Downloads\trimmed"

## Available Flags

### `--overwrite`

Overwrite existing files in the output folder.

### `--suffix _trim`

Append a suffix to output filenames.

Example:

    hero.png -> hero_trim.png

### `--keep-empty`

Keep fully transparent images by outputting a `1x1` transparent PNG.

## Build

Requires Go installed.

Build executable:

    go build -o macicrop.exe ./cmd/macicrop

Run executable:

    ./macicrop.exe trim "./input" "./output"

## Why this project exists

This project was mainly built to:

- Practice Go project structure
- Learn CLI argument parsing
- Work with file systems and paths
- Experiment with image processing in Go
- Build a small real-world utility

## Status

Functional prototype / learning project.

## License

MIT