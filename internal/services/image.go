package services

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/disintegration/imaging"
)

// SaveAndCompressImage reads an uploaded image, resizes it if it exceeds 500x500,
// and saves it to the specified path within the UPLOADS_FOLDER_PATH.
func SaveAndCompressImage(file multipart.File, path string) error {
	uploadsPath := os.Getenv("UPLOADS_FOLDER_PATH")
	if uploadsPath == "" {
		uploadsPath = "./web/static/images" // Fallback
	}

	fullPath := uploadsPath + "/" + path

	// Decode the image
	// file implements io.Reader, which imaging.Decode accepts
	img, err := imaging.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return err
	}

	// Resize if the image is larger than 500px in either dimension.
	// imaging.Fit scales the image down to fit within the specified dimensions
	// while preserving the aspect ratio.
	bounds := img.Bounds()
	if bounds.Dx() > 500 || bounds.Dy() > 500 {
		img = imaging.Fit(img, 500, 500, imaging.Lanczos)
	}

	// Save the image to disk. imaging.Save automatically infers the format
	// (JPEG, PNG) from the filename extension and applies appropriate compression.
	err = imaging.Save(img, fullPath)
	if err != nil {
		fmt.Println("Error saving compressed image:", err)
		return err
	}

	return nil
}
