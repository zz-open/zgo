package stringx

import (
	"github.com/zzopen/go-saber/internal"
	"testing"
)

func TestIsEmptyString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsEmptyString")

	assert.Equal(true, IsEmptyString(""))
	assert.Equal(false, IsEmptyString("123"))
	assert.Equal(false, IsEmptyString(string(' ')))
}
