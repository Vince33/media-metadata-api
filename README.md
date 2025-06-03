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

## 📌 Project Status

This project is in early development.

✅ Basic upload and metadata extraction from `.mp4` files is working  
🚧 Support for additional media types (e.g., `.mp3`, `.wav`, `.mov`) is planned  
🧪 Only tested with `.mp4` files so far  
📂 Internal structure and utility functions are still evolving  


## Getting Started

```bash
go run main.go
``` 
Then visit: http://localhost:8080/health