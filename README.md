# 🎥 Media Metadata API

A lightweight Go-based API that extracts metadata from uploaded media files (e.g., MP4, MP3) using `ffprobe`. Built with the Gin web framework and designed to demonstrate modular project structure, file handling, and external tool integration.

---

## 🚀 Features

- Upload a media file via HTTP POST
- Extracts metadata such as duration, codec info, resolution, bit rate, etc.
- Returns a clean JSON response
- Organizes uploads in a dedicated `/media` folder
- Modular file structure (handlers, utils, etc.)
- Request size limit (10 MiB) and MIME type validation enforced server-side
- Unit and integration test coverage for handlers, middleware, and ffprobe utilities

---

## 🧰 Tech Stack

- **Language:** Go
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Media Tool:** [FFmpeg / ffprobe](https://ffmpeg.org/)
- **JSON Parsing:** `encoding/json`, `os/exec`

---

## 📂 Project Structure

```
media-metadata-api/
├── main.go              # Entry point
├── handlers/            # HTTP route logic
│   └── extract.go
├── middleware/           # Upload size/type enforcement
├── utils/               # FFprobe helpers
│   └── ffprobe.go
└── media/               # Uploaded or test media files (.gitkeep tracked)
```

---

## 📦 Installation & Setup

1. **Clone the repo**
   ```bash
   git clone https://github.com/Vince33/media-metadata-api.git
   cd media-metadata-api
   ```

2. **Install dependencies**
   Make sure you have Go and ffprobe (part of FFmpeg) installed:
   ```bash
   brew install ffmpeg
   ```

3. **Run the server**
   ```bash
   go run main.go
   ```

Note: 📦 Vendored dependencies included for offline testing. Periodically updated using ```go get -u && go mod vendor```
.

---

## 📫 API Usage

### `POST /extract`

- **Content-Type:** multipart/form-data
- **Body:** media file as `file` key
- **Limits:** 10 MiB max file size; accepted types are `video/mp4`, `video/mpeg`, and `video/quicktime`

#### Example using `curl`:
```bash
curl -X POST http://localhost:8080/extract \
  -F "file=@SampleVideo_1280x720_1mb.mp4"
```

#### Response:
```json
{
  "format": {
    "filename": "media/SampleVideo_1280x720_1mb.mp4",
    "format_name": "mov,mp4,m4a,3gp,3g2,mj2",
    "duration": "5.312000"
  },
  "streams": [
    {
      "codec_name": "h264",
      "codec_type": "video",
      "width": 1280,
      "height": 720
    },
    {
      "codec_name": "aac",
      "codec_type": "audio",
      "channels": 6,
      "sample_rate": "48000"
    }
  ]
}
```

#### Error responses:
```json
// File exceeds 10 MiB limit (413)
{ "error": "File size exceeds 10 MiB limit" }

// Invalid file type (400)
{ "error": "Invalid file type" }
```

---

## 🔒 Security Considerations

- The project currently stores uploads in the `media/` directory.
- Request body size is capped at 10 MiB via `http.MaxBytesReader`, rejecting oversized uploads before they're fully read.
- MIME type is validated after upload, and files of disallowed types are removed immediately.

---

## 🧪 Testing

This API is used as the real backend in [MetaPeek](https://github.com/Vince33/metapeek)'s 
Playwright end-to-end test suite — its GitHub Actions CI checks out this repo, 
starts the server, and runs real upload requests against it, rather than 
mocking the backend. This repo's own test suite (`*_test.go` files alongside 
handlers, middleware, and utils) covers extraction, validation, and upload 
handling directly.

---

## 🔧 TODO

- [ ] Extend to support images (EXIF), PDFs, etc.
- [ ] Dockerize the application

---

## 📜 License

This project is currently unlicensed and is intended for personal portfolio and educational purposes.

---

## ✍️ Author

Built by [Vince Hines](https://github.com/Vince33) — Software Development Engineer in Test (SDET) with a passion for clean code and creative problem solving.