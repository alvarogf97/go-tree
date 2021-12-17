package fs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFolderConstructor(t *testing.T) {
	assert := require.New(t)

	t.Run("test_constructor", func(t *testing.T) {
		name := "memes"
		path := "./memes"
		level := 1
		isRoot := false

		folder := NewFolder(name, path, level, isRoot)

		assert.Equal(name, folder.name)
		assert.Equal(path, folder.path)
		assert.Equal(level, folder.level)
		assert.Equal(isRoot, folder.isRoot)
	})
}

func TestFolderTree(t *testing.T) {
	assert := require.New(t)

	t.Run("test_tree_empty_root", func(t *testing.T) {
		name := "memes"
		path := "./memes"
		level := 1
		isRoot := true
		symbol := ":3"

		folder := NewFolder(name, path, level, isRoot)

		assert.Equal("", folder.tree(symbol))
	})

	t.Run("test_tree_empty", func(t *testing.T) {
		name := "memes"
		path := "./memes"
		level := 1
		isRoot := false
		symbol := ":3"

		expectedString := fmt.Sprintf("%s── %s\n", symbol, name)
		folder := NewFolder(name, path, level, isRoot)

		assert.Equal(expectedString, folder.tree(symbol))
	})

	t.Run("test_tree_full", func(t *testing.T) {
		name := "memes"
		path := "./memes"
		level := 1
		isRoot := false

		folder := NewFolder(name, path, level, isRoot)
		subfolder := NewFolder("dark", "./memes/dark", 2, false)
		subsubfile := NewFile("dog_bonks_godzilla.png", "./memes/dark/dog_bonks_godzilla.png", 3)
		subsubfile2 := NewFile("ankara_messi.png", "./memes/dark/ankara_messi.png", 3)
		subfolder2 := NewFolder("light", "./memes/light", 2, false)

		subfolder.AddFile(subsubfile)
		subfolder.AddFile(subsubfile2)
		folder.AddFolder(*subfolder)
		folder.AddFolder(*subfolder2)

		expectedString := fmt.Sprintf("├── %s\n│   ├── %s\n│   │   ├── %s\n│   │   └── %s\n│   └── %s\n", folder.name, subfolder.name, subsubfile.name, subsubfile2.name, subfolder2.name)
		assert.Equal(expectedString, folder.tree("├"))
	})

	t.Run("test_Tree", func(t *testing.T) {
		name := "memes"
		path := "./memes"
		level := 1
		isRoot := false

		expectedString := fmt.Sprintf("── %s\n", name)
		folder := NewFolder(name, path, level, isRoot)

		assert.Equal(expectedString, folder.Tree())
	})
}

func TestFolderOperations(t *testing.T) {
	assert := require.New(t)

	t.Run("test_add_folder", func(t *testing.T) {
		name := "memes"
		path := "./memes"
		level := 1
		isRoot := false

		folder := NewFolder(name, path, level, isRoot)
		subfolder := NewFolder("dark", "./memes/dark", 2, false)

		assert.Equal(0, len(folder.folders))
		folder.AddFolder(*subfolder)
		assert.Equal(1, len(folder.folders))
	})

	t.Run("test_add_file", func(t *testing.T) {
		name := "memes"
		path := "./memes"
		level := 1
		isRoot := false

		folder := NewFolder(name, path, level, isRoot)
		subfile := NewFile("ankara_messi.png", "./memes/ankara_messi.png", 2)

		assert.Equal(0, len(folder.files))
		folder.AddFile(subfile)
		assert.Equal(1, len(folder.files))
	})

	t.Run("test_get_level", func(t *testing.T) {
		name := "memes"
		path := "./memes"
		level := 1
		isRoot := false

		folder := NewFolder(name, path, level, isRoot)

		assert.Equal(level, folder.GetDepth())
	})

	t.Run("test_get_path", func(t *testing.T) {
		name := "memes"
		path := "./memes"
		level := 1
		isRoot := false

		folder := NewFolder(name, path, level, isRoot)

		assert.Equal(path, folder.GetPath())
	})
}
