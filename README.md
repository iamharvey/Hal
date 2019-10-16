# Hal - A lib of sort algorithms

## Done (From Fast to Slowest)
- Tim Sort
- Quick Sort (Support two partition methods: Lomuto And Hoare)
- Heap Sort
- B-Tree Sort
- Radix Sort
- Gnome Sort
- Cocktail Shaker Sort
- Odd–even Sort / Brick Sort
- Bubble Sort

## 
- Platform: MacBook Pro (15-inch, 2017)
- Processor: 2.9 GHz Intel Core i7, 4 Cores
- 16 GB 2133 MHz LPDDR3


### 1000 Random Numbers Sorting
| Algorithm | Rounds | Running Time | Memory Allocation | Distinct Memory Allocations | 
| --------- | --------- | --------- | --------- | --------- | 
QuickSort(Hoare) |                43159   |          85893 ns/ op |            8192 B/ op |          1 allocs/ op 
QuickSort(Lomuto) |               45590   |          71234 ns/ op |            8192 B/ op |          1 allocs/ op 
QuickSort(Golang Built-in) |      27972   |         127098 ns/ op |            8224 B/ op |          2 allocs/ op 