package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func ReadFile(input_path string) []byte {
	// Get the runtime environment of the caller function (hence the 1)
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Error getting current file path.")
	}
	dir := filepath.Dir(filename)
	file, err := os.ReadFile(path.Join(dir, input_path))
	if err != nil {
		panic(err)
	}
	return file
}
