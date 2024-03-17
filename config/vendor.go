package config

import "os"

type VendorConfig struct {
	S3AwsConfig S3AwsConfig
}

func GetVendorConfig() *VendorConfig {
	return &VendorConfig{
		S3AwsConfig: GetS3AwsConfig(),
	}
}

type S3AwsConfig struct {
	Id     string
	Secret string
}

func GetS3AwsConfig() S3AwsConfig {
	secret := ""
	id := ""
	if os.Getenv("ENV") == "production" {
		id = os.Getenv("S3_ID")
		secret = os.Getenv("S3_SECRET_KEY")
	} else {
		secret = conf.GetOptionalValue("A33_SECRET_KEY", "")
		id = conf.GetOptionalValue("A33_ID", "")
	}
	return S3AwsConfig{
		Id:     id,
		Secret: secret,
	}

}
