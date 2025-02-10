package utility

import (
	"fmt"
	"os"
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
