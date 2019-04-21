package karatechop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChopFactory(t *testing.T) {
	_, err := GetChopMethod(RecursiveChopType)
	assert.Nil(t, err)

	_, err = GetChopMethod(IterativeChopType)
	assert.Nil(t, err)

	_, err = GetChopMethod(FunctionalChopType)
	assert.Nil(t, err)

	_, err = GetChopMethod(SecondRecursiveChopType)
	assert.Nil(t, err)

	_, err = GetChopMethod(SecondFunctionalChopType)
	assert.Nil(t, err)
}
