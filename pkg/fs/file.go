package fs

import (
	"fmt"
	"strings"
)

// File descriptor
type File struct {
	name  string
	path  string
	level int
}

// Returns tree file representation
func (file File) tree(symbol string) string {
	offset := strings.Repeat("│   ", (file.level - 1))
	return fmt.Sprintf("%s%s── %s\n", offset, symbol, file.name)
}

// Creates a new file
func NewFile(name string, path string, level int) File {
	return File{name: name, path: path, level: level}
}
