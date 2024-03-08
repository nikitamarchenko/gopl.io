Exercise 9.6: Measure how the performance of a compute-bound parallel program (see Exercise 8.5) varies with GOMAXPROCS. What is the optimal value on your computer? How many CPUs does your computer have?


Best:
cpu 6,7 with 8 goroutines
```
BenchmarkCalc/calc(4)-4           	    120	  9972125 ns/op
BenchmarkCalc/calc(5)-4           	    126	  9522194 ns/op
BenchmarkCalc/calc(6)-8           	    127	  9369286 ns/op
BenchmarkCalc/calc(7)-8           	    127	  9409806 ns/op
```

```
go test -bench . -cpu 1,2,4,8,16,32,64,128
goos: linux
goarch: amd64
pkg: gopl.io/ch8/ex8.5
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkCalc/calc(1)             	     60	 19249408 ns/op
BenchmarkCalc/calc(1)-2           	     63	 18435800 ns/op
BenchmarkCalc/calc(1)-4           	     64	 18117609 ns/op
BenchmarkCalc/calc(1)-8           	     63	 18858663 ns/op
BenchmarkCalc/calc(1)-16          	     61	 18405766 ns/op
BenchmarkCalc/calc(1)-32          	     64	 18497430 ns/op
BenchmarkCalc/calc(1)-64          	     62	 18670489 ns/op
BenchmarkCalc/calc(1)-128         	     62	 18697638 ns/op
BenchmarkCalc/calc(2)             	    100	 22964250 ns/op
BenchmarkCalc/calc(2)-2           	     92	 11734706 ns/op
BenchmarkCalc/calc(2)-4           	    100	 12745175 ns/op
BenchmarkCalc/calc(2)-8           	     92	 12700461 ns/op
BenchmarkCalc/calc(2)-16          	     93	 11979536 ns/op
BenchmarkCalc/calc(2)-32          	     94	 12708779 ns/op
BenchmarkCalc/calc(2)-64          	     90	 13023240 ns/op
BenchmarkCalc/calc(2)-128         	     88	 12533687 ns/op
BenchmarkCalc/calc(3)             	     96	 24556109 ns/op
BenchmarkCalc/calc(3)-2           	     90	 13248033 ns/op
BenchmarkCalc/calc(3)-4           	    100	 11241136 ns/op
BenchmarkCalc/calc(3)-8           	    100	 11100133 ns/op
BenchmarkCalc/calc(3)-16          	    100	 11140330 ns/op
BenchmarkCalc/calc(3)-32          	     88	 11978569 ns/op
BenchmarkCalc/calc(3)-64          	    100	 12358159 ns/op
BenchmarkCalc/calc(3)-128         	    100	 12083347 ns/op
BenchmarkCalc/calc(4)             	    100	 22825705 ns/op
BenchmarkCalc/calc(4)-2           	     87	 13811919 ns/op
BenchmarkCalc/calc(4)-4           	    120	  9972125 ns/op
BenchmarkCalc/calc(4)-8           	    100	 10931689 ns/op
BenchmarkCalc/calc(4)-16          	    100	 11113162 ns/op
BenchmarkCalc/calc(4)-32          	    100	 10832603 ns/op
BenchmarkCalc/calc(4)-64          	    100	 11717951 ns/op
BenchmarkCalc/calc(4)-128         	     92	 11995302 ns/op
BenchmarkCalc/calc(5)             	     94	 23088584 ns/op
BenchmarkCalc/calc(5)-2           	     85	 14146273 ns/op
BenchmarkCalc/calc(5)-4           	    126	  9522194 ns/op
BenchmarkCalc/calc(5)-8           	    100	 10190130 ns/op
BenchmarkCalc/calc(5)-16          	    100	 10752257 ns/op
BenchmarkCalc/calc(5)-32          	     97	 11120820 ns/op
BenchmarkCalc/calc(5)-64          	    100	 10404652 ns/op
BenchmarkCalc/calc(5)-128         	     96	 11587612 ns/op
BenchmarkCalc/calc(6)             	    100	 24421850 ns/op
BenchmarkCalc/calc(6)-2           	     80	 13673768 ns/op
BenchmarkCalc/calc(6)-4           	    100	 10223380 ns/op
BenchmarkCalc/calc(6)-8           	    127	  9369286 ns/op
BenchmarkCalc/calc(6)-16          	    100	 10128913 ns/op
BenchmarkCalc/calc(6)-32          	    100	 10090436 ns/op
BenchmarkCalc/calc(6)-64          	    100	 10078030 ns/op
BenchmarkCalc/calc(6)-128         	    100	 10087054 ns/op
BenchmarkCalc/calc(7)             	    100	 21632388 ns/op
BenchmarkCalc/calc(7)-2           	     91	 13055217 ns/op
BenchmarkCalc/calc(7)-4           	    100	 10822851 ns/op
BenchmarkCalc/calc(7)-8           	    127	  9409806 ns/op
BenchmarkCalc/calc(7)-16          	    100	 10350271 ns/op
BenchmarkCalc/calc(7)-32          	    100	 10369896 ns/op
BenchmarkCalc/calc(7)-64          	    100	 10114960 ns/op
BenchmarkCalc/calc(7)-128         	    100	 10545929 ns/op
BenchmarkCalc/calc(8)             	    100	 23491241 ns/op
BenchmarkCalc/calc(8)-2           	     91	 12979286 ns/op
BenchmarkCalc/calc(8)-4           	    100	 10450487 ns/op
BenchmarkCalc/calc(8)-8           	    122	 10195633 ns/op
BenchmarkCalc/calc(8)-16          	    100	 11630314 ns/op
BenchmarkCalc/calc(8)-32          	     91	 11117254 ns/op
BenchmarkCalc/calc(8)-64          	     98	 11101305 ns/op
BenchmarkCalc/calc(8)-128         	    100	 11243760 ns/op
BenchmarkCalc/calc(16)            	    100	 24951121 ns/op
BenchmarkCalc/calc(16)-2          	     88	 13386629 ns/op
BenchmarkCalc/calc(16)-4          	    100	 10092958 ns/op
BenchmarkCalc/calc(16)-8          	    122	  9851716 ns/op
BenchmarkCalc/calc(16)-16         	    100	 10885540 ns/op
BenchmarkCalc/calc(16)-32         	    100	 11765517 ns/op
BenchmarkCalc/calc(16)-64         	    100	 11736228 ns/op
BenchmarkCalc/calc(16)-128        	    100	 11575112 ns/op
BenchmarkCalc/calc(24)            	     90	 25510269 ns/op
BenchmarkCalc/calc(24)-2          	     75	 14446575 ns/op
BenchmarkCalc/calc(24)-4          	    100	 10021668 ns/op
BenchmarkCalc/calc(24)-8          	    100	 10258152 ns/op
BenchmarkCalc/calc(24)-16         	    100	 10693923 ns/op
BenchmarkCalc/calc(24)-32         	    100	 11371739 ns/op
BenchmarkCalc/calc(24)-64         	    100	 11206001 ns/op
BenchmarkCalc/calc(24)-128        	    100	 11366936 ns/op
BenchmarkCalc/calc(32)            	     92	 26578546 ns/op
BenchmarkCalc/calc(32)-2          	     72	 14888123 ns/op
BenchmarkCalc/calc(32)-4          	    100	 10448624 ns/op
BenchmarkCalc/calc(32)-8          	    120	  9866543 ns/op
BenchmarkCalc/calc(32)-16         	    100	 11204752 ns/op
BenchmarkCalc/calc(32)-32         	     92	 11143530 ns/op
BenchmarkCalc/calc(32)-64         	    100	 11551952 ns/op
BenchmarkCalc/calc(32)-128        	     92	 11412085 ns/op
BenchmarkCalc/calc(40)            	     78	 23478730 ns/op
BenchmarkCalc/calc(40)-2          	     84	 14431788 ns/op
BenchmarkCalc/calc(40)-4          	    100	 11397064 ns/op
BenchmarkCalc/calc(40)-8          	    100	 10119703 ns/op
BenchmarkCalc/calc(40)-16         	    100	 11138634 ns/op
BenchmarkCalc/calc(40)-32         	     99	 11114303 ns/op
BenchmarkCalc/calc(40)-64         	    100	 11466282 ns/op
BenchmarkCalc/calc(40)-128        	     90	 11541514 ns/op
BenchmarkCalc/calc(48)            	     99	 24214412 ns/op
BenchmarkCalc/calc(48)-2          	     76	 14541646 ns/op
BenchmarkCalc/calc(48)-4          	     96	 10797972 ns/op
BenchmarkCalc/calc(48)-8          	    100	 10054731 ns/op
BenchmarkCalc/calc(48)-16         	    100	 11365242 ns/op
BenchmarkCalc/calc(48)-32         	    100	 11226833 ns/op
BenchmarkCalc/calc(48)-64         	    100	 11158629 ns/op
BenchmarkCalc/calc(48)-128        	    100	 11562179 ns/op
BenchmarkCalc/calc(56)            	    100	 24107693 ns/op
BenchmarkCalc/calc(56)-2          	     86	 15082451 ns/op
BenchmarkCalc/calc(56)-4          	     99	 10865392 ns/op
BenchmarkCalc/calc(56)-8          	    112	 10295069 ns/op
BenchmarkCalc/calc(56)-16         	    100	 11164205 ns/op
BenchmarkCalc/calc(56)-32         	    100	 11198248 ns/op
BenchmarkCalc/calc(56)-64         	    100	 11372373 ns/op
BenchmarkCalc/calc(56)-128        	    100	 11581528 ns/op
BenchmarkCalc/calc(64)            	     97	 23350514 ns/op
BenchmarkCalc/calc(64)-2          	     85	 14633665 ns/op
BenchmarkCalc/calc(64)-4          	    100	 10255805 ns/op
BenchmarkCalc/calc(64)-8          	    100	 10417780 ns/op
BenchmarkCalc/calc(64)-16         	    100	 11140219 ns/op
BenchmarkCalc/calc(64)-32         	    100	 11252764 ns/op
BenchmarkCalc/calc(64)-64         	    100	 11626020 ns/op
BenchmarkCalc/calc(64)-128        	     94	 11673032 ns/op
PASS
ok  	gopl.io/ch8/ex8.5	163.520s
```