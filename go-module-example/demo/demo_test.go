package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorateString(t *testing.T) {

	argument := "my argument"
	expectedResult := "my argument - by JFrog"
	gotResult := DecorateString(argument)

	assert.Equal(t, expectedResult, gotResult, "Got wrong result!")

}
