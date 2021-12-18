package scan

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	treefs "github.com/alvarogf97/go-tree/pkg/fs"
)

type Scanner struct {
	options   ScannerOptions
	walker    func(*treefs.Folder, bool) func(string, fs.DirEntry, error) error
	normalize func(string) (string, error)
	decompose func(string) (string, string, string)
}

func (scanner Scanner) Scan(path string) (*treefs.Folder, error) {
	// normalize path
	rpath, err := scanner.normalize(path)
	if err != nil {
		return nil, err
	}

	// checks the given path exists
	fileInfo, err := os.Stat(rpath)
	if err != nil {
		return nil, fmt.Errorf("directory `%s` does not exists", rpath)
	}
	if !fileInfo.IsDir() {
		return nil, fmt.Errorf("file cannot be scanned")
	}

	// creates root folder according
	// to the checked path
	_, name, extension := scanner.decompose(rpath)
	root := treefs.NewFolder(fmt.Sprintf("%s.%s", name, extension), rpath, 0, true)

	// scans root folder searching for
	// subfolders and containing files
	filepath.WalkDir(rpath, scanner.walker(root, scanner.options.showHidden))
	return root, nil
}

// Creates a new scanner
func NewScanner(options ScannerOptions) *Scanner {
	return &Scanner{
		options:   options,
		walker:    recursiveWalker,
		normalize: normalizePath,
		decompose: decomposePath,
	}
}
