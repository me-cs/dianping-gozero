package utils

import (
	"strings"

	"github.com/google/uuid"
)

// GenerateUUID 生成不带横杠的UUID
func GenerateUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}