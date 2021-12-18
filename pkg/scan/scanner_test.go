package scan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func failNormalizePath(path string) (string, error) {
	return "", fmt.Errorf("error")
}

func TestScannerConstructor(t *testing.T) {
	assert := require.New(t)

	t.Run("test_constructor", func(t *testing.T) {
		showHidden := true
		scanner := NewScanner(NewScannerOptions(showHidden))

		assert.Equal(showHidden, scanner.options.showHidden)
	})
}

func TestScannerScan(t *testing.T) {
	assert := require.New(t)

	t.Run("test_scanner_scan_success", func(t *testing.T) {
		scanner := NewScanner(NewScannerOptions(false))
		folder, err := scanner.Scan("./")

		assert.NoError(err)
		assert.Contains(folder.GetPath(), "go-tree")
	})

	t.Run("test_scanner_scan_fail_normalize_path", func(t *testing.T) {
		scanner := NewScanner(NewScannerOptions(false))
		scanner.normalize = failNormalizePath

		_, err := scanner.Scan("./")

		assert.Error(err)
	})

	t.Run("test_scanner_scan_fail_stat_path", func(t *testing.T) {
		scanner := NewScanner(NewScannerOptions(false))
		_, err := scanner.Scan("fakepath")

		assert.Error(err)
	})

	t.Run("test_scanner_scan_fail_path_is_file", func(t *testing.T) {
		scanner := NewScanner(NewScannerOptions(false))
		_, err := scanner.Scan("./path.go")

		assert.Error(err)
	})

}
