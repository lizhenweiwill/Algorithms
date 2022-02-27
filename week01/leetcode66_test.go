package week01

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var source []int // 源
	var target []int // 目标
	var result []int // 结果
	source = []int{1, 2, 3}
	target = []int{1, 2, 4}
	result = plusOne(source)
	if !reflect.DeepEqual(target, result) {
		t.Log("error - 1")
	}
	source = []int{9, 9, 9}
	target = []int{1, 0, 0, 0}
	result = plusOne(source)
	if !reflect.DeepEqual(target, result) {
		t.Log("error - 2")
	}
	source = []int{0}
	target = []int{1}
	result = plusOne(source)
	if !reflect.DeepEqual(target, result) {
		t.Log("error - 3")
	}
}
