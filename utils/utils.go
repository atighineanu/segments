package utils

import (
	"sort"
)

type Interval struct {
	Begin int
	End   int
}

//--------------SOLUTION I-------------------------------
func MERGEI(intervals []Interval) ([]Interval, error) {
	var mergedIntervals []Interval
	var value Interval
	var candidate bool //---we create a trigger that will memorize if we want to add new element to merged intervals slice or not

	if len(mergedIntervals) == 0 { //---if the slice of merged intervals is empty - we populate it with first element from dataset
		mergedIntervals = append(mergedIntervals, intervals[0])
	}

	for _, value = range intervals {
		for index, _ := range mergedIntervals { //---we keep passing through the newly created slice in internal loop, it is easier to process and compare this way
			if mergedIntervals[index].End < value.Begin || mergedIntervals[index].Begin > value.End { //-- classic condition to check intersection absence
				candidate = true
				for _, val := range mergedIntervals { //--- checking for overlaps in newly populated slice
					if val.Begin <= value.Begin && val.End >= value.End { //-- classic condition to check intersection absence
						candidate = false //--- when overlapping - no need to add new element
					}
				}
			} else {
				if mergedIntervals[index].Begin >= value.Begin {
					mergedIntervals[index].Begin = value.Begin
				}
				if mergedIntervals[index].End <= value.End {
					mergedIntervals[index].End = value.End
				}
				candidate = false //--- when overlapping - no need to add new element
			}
		}
		if candidate {
			mergedIntervals = append(mergedIntervals, value)
			candidate = false //--- after we added new element -> trigger set to 0
		}
	}
	return mergedIntervals, nil
}

//----------------SOLUTION II--------------------------
func RemoveIndex(slice []Interval, index int) []Interval { //-- we created an auxiliary function to remove element "i" from a given array(slice)
	return append(slice[:index], slice[index+1:]...)
}

func MERGEII(intervals []Interval) ([]Interval, error) {
	for j := 0; j < len(intervals)-1; j++ {
		for i := len(intervals) - 1; i > j; i-- { //--- we pass the inner loop from the end of intervals array to the j, because this way we will
			//--- do RemoveIndex easier; if we memorized a lower index "i" and removed it, the next memorized index
			//--- on the shorter intervals slice would have the index "i" (not i+1), that will make us write lots of checks and ifs in order to make
			//--- it work; thus, trimming elements from a HIGHER slice index is just easier.
			if (intervals[j].End < intervals[i].Begin) || (intervals[j].Begin > intervals[i].End) { //-- classic condition to check intersection absence
			} else {
				if intervals[j].Begin >= intervals[i].Begin {
					intervals[j].Begin = intervals[i].Begin
				}
				if intervals[j].End <= intervals[i].End {
					intervals[j].End = intervals[i].End
				}
				intervals = RemoveIndex(intervals, i)
				if i < len(intervals)-1 {
					i++ //--- since we removed the index i, we would like to compare j to i+1 as well, -> to do that we will freeze at i for another cycle
				} //--- because after removal of i, now i+1 has the value of i
			}
		}
	}
	return intervals, nil
}

//------------------SOLUTION III--------------------
func MERGEIII(intervals []Interval) ([]Interval, error) {
	sort.SliceStable(intervals, func(i, j int) bool { //----- if we sort by element.Begin, it will be easier (and clearer) then to write
		return intervals[i].Begin < intervals[j].Begin //-----  proper conditions for a merged segments slice.
	})
	for i := 0; i < len(intervals)-1; i++ {
		if !(intervals[i].End < intervals[i+1].Begin) { //--- because this slice is ordered as f(elem.Begin) -> it's enough to just compare
			if intervals[i].End < intervals[i+1].End { //--- element i to the next one. if overlapping -> take new values and destroy i+1
				intervals[i].End = intervals[i+1].End
			}
			intervals = RemoveIndex(intervals, i+1)
			i -= 1 //--- since we removed the index i+1, we would like to compare i to i+2 as well, -> to do that we will freeze at i for another cycle
		}
	}
	return intervals, nil
}

//====>>>>>>>>>>>> CONCLUSION:
//			1. All three testcases (from solution_test.go) are PASSED for all the three solutions.
//			2. The benchmarks with tiny, small and bigger input(s) show SOLUTIONII to be the fastest,
//			   SOLUTIONI - the next fastest, and the slowest is SOLUTIONIII.
//			user@shakdahack:~/scripts/github.com/atighineanu/megaproj1$ go test -v
//			=== RUN   TestIntervals
//			--- PASS: TestIntervals (0.00s)
//			PASS
//			ok      megaproj1       0.003s
//			user@shakdahack:~/scripts/github.com/atighineanu/megaproj1$ go test -bench=.
//			goos: linux
//			goarch: amd64
//			pkg: megaproj1
//			BenchmarkForSpeedSolItest0-12            7250754               150 ns/op
//			BenchmarkForSpeedSolIItest0-12          55425487                21.5 ns/op
//			BenchmarkForSpeedSolIIItest0-12          3994804               298 ns/op
//			BenchmarkForSpeedSolItest1-12            1504713               822 ns/op
//			BenchmarkForSpeedSolIItest1-12          11000529                99.5 ns/op
//			BenchmarkForSpeedSolIIItest1-12          3235483               392 ns/op
//			BenchmarkForSpeedSolItest2-12              63642             19251 ns/op
//			BenchmarkForSpeedSolIItest2-12            225406              4955 ns/op
//			BenchmarkForSpeedSolIIItest2-12            40290             28735 ns/op
//			PASS
//			ok      megaproj1       13.874s
