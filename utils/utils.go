package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func getHomePath() string {
	osys := runtime.GOOS
	if osys == "windows" {
		return os.Getenv("HOMEPATH")
	}
	return os.Getenv("HOME")
}

// Constants
var (
	BroConfPath = filepath.Join(getHomePath(), ".brocli")
)

// DoesExist checks if the given path exists
func DoesExist(pth string) bool {
	if _, err := os.Stat(pth); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsAbsolute checks if the given path is absolute or not
func IsAbsolute(pth string) bool {
	return path.IsAbs(pth)
}

// IsDir checks if the path is a directory
func IsDir(pth string) bool {
	// Assuming it exists
	stat, _ := os.Stat(pth)
	if stat.IsDir() {
		return true
	}
	return false
}
