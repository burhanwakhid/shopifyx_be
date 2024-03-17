package uuid

import (
	"github.com/google/uuid"
)

func GenerateV4() string {
	val := uuid.New()
	return val.String()
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
