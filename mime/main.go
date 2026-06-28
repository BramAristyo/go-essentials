package main

import (
	"fmt"
	"mime"
)

func main() {
	mimeTypes := []string{
		// Images
		"image/jpeg",
		"image/png",
		"image/gif",
		"image/webp",
		"image/heic",
		"image/heif",
		"image/avif",
		"image/bmp",
		"image/tiff",
		"image/svg+xml",
		"image/x-icon",

		// Documents
		"application/pdf",
		"text/plain",
		"text/csv",
		"application/json",
		"application/xml",

		// Microsoft Office
		"application/msword",
		"application/vnd.ms-excel",
		"application/vnd.ms-powerpoint",

		// Microsoft Office Open XML
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation",

		// Archives
		"application/zip",
		"application/x-rar-compressed",
		"application/x-7z-compressed",
		"application/gzip",
		"application/x-tar",

		// Audio
		"audio/mpeg",
		"audio/wav",
		"audio/ogg",
		"audio/flac",

		// Video
		"video/mp4",
		"video/x-msvideo",
		"video/x-matroska",
		"video/quicktime",
		"video/webm",
	}

	fmt.Printf("%-80s %s\n", "MIME Type", "Extensions")
	fmt.Println("-------------------------------------------------------------------------------------------------------------------")

	for _, mt := range mimeTypes {
		exts, err := mime.ExtensionsByType(mt)
		if err != nil {
			fmt.Printf("%-80s ERROR: %v\n", mt, err)
			continue
		}

		if len(exts) == 0 {
			fmt.Printf("%-80s (not found)\n", mt)
			continue
		}

		fmt.Printf("%-80s %v\n", mt, exts)
	}
}
