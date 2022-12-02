package main

import (
	"strings"
)

func part1(input string) int {
    var total_score int

    round_strs := strings.Split(input, "\n")
    for _, round_str := range round_strs {
        if round_str != "" {
            round := game_round_from_str(round_str, false)
            total_score += score(round)
        }
    }

    return total_score
}
