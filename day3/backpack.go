package main

type Backpack struct {
    compartment1 map[rune]bool
    compartment2 map[rune]bool
    items map[rune]bool
}

func backpack_from_str(backpack_str string) Backpack {
    compartment1_str := backpack_str[0:(len(backpack_str) / 2)]
    compartment2_str := backpack_str[(len(backpack_str) / 2):]

    var backpack Backpack
    backpack.compartment1 = make(map[rune]bool)
    backpack.compartment2 = make(map[rune]bool)
    backpack.items = make(map[rune]bool)

    for _, c := range compartment1_str {
        backpack.compartment1[c] = true
        backpack.items[c] = true
    }

    for _, c := range compartment2_str {
        backpack.compartment2[c] = true
        backpack.items[c] = true
    }

    return backpack
}

func priority(char rune) int {
    if int(char) >= int('a') && int(char) <= int('z') {
        return int(char) - int('a') + 1
    }

    if int(char) >= int('A') && int(char) <= int('Z') {
        return 27 + int(char) - int('A') 
    }

    return 0
}

func problem_item(backpack Backpack) rune {
    for item := range backpack.compartment1 {
        if backpack.compartment2[item] {
            return item
        }
    }

    for item := range backpack.compartment2 {
        if backpack.compartment1[item] {
            return item
        }
    }

    return '*'
}

func badge_item(backpacks []Backpack) rune {
    if len(backpacks) != 3 {
        return '*'
    }

    for i := range backpacks[0].items {
        if backpacks[1].items[i] && backpacks[2].items[i] {
            return i
        }
    }

    for i := range backpacks[1].items {
        if backpacks[1].items[0] && backpacks[2].items[i] {
            return i
        }
    }

    for i := range backpacks[2].items {
        if backpacks[0].items[i] && backpacks[1].items[i] {
            return i
        }
    }

    return '*'
}
