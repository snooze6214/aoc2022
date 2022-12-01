package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
    input_bytes, err := ioutil.ReadFile("input")
    if err != nil {
        log.Fatal("Failed to read input file.")
    }

    input := string(input_bytes)
    inventories := invs_from_string(input)

    part1_res := part1(inventories)
    // part2_res := part2a(inventories)
    // part2_res := part2b(inventories)
    part2_res := part2c(inventories)

    fmt.Println("Part 1: ", part1_res)
    fmt.Println("Part 2: ", part2_res)
}
