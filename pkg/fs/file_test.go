package fs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileConstructor(t *testing.T) {
	assert := require.New(t)

	t.Run("test_constructor", func(t *testing.T) {
		name := "dog_bonks_godzilla.png"
		path := "./memes/dog_bonks_godzilla.png"
		level := 1

		file := NewFile(name, path, level)

		assert.Equal(name, file.name)
		assert.Equal(path, file.path)
		assert.Equal(level, file.level)
	})
}

func TestFileTree(t *testing.T) {
	assert := require.New(t)

	t.Run("test_tree", func(t *testing.T) {
		name := "dog_bonks_godzilla.png"
		path := "./memes/dog_bonks_godzilla.png"
		level := 1
		symbol := ":3"

		expectedString := fmt.Sprintf("%s── %s\n", symbol, name)
		file := NewFile(name, path, level)

		assert.Equal(expectedString, file.tree(symbol))
	})
}
