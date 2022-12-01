package main

import "sort"

func part2a(inventories []ElfInventory) int {
    sort.Slice(inventories, func(i, j int) bool {
        return inventories[i].total > inventories[j].total
    })

    return inventories[0].total + inventories[1].total + inventories[2].total
}

func part2a_cache_opt(inventories []int) int {
    sort.Slice(inventories, func(i, j int) bool {
        return inventories[i] > inventories[j]
    })

    return inventories[0] + inventories[1] + inventories[2]
}

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

func part2b_cache_opt(inventories []int) int {
    var top1 int = -1
    var top2 int = -1
    var top3 int = -1
    var t1, t2 int

    for i, inv := range inventories {
        if inv > top1 {
            top1 = inv
            t1 = i
        }
    }

    for i, inv := range inventories {
        if i != t1 && inv > top2 {
            top2 = inv
            t2 = i
        }
    }

    for i, inv := range inventories {
        if i != t1 && i != t2 && inv > top3 {
            top3 = inv
        }
    }

    return top1 + top2 + top3
}

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

func part2c_cache_opt(inventories []int) int {
    var top1 int = -1
    var top2 int = -1
    var top3 int = -1

    for _, inv := range inventories {
        if inv > top1 {
            top3 = top2
            top2 = top1
            top1 = inv
            continue
        }

        if inv > top2 {
            top3 = top2
            top2 = inv
            continue
        }

        if inv > top3 {
            top3 = inv
        }
    }

    return top1 + top2 + top3
}
