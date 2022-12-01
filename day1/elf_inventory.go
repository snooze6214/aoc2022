package main

import (
    "fmt"
    "strings"
)

type ElfInventory struct {
    items []int
    total int
}

func invs_from_string(input string) []ElfInventory {
    inv_lists := strings.Split(input, "\n\n")
    var result []ElfInventory

    for _, inv_list_str := range inv_lists {
        var inv ElfInventory
        inv_list := strings.Split(inv_list_str, "\n")

        for _, inv_amt_str := range inv_list {
            var amt int
            fmt.Sscan(inv_amt_str, &amt)
            inv.items = append(inv.items, amt)
            inv.total += amt
        }
        
        result = append(result, inv)
    }

    return result
}

func invs_from_string2(input string) []int {
    inv_lists := strings.Split(input, "\n\n")
    var result []int

    for _, inv_list_str := range inv_lists {
        var total int
        inv_list := strings.Split(inv_list_str, "\n")

        for _, inv_amt_str := range inv_list {
            var amt int
            fmt.Sscan(inv_amt_str, &amt)
            total += amt
        }
        
        result = append(result, total)
    }

    return result
}
