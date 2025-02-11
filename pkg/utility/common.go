package utility

import (
	"fmt"
	"os"
	"strings"
)

func GetDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	fmt.Println("Current working directory:", dir)
	return dir
}

func FirstToUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstToLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
