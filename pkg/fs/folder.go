package fs

import (
	"fmt"
	"strings"
)

const (
	CONTINUE_TREE_SYMBOL = "├"
	BREAK_TREE_SYMBOL    = "└"
)

// Folder descriptor
type Folder struct {
	name    string
	path    string
	level   int
	isRoot  bool
	folders []Folder
	files   []File
}

// Adds a folder to the current one
func (folder *Folder) AddFolder(f Folder) {
	folder.folders = append(folder.folders, f)
}

// Adds file to the current folder
func (folder *Folder) AddFile(f File) {
	folder.files = append(folder.files, f)
}

// Retrieves folder's path
func (folder Folder) GetPath() string {
	return folder.path
}

// Retrieves folder's depth level
func (folder Folder) GetDepth() int {
	return folder.level
}

// Returns tree folder representation. Folders
// will be printed before files
func (folder Folder) tree(symbol string) string {

	var formattedFiles []string
	var formattedFolders []string

	totalFiles := len(folder.files) - 1
	totalFolders := len(folder.folders) - 1

	// format subfolders
	for i, file := range folder.files {
		if i != totalFiles {
			formattedFiles = append(formattedFiles, file.tree("├"))
		} else {
			formattedFiles = append(formattedFiles, file.tree("└"))
		}
	}

	// format subfiles
	for i, folder := range folder.folders {
		if i != totalFolders || totalFiles >= 0 {
			formattedFolders = append(formattedFolders, folder.tree("├"))
		} else {
			formattedFolders = append(formattedFolders, folder.tree("└"))
		}
	}

	if folder.isRoot {
		return fmt.Sprintf("%s%s", strings.Join(formattedFolders, ""), strings.Join(formattedFiles, ""))
	} else {
		offset := strings.Repeat("│   ", (folder.level - 1))
		return fmt.Sprintf("%s%s── %s\n%s%s", offset, symbol, folder.name, strings.Join(formattedFolders, ""), strings.Join(formattedFiles, ""))
	}
}

// Returns tree folder representation without
// giving the initial recursion symbol
func (folder Folder) Tree() string {
	return folder.tree("")
}

// Creates a new folder
func NewFolder(name string, path string, level int, isRoot bool) *Folder {
	return &Folder{name: name, path: path, level: level, isRoot: isRoot}
}
