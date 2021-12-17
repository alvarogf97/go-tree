package shortcuts

import "github.com/alvarogf97/go-tree/pkg/scan"

// Tree folder path shortcut
func Tree(path string, showHidden bool) (string, error) {
	scanner := scan.NewScanner(scan.NewScannerOptions(showHidden))
	folder, err := scanner.Scan(path)
	if err != nil {
		return "", err
	}

	return folder.Tree(), nil
}
