package model_test

import (
	"testing"

	"github.com/sir-farfan/mb2cue/model"
	"github.com/stretchr/testify/assert"
)

func Test_FormatString(t *testing.T) {
	index := model.FormatIndex(102693)
	assert.Equal(t, "01:42:69", index)

	index = model.FormatIndex(307000)
	assert.Equal(t, "05:07:00", index)
}
