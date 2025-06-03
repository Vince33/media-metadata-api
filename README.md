# Media Metadata API

A lightweight, fast Go-based API for extracting metadata from media files using FFmpeg/FFprobe.

## 🚀 Current Features

- **Upload Media Files:** Accepts video files via a `POST /extract` endpoint.
- **Extract Metadata:** Uses `ffprobe` to parse format and stream-level metadata.
- **Clean JSON Response:** Returns structured metadata including:
  - File format, duration, and bit rate
  - Video codec, resolution, and stream info
  - Audio stream details (if present in the file)
- **Tested With:** `.mp4` files only (more formats to be tested)


## Project Structure
```
media-metadata-api/
├── main.go              # Entry point
├── handlers/            # HTTP route logic
├── utils/               # FFmpeg/FFprobe helpers
└── media/               # Uploaded or test media
```

## Status

🚧 Early development — core file upload and extraction coming soon.

## Getting Started

```bash
go run main.go
``` 
Then visit: http://localhost:8080/health