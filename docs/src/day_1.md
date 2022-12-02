# Day 1: Calorie Counting

[Link](https://adventofcode.com/2022/day/1)

The first day's problem is easy, as expected. For this problem, I think the 
code will be self explanatory.

## Input handling:
```go
// elf_inventory.go
type ElfInventory struct {
    items []int
    total int
}

func invs_from_string(input string) []ElfInventory {
    inv_lists := strings.Split(input, "\n\n")
    var result []ElfInventory

    for _, inv_list_str := range inv_lists {
        var inv ElfInventory
        inv_list := strings.Split(inv_list_str, "\n")

        for _, inv_amt_str := range inv_list {
            var amt int
            fmt.Sscan(inv_amt_str, &amt)
            inv.items = append(inv.items, amt)
            inv.total += amt
        }
        
        result = append(result, inv)
    }

    return result
}
```
```go
// main.go
func main() {
    input_bytes, err := ioutil.ReadFile("input")
    if err != nil {
        log.Fatal("Failed to read input file.")
    }

    input := string(input_bytes)
    inventories := invs_from_string(input)

    // ...
}
```

## Part 1:
```go
func part1(inventories []ElfInventory) int {
    var max int = -1
    for _, inv := range inventories {
        if inv.total > max {
            max = inv.total
        }
    }

    return max
}
```

## Part 2:
```go
// approach 1: Sort and add top 3 (lazy man's approach)
func part2a(inventories []ElfInventory) int {
    sort.Slice(inventories, func(i, j int) bool {
        return inventories[i].total > inventories[j].total
    })

    return inventories[0].total + inventories[1].total + inventories[2].total
}
```

```go
// approach 2: Three linear scans
func part2b(inventories []ElfInventory) int {
    var top1 int = -1
    var top2 int = -1
    var top3 int = -1
    var t1, t2 int

    for i, inv := range inventories {
        if inv.total > top1 {
            top1 = inv.total
            t1 = i
        }
    }

    for i, inv := range inventories {
        if i != t1 && inv.total > top2 {
            top2 = inv.total
            t2 = i
        }
    }

    for i, inv := range inventories {
        if i != t1 && i != t2 && inv.total > top3 {
            top3 = inv.total
        }
    }

    return top1 + top2 + top3
}
```

```go
// approach 3: Single linear scan
func part2c(inventories []ElfInventory) int {
    var top1 int = -1
    var top2 int = -1
    var top3 int = -1

    for _, inv := range inventories {
        if inv.total > top1 {
            top3 = top2
            top2 = top1
            top1 = inv.total
            continue
        }

        if inv.total > top2 {
            top3 = top2
            top2 = inv.total
            continue
        }

        if inv.total > top3 {
            top3 = inv.total
        }
    }

    return top1 + top2 + top3
}

```

### Benchmarking 
Since I ended up with three approaches for part 2, I ended up benchmarking them
against eachother (using golang's built in benchmarking utility in the `testing`
module). Here are the results:
```
$ go test -run=Bench -bench=.
goos: linux
goarch: amd64
pkg: day1
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkPart1-8        13796878                85.93 ns/op
BenchmarkPart2a-8        1570696               768.9 ns/op
BenchmarkPart2b-8        2083321               578.6 ns/op
BenchmarkPart2c-8        5433682               225.5 ns/op
PASS
ok      day1    6.521s
```

### Performance meme 2.0
I wondered how cache friendly having an array of numbers and the total in the 
struct `ElfInventory` was. So, I ended up tweeking the program to simply use
an array of totals. The benchmark results are as follows:
```
$ go test -run=Bench -bench=.
goos: linux
goarch: amd64
pkg: day1
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkPart1-8                13914895                86.54 ns/op
BenchmarkPart2a-8                1536236               768.9 ns/op
BenchmarkPart2b-8                2039821               584.0 ns/op
BenchmarkPart2c-8                5441077               221.0 ns/op
BenchmarkPart1_cache_opt-8      14915084                80.29 ns/op
BenchmarkPart2a_cache_opt-8      1780618               664.9 ns/op
BenchmarkPart2b_cache_opt-8      2295934               524.4 ns/op
BenchmarkPart2c_cache_opt-8      4852630               246.3 ns/op
PASS
ok      day1    12.852s

```

Funnily enough, all solutions benifit from this optimization except for the third
approach, and honestly I don't quite know why as of now.

_Addendum_: 
Actually some of the performance implications from the previous
benchmark makes sense:
- The sorting solution and three linear pass solution (`2a` and `2b`) benifitted 
  the most out of making the underlying datastructure cache friendly, which makes
  sense, because the cache only helps optimize memory access to the same data
  object over and over again. Both sorting and the three linear pass solution 
  have this property, hence the somewhat sizable performance improvement.
- I still dont know why the `2c` performs worse tho :/ That's something to ponder
  about.
