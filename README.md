# Media Metadata API

A lightweight, fast Go-based API for extracting metadata from media files using FFmpeg/FFprobe.

## Features

- Upload media files (audio/video)
- Extract technical metadata using FFprobe
- Return results as JSON
- Built with [Gin](https://github.com/gin-gonic/gin) and designed for easy deployment

## Project Structure

media-metadata-api/
|----- main.go # Entry point
|----- handlers/ #HTTP route logic
|----- utils/ # FFmpeg/FFprobe helpers
|----- media/ # Uploaded or test media

## Status

🚧 Early development — core file upload and extraction coming soon.

## Getting Started

```bash
go run main.go
``` 
Then visit: http://localhost:8080/health