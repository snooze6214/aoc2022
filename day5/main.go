package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
    input_bytes, err := ioutil.ReadFile("input")
    if err != nil {
        log.Fatal("Failed to read input file.")
    }

    input := string(input_bytes)
    input = strings.Trim(input, "\n")

    input_split := strings.Split(input, "\n\n")
    crate_stacks := crate_stacks_from_str(input_split[0])
    instr_strs := strings.Split(strings.Trim(input_split[1], "\n"), "\n")

    part1_res := part1(crate_stacks, instr_strs)
    fmt.Println("Part 1: ", part1_res)

    crate_stacks = crate_stacks_from_str(input_split[0])
    part2_res := part2(crate_stacks, instr_strs)
    fmt.Println("Part 2: ", part2_res)
}
