package testify_learning

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTestifyMethod(t *testing.T) {
	//break the test immediately when an assertion fails
	assert := require.New(t)
	//assert := assert.New(t)

	//asserts that two objects are equal.
	assert.Equal(1, 1, "1 == 1")

	assert.Empty(struct{}{}, "判断是否为空")

	//judge the Comparison func
	assert.Conditionf(func() (success bool) {
		success = true
		return
	}, "%s\n", "assert.Conditionf")

	//assert the list1 is equal to the list2
	assert.ElementsMatch([]int{1, 2, 3}, []int{2, 1, 3}, "assert.ElementsMatch")

	//asserts that the specified object is empty. 零值和长度为0的切片和通道都包括
	c := make(chan (int), 1)
	assert.Empty(c, "assert.Empty")

}
