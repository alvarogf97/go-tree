package scan

import (
	"io/fs"
	"path/filepath"
	"testing"

	treefs "github.com/alvarogf97/go-tree/pkg/fs"
	"github.com/stretchr/testify/require"
)

type fDirentry struct {
	name  string
	isDir bool
}

func (di fDirentry) Name() string {
	return di.name
}
func (di fDirentry) IsDir() bool {
	return di.isDir
}
func (di fDirentry) Type() fs.FileMode {
	return 0
}
func (di fDirentry) Info() (fs.FileInfo, error) {
	return nil, nil
}

func TestRecursiveWalk(t *testing.T) {
	assert := require.New(t)

	t.Run("test_recursive_walker_success_with_dir", func(t *testing.T) {
		path, _ := filepath.Abs("./")
		gpath, _ := filepath.Abs("../fs")
		folder := treefs.NewFolder("scan", path, 0, true)
		di := fDirentry{"pkg", true}
		fn := recursiveWalker(folder, false)

		err := fn(gpath, di, nil)

		assert.NoError(err)
	})

	t.Run("test_recursive_walker_success_with_file", func(t *testing.T) {
		path, _ := filepath.Abs("./")
		gpath, _ := filepath.Abs("./path.go")
		folder := treefs.NewFolder("scan", path, 0, true)
		di := fDirentry{"path.go", true}
		fn := recursiveWalker(folder, false)

		err := fn(gpath, di, nil)

		assert.NoError(err)
	})

	t.Run("test_recursive_walker_success_hidden", func(t *testing.T) {
		path, _ := filepath.Abs("./")
		gpath, _ := filepath.Abs("../../.travis.yml")
		folder := treefs.NewFolder("scan", path, 0, true)
		di := fDirentry{".travis.yml", false}
		fn := recursiveWalker(folder, false)

		err := fn(gpath, di, nil)

		assert.NoError(err)
	})

}
