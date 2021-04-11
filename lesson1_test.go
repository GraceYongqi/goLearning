package goLearning

import (
	"fmt"
	"goLearning/utils"
	"testing"
)


func TestCalulateMax(t *testing.T) {
	target := []int {199,23,41,67,34,90,1}
	fmt.Println(utils.CalculateMax(target))
}


