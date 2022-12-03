package main

import (
	"strings"
)

func part1(input string) int {
    var result int
    lines := strings.Split(input, "\n")
    for _, line := range lines {
        backpack := backpack_from_str(line)
        problem := problem_item(backpack)
        result += priority(problem)
    }

    return result
}
