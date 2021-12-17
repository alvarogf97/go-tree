package scan

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizePath(t *testing.T) {
	assert := require.New(t)

	t.Run("test_normalize_path", func(t *testing.T) {
		cleanedPath, _ := normalizePath("./")
		assert.True(filepath.IsAbs(cleanedPath))
	})
}

func TestDecomposePath(t *testing.T) {
	assert := require.New(t)

	t.Run("test_decompose_path", func(t *testing.T) {
		expectedRoot := "memes/2022"
		expectedFilename := "quackked_eated_by_jankos"
		expectedExtension := "png"
		path := "memes/2022/quackked_eated_by_jankos.png"
		root, filename, extension := decomposePath(path)

		assert.Equal(expectedRoot, root)
		assert.Equal(expectedFilename, filename)
		assert.Equal(expectedExtension, extension)
	})

}
