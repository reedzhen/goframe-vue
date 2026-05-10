package tools

import "strings"

func TrimLike(value string) string {
	return "%" + strings.TrimSpace(value) + "%"
}

func TrimLikeLeft(value string) string {
	return "%" + strings.TrimSpace(value)
}

func TrimLikeRight(value string) string {
	return strings.TrimSpace(value) + "%"
}
