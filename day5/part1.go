package main

import "strings"

func part1(crate_stacks CrateStacks, instr_strs []string) string {
    for _, instr_str := range instr_strs {
        crate_stacks.move_from_instr_str(instr_str, false)
    }
    
    var result_builder strings.Builder

    for _, stack := range crate_stacks {
        result_builder.WriteByte(stack[len(stack) - 1])
    }

    return result_builder.String()
}
