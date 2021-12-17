package shortcuts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTree(t *testing.T) {
	assert := require.New(t)
	t.Run("test_tree_success", func(t *testing.T) {
		_, err := Tree("./", false)
		assert.NoError(err)
	})

	t.Run("test_tree_fail", func(t *testing.T) {
		_, err := Tree("fakepath", false)
		assert.Error(err)
	})

}
