package main

import (
	"reflect"
	"testing"

	"megaproj1/utils"
)

func TestIntervals(t *testing.T) {
	//testcase#1
	test1, _ := utils.MERGEI([]utils.Interval{{-20, 0}, {300, 455}, {1, 2}, {5, 10}, {4, 23}, {4, 54}, {55, 100}, {102, 200}})
	test2, _ := utils.MERGEII([]utils.Interval{{-20, 0}, {300, 455}, {1, 2}, {5, 10}, {4, 23}, {4, 54}, {55, 100}, {102, 200}})
	test5, _ := utils.MERGEIII([]utils.Interval{{-20, 0}, {300, 455}, {1, 2}, {5, 10}, {4, 23}, {4, 54}, {55, 100}, {102, 200}})
	check1 := []utils.Interval{{-20, 0}, {300, 455}, {1, 2}, {4, 54}, {55, 100}, {102, 200}}
	check2 := []utils.Interval{{-20, 0}, {300, 455}, {1, 2}, {4, 54}, {102, 200}, {55, 100}}
	check3 := []utils.Interval{{-20, 0}, {1, 2}, {4, 54}, {300, 455}, {102, 200}, {55, 100}}
	check5 := []utils.Interval{{-20, 0}, {1, 2}, {4, 54}, {55, 100}, {102, 200}, {300, 455}}
	if !reflect.DeepEqual(test1, check1) && !reflect.DeepEqual(test1, check2) {
		t.Errorf("Interval overlapping was incorrect, got: %v, want: %v, or: %v\n", test1, check1, check2)
	}
	if !reflect.DeepEqual(test2, check1) && !reflect.DeepEqual(test2, check2) {
		t.Errorf("Interval overlapping was incorrect, got: %v, want: %v or: %v\n", test2, check1, check2)
	}
	if !reflect.DeepEqual(test5, check1) && !reflect.DeepEqual(test5, check2) && !reflect.DeepEqual(test5, check3) && !reflect.DeepEqual(test5, check5) {
		t.Errorf("Interval overlapping was incorrect, got: %v, want: %v or: %v or: %v or %v\n", test5, check1, check2, check3, check5)
	}

	//testcase#2
	check4 := []utils.Interval{{25, 30}, {2, 23}}
	check6 := []utils.Interval{{2, 23}, {25, 30}}
	test3, _ := utils.MERGEI([]utils.Interval{{25, 30}, {2, 19}, {14, 23}, {4, 8}})
	test4, _ := utils.MERGEII([]utils.Interval{{25, 30}, {2, 19}, {14, 23}, {4, 8}})
	test6, _ := utils.MERGEIII([]utils.Interval{{25, 30}, {2, 19}, {14, 23}, {4, 8}})
	if !reflect.DeepEqual(test3, check4) {
		t.Errorf("Interval overlapping was incorrect, got: %v, want: %v.\n", test3, check4)
	}
	if !reflect.DeepEqual(test4, check4) {
		t.Errorf("Interval overlapping was incorrect, got: %v, want: %v.\n", test4, check4)
	}
	if !reflect.DeepEqual(test6, check4) && !reflect.DeepEqual(test6, check6) {
		t.Errorf("Interval overlapping was incorrect, got: %v, want: %v, or: %v\n", test6, check4, check6)
	}

	//testcase#3
	bigslice := []utils.Interval{{-100, -30}, {-99, -50}, {-95, -40}, {0, 1}, {-99, -49}, {-97, -91}, {-70, -49}, {-69, -61}, {-67, -62}, {-65, -63}, {-20, 0}, {-19, -10}, {-17, -15}, {-15, -13}, {-14, -11}, {300, 455}, {1, 2}, {5, 10}, {4, 23}, {4, 54}, {55, 100}, {102, 200}}
	check7 := []utils.Interval{{-100, -30}, {-20, 2}, {300, 455}, {4, 54}, {55, 100}, {102, 200}}
	check8 := []utils.Interval{{-100, -30}, {-20, 2}, {4, 54}, {55, 100}, {102, 200}, {300, 455}}
	test7, _ := utils.MERGEI(bigslice)
	test8, _ := utils.MERGEII(bigslice)
	test9, _ := utils.MERGEIII(bigslice)
	if !reflect.DeepEqual(test7, check7) && !reflect.DeepEqual(test7, check8) {
		t.Errorf("Interval overlapping was incorrect, got: %v, want: %v, or: %v\n", test6, check4, check6)
	}
	if !reflect.DeepEqual(test8, check7) && !reflect.DeepEqual(test8, check8) {
		t.Errorf("Interval overlapping was incorrect, got: %v, want: %v, or: %v\n", test6, check4, check6)
	}
	if !reflect.DeepEqual(test9, check7) && !reflect.DeepEqual(test9, check8) {
		t.Errorf("Interval overlapping was incorrect, got: %v, want: %v, or: %v\n", test6, check4, check6)
	}
}
