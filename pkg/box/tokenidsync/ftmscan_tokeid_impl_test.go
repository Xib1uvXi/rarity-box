package tokenidsync

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ftmscanTokenIDSynchronizer_Sync(t *testing.T) {
	fds := NewFtmscanTokenIDSynchronizer()

	result, err := fds.Sync("0x6d81145629f154dbf07fdab18d2892818626fecf")
	assert.NoError(t, err)

	assert.Len(t, result, 309)
}
