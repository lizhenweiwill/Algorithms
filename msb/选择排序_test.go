package msb

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{7, 1, 3, 5, 1, 6, 8, 3, 5, 7, 5, 6}
	//arr := []int{7, 1, 3, 5, 1}
	fmt.Println(arr)
	SelectSort(arr)
	fmt.Println(arr)
}
