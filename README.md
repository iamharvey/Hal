# Hal - A lib of sort algorithms

## Provided Sorting Algorithms (From Fast to Slowest In General)
- Tim Sort
- Quick Sort (Support two partition methods: Lomuto And Hoare)
- Heap Sort
- B-Tree Sort
- Radix Sort
- Gnome Sort
- Cocktail Shaker Sort
- Odd–even Sort / Brick Sort
- Bubble Sort
- Insert Sort

## TO-DO
- Support string sort
- Search algorithms

## Benchmark

### Test Bed Description
| Hardware Platform | MacBook Pro (15-inch, 2017) |
| --- | --- |
| OS | macOS 10.14.4 |
| Processor | 2.9 GHz Intel Core i7, 4 Cores | 
| Memory | 16 GB 2133 MHz LPDDR3 | 

### Task
Sorting N Random Numbers, range from 1 to 2N. In my benchmark experiment, N is set to 100, 1000, 10000, 100000.

### Results - 100 Random Numbers
- Fast in speed: Tim Sort(run size=32), 
- Best in general (both speed and memory): Insert Sort

<br>

| Algorithm | Rounds | Running Time | Memory  | Allocations | 
| --------- | --------- | --------- | --------- | --------- | 
**Tim Sort(run size=32)** |             1157804   |          **3102 ns/op** |            896 B/op |          1 allocs/op
Tim Sort(run size=64) |             1000000   |          3147 ns/op |            896 B/op |          1 allocs/op
Tim Sort(run size=128) |            629481   |          5964 ns/op |           2496 B/op |          7 allocs/op
Quick Sort(Hoare) |                 482629   |          7420 ns/op |            896 B/op |          1 allocs/op 
Quick Sort(Lomuto) |                597049   |          5837 ns/op |            896 B/op |          1 allocs/op 
Quick Sort(Golang Built-in) |       384768   |          9351 ns/op |            928 B/op |          2 allocs/op
Insert Sort |                       730724   |          4890 ns/op |            896 B/op |          1 allocs/op

### Results - 1000 Random Numbers
- Fast in speed: Tim Sort(run size=128)
- Best in general (both speed and memory): Quick Sort(Lomuto)

<br>

| Algorithm | Rounds | Running Time | Memory  | Allocations | 
| --------- | --------- | --------- | --------- | --------- | 
Tim Sort(run size=32) |             49072   |          73275 ns/op |           48640 B/op |         63 allocs/op
Tim Sort(run size=64) |             45804   |          76612 ns/op |           48640 B/op |         31 allocs/op
**Tim Sort(run size=128)**|             87736   |          **41032 ns/op** |           32640 B/op |         15 allocs/op
Quick Sort(Hoare) |                 43159   |          85893 ns/op |            8192 B/op |          1 allocs/op 
Quick Sort(Lomuto) |                45590   |          71234 ns/op |            8192 B/op |          1 allocs/op 
Quick Sort(Golang Built-in) |       27972   |         127098 ns/op |            8224 B/op |          2 allocs/op
Insert Sort |                       20487   |         172544 ns/op |            8192 B/op |          1 allocs/op

### Results - 10K Random Numbers
- Fast in speed: Tim Sort(run size=128)
- Best in general (both speed and memory): Quick Sort(Lomuto)

<br>

| Algorithm | Rounds | Running Time | Memory  | Allocations | 
| --------- | --------- | --------- | --------- | --------- | 
Tim Sort(run size=32) |              3559   |         927927 ns/op |         774914 B/op  |      625 allocs/op
Tim Sort(run size=64) |              5282   |         707896 ns/op |         695044 B/op   |     313 allocs/op
**Tim Sort(run size=128)** |             6616     |       **559145 ns/op**  |        615171 B/op   |     157 allocs/op
Quick Sort(Hoare) |                  3795    |        974970 ns/op |          81920 B/op  |        1 allocs/op 
Quick Sort(Lomuto) |                 4334     |       872705 ns/op |          81920 B/op |         1 allocs/op 
Quick Sort(Golang Built-in) |        2158     |      1629768 ns/op |          81952 B/op  |        2 allocs/op
Insert Sort |                         235     |     14666689 ns/op   |        81920 B/op     |     1 allocs/op

### Results - 100K Random Numbers 
- Fast in speed: Tim Sort(run size=128) (0.007 seconds)
- Best in general (both speed and memory): Quick Sort(Lomuto)

<br>

| Algorithm | Rounds | Running Time | Memory Allocation | Allocations | 
| --------- | --------- | --------- | --------- | --------- | 
Tim Sort(run size=32) |              285 |         11812962 ns/op |       10349629 B/op |      6249 allocs/op
Tim Sort(run size=64) |              414 |          8985237 ns/op |        9549867 B/op |      3125 allocs/op
**Tim Sort(run size=128)** |             495 |          **7144921 ns/op** |        8750111 B/op |      1563 allocs/op
Quick Sort(Hoare) |                  344 |         10669909 ns/op |         802817 B/op |         1 allocs/op
Quick Sort(Lomuto) |                 381 |          9409562 ns/op |         802819 B/op |         1 allocs/op
Quick Sort(Golang Built-in) |        194 |         18554115 ns/op |         802848 B/op |         2 allocs/op
Insert Sort |                        3   |       1470262444 ns/op   |       802816 B/op   |       1 allocs/op


### Tim Sort(128) VS Quick Sort(Lomuto) - 100 Million Random Numbers 
- Tim Sort(run size=128): 11.5 seconds, 15GB per operation
- Quick Sort(Lomuto): 14.4 seconds, 760MB per operation

| Algorithm | Rounds | Running Time | Memory Allocation | Allocations | 
| --------- | --------- | --------- | --------- | --------- | 
Tim Sort(run size=128)     |          1 |       11525982358 ns/op  |     16532869296 B/op  |       1562525 allocs/op
Quick Sort(Lomuto)         |          1 |       14375655148 ns/op  |     800006144 B/op    |       1 allocs/op