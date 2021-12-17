package scan

import (
	"io/fs"
	"path/filepath"
	"strings"

	treefs "github.com/alvarogf97/go-tree/pkg/fs"
)

// Retruns an fs.WalkDir function that traverse
// the given folder and fill it with the found
// subfolders and files in a recursive way
func recursiveWalker(root *treefs.Folder, showHidden bool) func(string, fs.DirEntry, error) error {
	return func(path string, di fs.DirEntry, err error) error {
		path = filepath.ToSlash(path)
		splitted := strings.Split(path, "/")

		// if showHidden is false, hidden folder will not
		// be scanned
		if !showHidden && strings.HasPrefix(di.Name(), ".") {
			return nil
		}

		if root.GetPath() == strings.Join(splitted[:len(splitted)-1], "/") {
			if di.IsDir() {
				folder := treefs.NewFolder(di.Name(), path, root.GetDepth()+1, false)
				filepath.WalkDir(path, recursiveWalker(folder, showHidden))
				root.AddFolder(*folder)
			} else {
				root.AddFile(treefs.NewFile(di.Name(), path, root.GetDepth()+1))
			}
		}
		return nil
	}
}
