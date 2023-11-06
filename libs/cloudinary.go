package libs

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type ConfigCloudinary struct {
	CloudName      string
	CloudApiKey    string
	CloudApiSecret string
	CloudFolder    string
}

type CloudinaryService struct {
	cld    *cloudinary.Cloudinary
	folder string
}

func InitCloudinary(opt *ConfigCloudinary) *CloudinaryService {
	cld, err := cloudinary.NewFromParams(opt.CloudName, opt.CloudApiKey, opt.CloudApiSecret)
	if err != nil {
		return nil
	}
	log.Println("Successfully connected to cloudinary")
	return &CloudinaryService{
		cld:    cld,
		folder: opt.CloudFolder,
	}
}

func (service *CloudinaryService) UploadFile(fileHeader *multipart.FileHeader, fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validation Image
	isValid := isValidFile(fileHeader)
	if isValid != nil {
		return "", isValid
	}

	// Convert file
	fileReader, err := convertFile(fileHeader)
	if err != nil {
		return "", err
	}

	// Remove ext
	realFileName := RemoveExtension(fileName)

	// Upload file
	uploadParam, err := service.cld.Upload.Upload(ctx, fileReader, uploader.UploadParams{
		PublicID: realFileName,
		Folder:   service.folder,
	})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}

func convertFile(fileHeader *multipart.FileHeader) (*bytes.Reader, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content into an in-memory buffer
	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, file); err != nil {
		return nil, err
	}

	// Create a bytes.Reader from the buffer
	fileReader := bytes.NewReader(buffer.Bytes())
	return fileReader, nil
}

func RemoveExtension(filename string) string {
	return path.Base(filename[:len(filename)-len(path.Ext(filename))])
}

func isValidFile(fileHeader *multipart.FileHeader) error {
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	maxFileSize := int64(3 * 1024 * 1024) // 3MB

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	// Validasi ekstensi
	validExt := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			validExt = true
			break
		}
	}

	if !validExt {
		return errors.New("file extension is not allowed")
	}

	// Validasi ukuran file
	if fileHeader.Size > maxFileSize {
		return errors.New("file size exceeds 3MB")
	}

	return nil
}
