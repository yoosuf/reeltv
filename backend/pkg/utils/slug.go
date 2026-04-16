package utils

import (
	"github.com/gosimple/slug"
)

// GenerateSlug generates a URL-friendly slug from a string
func GenerateSlug(s string) string {
	return slug.Make(s)
}
