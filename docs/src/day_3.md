# Day 3: Rucksack Reorganization

[Link](https://adventofcode.com/2022/day/3)

Again easy enough problem. BUT GO HAS NO SETS IN ITS STANDARD LIBRARY AARRRHHH.
To be fair, it makes sense, GoLang didn't have generics up until 1.18. Now that
we have generics, I might implement a set datastructure. For now I just abuse
a `map` as a set as follows:

```go
var mySet map[int]bool = make(map[int]bool)

// insert key k
mySet[k] = true

// check if k is in set
if mySet[k] {
    fmt.Println("YAY")
}
```
