// test run
// author: baoqiang
// time: 2020/12/15 8:54 下午
package d1

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDemo(t *testing.T) {
	testGetContainerWithMostWater(t)
}

func testGetContainerWithMostWater(t *testing.T) {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, GetContainerWithMostWater(a), 49, "")
}
