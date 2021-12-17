package scan

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScannerOptionsConstructor(t *testing.T) {
	assert := require.New(t)

	t.Run("test_constructor_success", func(t *testing.T) {
		showHidden := true
		options := NewScannerOptions(showHidden)

		assert.Equal(showHidden, options.showHidden)
	})
}
