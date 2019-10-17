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
| Algorithm | Rounds | Running Time | Memory Allocation | Distinct Memory Allocations | 
| --------- | --------- | --------- | --------- | --------- | 
Tim Sort(run size=32) |             553020   |          5897 ns/op |            896 B/op |          1 allocs/op
Tim Sort(run size=64) |             591276   |          5604 ns/op |            896 B/op |          1 allocs/op
Tim Sort(run size=128) |            611834   |          5964 ns/op |           2496 B/op |          7 allocs/op
Quick Sort(Hoare) |                 482629   |          7420 ns/op |            896 B/op |          1 allocs/op 
Quick Sort(Lomuto) |                597049   |          5837 ns/op |            896 B/op |          1 allocs/op 
Quick Sort(Golang Built-in) |       384768   |          9351 ns/op |            928 B/op |          2 allocs/op
Insert Sort |                       730724   |          4890 ns/op |            896 B/op |          1 allocs/op

### Results - 1000 Random Numbers
| Algorithm | Rounds | Running Time | Memory Allocation | Distinct Memory Allocations | 
| --------- | --------- | --------- | --------- | --------- | 
Tim Sort(run size=32) |             49072   |          73275 ns/op |           48640 B/op |         63 allocs/op
Tim Sort(run size=64) |             45804   |          76612 ns/op |           48640 B/op |         31 allocs/op
Tim Sort(run size=128)|             49656   |          72081 ns/op |           48640 B/op |         63 allocs/op
Quick Sort(Hoare) |                 43159   |          85893 ns/op |            8192 B/op |          1 allocs/op 
Quick Sort(Lomuto) |                45590   |          71234 ns/op |            8192 B/op |          1 allocs/op 
Quick Sort(Golang Built-in) |       27972   |         127098 ns/op |            8224 B/op |          2 allocs/op
Insert Sort |                       20487   |         172544 ns/op |            8192 B/op |          1 allocs/op

### Results - 10000 Random Numbers 


### Results - 100000 Random Numbers 