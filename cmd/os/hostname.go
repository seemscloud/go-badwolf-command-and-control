package os

import (
	"fmt"
	"os"
)

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Failed to get hostname")
	}
	return name
}
