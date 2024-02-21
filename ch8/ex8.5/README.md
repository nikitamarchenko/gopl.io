
Exercise 8.5: Take an existing CPU-bound sequential program, such as the
Mandelbrot program of Section 3.3 or the 3-D surface computation of
Section 3.2, and execute its main loop in parallel using channels for
communication. How much faster does it run on a multiprocessor machine? What
is the optimal number of goroutines to use?


```
goos: linux
goarch: amd64
pkg: gopl.io/ch8/ex8.5
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkCalc/calc(1)-8         	     56	 18172922 ns/op	11036832 B/op	  90045 allocs/op
BenchmarkCalc/calc(2)-8         	    103	 11959692 ns/op	11037210 B/op	  90050 allocs/op
BenchmarkCalc/calc(3)-8         	    100	 10326202 ns/op	11038801 B/op	  90061 allocs/op
BenchmarkCalc/calc(4)-8         	    100	 10575599 ns/op	11039500 B/op	  90069 allocs/op
BenchmarkCalc/calc(5)-8         	    121	  9789189 ns/op	11039981 B/op	  90075 allocs/op
BenchmarkCalc/calc(6)-8         	    130	  9379015 ns/op	11040756 B/op	  90082 allocs/op
BenchmarkCalc/calc(7)-8         	    126	  9852112 ns/op	11041668 B/op	  90091 allocs/op
BenchmarkCalc/calc(8)-8         	    123	  9612305 ns/op	11042358 B/op	  90098 allocs/op
BenchmarkCalc/calc(16)-8        	    122	  9739102 ns/op	11044309 B/op	  90121 allocs/op
BenchmarkCalc/calc(24)-8        	    100	 12102041 ns/op	11045399 B/op	  90137 allocs/op
BenchmarkCalc/calc(32)-8        	     98	 12139896 ns/op	11046089 B/op	  90150 allocs/op
BenchmarkCalc/calc(40)-8        	     92	 12089333 ns/op	11047429 B/op	  90168 allocs/op
BenchmarkCalc/calc(48)-8        	     99	 12328215 ns/op	11048515 B/op	  90183 allocs/op
BenchmarkCalc/calc(56)-8        	     98	 12310388 ns/op	11049007 B/op	  90193 allocs/op
BenchmarkCalc/calc(64)-8        	     94	 12324857 ns/op	11050429 B/op	  90212 allocs/op
PASS
ok  	gopl.io/ch8/ex8.5	23.183s
```

As expected best is from 6gr to 16gr.
