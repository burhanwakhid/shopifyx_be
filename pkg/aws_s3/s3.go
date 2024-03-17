package awss3

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/burhanwakhid/shopifyx_backend/config"
)

// UploadImageToS3 uploads an image to the specified S3 bucket with validation
func UploadImageToS3(form *multipart.FileHeader) (string, error) {

	bucketName := "sprint-bucket-public-read"

	awsConfig := config.GetS3AwsConfig()
	id := awsConfig.Id
	secret := awsConfig.Secret
	// Validate image size
	if err := validateImageSize(form); err != nil {
		return "", err
	}
	// Get Buffer from file
	buffer, err := form.Open()
	if err != nil {
		return "", err
	}
	defer buffer.Close()

	// Get the file extension
	ext := filepath.Ext(form.Filename) // Assuming file.Filename contains the original file name
	if ext == "" {
		return "", errors.New("missing file extension")
	}

	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			id,
			secret,
			"",
		),
	})
	if err != nil {
		return "", fmt.Errorf("failed to create AWS session: %w", err)
	}

	// Create a new S3 service client
	svc := s3.New(sess)

	// Read the image data from the file
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, buffer)
	if err != nil {
		return "", fmt.Errorf("failed to read image data: %w", err)
	}

	now := time.Now().UTC()
	timestamp := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	// Construct the object key
	objectKey := fmt.Sprintf("%s_%s", timestamp, ext)

	// Create a PutObject input struct
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		ACL:    aws.String("public-read"),
		Body:   bytes.NewReader(buf.Bytes()),
	}

	// Upload the image to S3
	_, err = svc.PutObject(input)
	if err != nil {
		fmt.Print("failed to upload image to S3: %w", err)
		return "", err
	}

	fmt.Println("Image uploaded successfully!: 1")

	// Generate public URL for the uploaded object
	publicURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, "ap-southeast-1", objectKey)

	// Print or use the public URL as needed
	fmt.Println("Public URL:", publicURL)

	fmt.Println("Image uploaded successfully!: 2 ", publicURL)
	return publicURL, nil

}

// validateImage checks image format and size
func validateImageSize(imageFile *multipart.FileHeader) error {
	// Detect content type using mime package
	contentType := imageFile.Header.Get("Content-Type")

	fmt.Print("ini content typppe: %w", contentType)

	// Check allowed formats (modify as needed)
	if contentType != "image/jpeg" && contentType != "image/jpg" {
		return errors.New("invalid image format. Only JPG or JPEG allowed")
	}

	// Check image size
	if imageFile.Size < 10*1024 {
		return errors.New("image size too small. Minimum size is 10KB")
	}

	if imageFile.Size > 2*1024*1024 {
		return errors.New("image size too large. Maximum size is 2MB")
	}

	return nil
}
