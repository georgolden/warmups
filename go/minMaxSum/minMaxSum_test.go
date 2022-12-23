package minMaxSum

import (
	"testing"
)

func TestMinMaxSum(t *testing.T) {
	input := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	var w1, w2 uint32 = 2063136757, 2744467344
	minS, maxS := MinMaxSum(input)
	if minS != w1 {
		t.Errorf("Expected %v to be equal %v", minS, w1)
	}
	if maxS != w2 {
		t.Errorf("Expected %v to be equal %v", maxS, w2)
	}
}
