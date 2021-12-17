package scan

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

	t.Run("test_scanner_scan_success", func(t *testing.T) {})

	t.Run("test_scanner_scan_fail_normalize_path", func(t *testing.T) {})

	t.Run("test_scanner_scan_fail_stat_path", func(t *testing.T) {})

	t.Run("test_scanner_scan_fail_path_is_dir", func(t *testing.T) {})

	t.Run("test_scanner_scan_success", func(t *testing.T) {})

}
