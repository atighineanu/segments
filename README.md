## CONCLUSION:
##### 1. All three testcases (from solution_test.go) are PASSED for all the three solutions.
##### (solution implementation can be found in package utils)
##### GO TEST
```
user@shakdahack:~/scripts/github.com/atighineanu/megaproj1$ go test -v
=== RUN   TestIntervals
--- PASS: TestIntervals (0.00s)
PASS
ok      megaproj1       0.003s
```
##### TESTCASE #2 in solutions_test.go corresponds to the given (challenge) task. As you can see - the test PASSES as well.
	
##### 2. The benchmarks with tiny, small and bigger input(s) show SOLUTIONII to be the fastest, SOLUTIONI - the next fastest, and the slowest is SOLUTIONIII.

GO BENCHMARK
```
user@shakdahack:~/scripts/github.com/atighineanu/megaproj1$ go test -bench=.
goos: linux
goarch: amd64
pkg: megaproj1
BenchmarkForSpeedSolItest0-12            7250754               150 ns/op
BenchmarkForSpeedSolIItest0-12          55425487                21.5 ns/op
BenchmarkForSpeedSolIIItest0-12          3994804               298 ns/op
BenchmarkForSpeedSolItest1-12            1504713               822 ns/op
BenchmarkForSpeedSolIItest1-12          11000529                99.5 ns/op
BenchmarkForSpeedSolIIItest1-12          3235483               392 ns/op
BenchmarkForSpeedSolItest2-12              63642             19251 ns/op
BenchmarkForSpeedSolIItest2-12            225406              4955 ns/op
BenchmarkForSpeedSolIIItest2-12            40290             28735 ns/op
PASS
ok      megaproj1       13.874s
```
###### BUT, with increasing the size of the INPUT, the solution II becomes less and less efficient if compared to solution III and I.
###### Last but not least: solution I and II don't sort the intervals in ascending order, while solution III does. 
