package main

import (
	"strings"
)

func part2(input string) int {
    var result int
    lines := strings.Split(input, "\n")

    for i := 0; i < len(lines); i += 3 {
        backpacks := []Backpack {
            backpack_from_str(lines[i]),
            backpack_from_str(lines[i + 1]),
            backpack_from_str(lines[i + 2]),
        }

        result += priority(badge_item(backpacks))
    }

    return result
}
