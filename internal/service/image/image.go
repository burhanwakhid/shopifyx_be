package imageService

import (
	"mime/multipart"

	awss3 "github.com/burhanwakhid/shopifyx_backend/pkg/aws_s3"
)

type ImageService struct {
}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (s *ImageService) UploadImage(form *multipart.FileHeader) (string, error) {
	url, err := awss3.UploadImageToS3(form)

	if err != nil {
		return "", err
	}

	return url, nil
}
