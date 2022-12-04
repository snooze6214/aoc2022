package main

import (
	"strings"
)

func part2(input string) int {
    var result int

    pair_ranges := strings.Split(input, "\n")
    for _, pair_str := range pair_ranges {
        pair_str_split := strings.Split(pair_str, ",")

        r1, err := range_from_str(pair_str_split[0])
        if err != nil {
            return -1
        }

        r2, err := range_from_str(pair_str_split[1])
        if err != nil {
            return -1
        }

        if overlaps(r1, r2) {
            result++
        }
    }

    return result
}
