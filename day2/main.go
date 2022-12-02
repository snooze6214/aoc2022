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

    part1_res := part1(input)
    fmt.Println("Part 1: ", part1_res)

    part2_res := part2(input)
    fmt.Println("Part 1: ", part2_res)
}
