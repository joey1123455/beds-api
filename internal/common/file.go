package common

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"

	"github.com/google/uuid"
)

var ImageUploadFolder = "fluxapp-uploads"

const (
	MaxImageWidth  = 1400
	MaxImageHeight = 1400
)

func ScaleToFit(imageData image.Image) image.Image {
	bounds := imageData.Bounds()
	if bounds.Dx() > MaxImageWidth || bounds.Dy() > MaxImageHeight {
		resizedImage := imaging.Fit(imageData, int(MaxImageWidth), int(MaxImageHeight), imaging.Lanczos)
		return resizedImage
	} else {
		return imageData
	}
}

func SaveFileToHomeDir(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user's home directory: %v", err)
	}

	fileName := fmt.Sprintf("%s.jpg", uuid.New().String())
	destPath := filepath.Join(homeDir, ImageUploadFolder, fileName)
	fmt.Println(destPath)

	imageData, err := imaging.Decode(file, imaging.AutoOrientation(true))
	if err != nil {
		return "", err
	}

	saveErr := imaging.Save(imageData, destPath, imaging.JPEGQuality(80))
	if saveErr != nil {
		return "", saveErr
	}

	return fileName, nil
}

func SaveQrImageToHomeDir(img image.Image) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Create a directory for QR codes if it doesn't exist
	qrCodeDir := filepath.Join(homeDir, ImageUploadFolder)
	if err := os.MkdirAll(qrCodeDir, os.ModePerm); err != nil {
		return "", err
	}

	// Define the path to save the image
	////nolint
	fileName := fmt.Sprintf("%s.png", uuid.New().String())
	filePath := filepath.Join(qrCodeDir, fileName)

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Encode the image to the file in PNG format
	if err := png.Encode(file, img); err != nil {
		return "", err
	}

	return fileName, nil
}
