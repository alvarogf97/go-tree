package scan

import (
	"path/filepath"
	"strings"
)

// Converts the given path to absolute and normalize
// backslash to make it compatible with all systems
func normalizePath(path string) (string, error) {
	return filepath.Abs(filepath.ToSlash(path))
}

// Splits the given path into the root, the filename
// and the file extension
func decomposePath(path string) (string, string, string) {
	splitted := strings.Split(path, "/")
	root := strings.Join(splitted[:len(splitted)-1], "/")
	fullname := splitted[len(splitted)-1]

	splittedFullname := strings.Split(fullname, ".")
	extension := splittedFullname[len(splittedFullname)-1]
	filename := strings.Join(splittedFullname[:len(splittedFullname)-1], ".")
	return root, filename, extension
}
