package array_test

import "testing"

func TestArrayInit(t *testing.T)  {
	 var arr [3]int
	 t.Log(arr[0], arr[1])

	 arr1 := [...]int{1,3,5,6,7}
	 t.Log(arr1)
	 t.Log(len(arr1))
}