package main

func part1(inventories []ElfInventory) int {
    var max int = -1
    for _, inv := range inventories {
        if inv.total > max {
            max = inv.total
        }
    }

    return max
}

func part1_cache_opt(inventories []int) int {
    var max int = -1
    for _, inv := range inventories {
        if inv > max {
            max = inv
        }
    }

    return max
}
