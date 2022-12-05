package main

import (
	"strconv"
	"strings"
)

type CrateStack []byte
type CrateStacks []CrateStack

func crate_stacks_from_str(crate_stacks_str string) CrateStacks {
    lines := strings.Split(crate_stacks_str, "\n")
    row_width := len(lines[0])
    num_stacks := (row_width + 1) / 4

    var result []CrateStack = make([]CrateStack, num_stacks)

    for i := 0; i < num_stacks; i++ {
        index := 1 + 4*i
        for j := len(lines) - 2; j >= 0; j-- {
            if lines[j][index] == byte(' ') {
                break
            }
            result[i] = append(result[i], lines[j][index])
        }
    }

    return result
}

func (s CrateStacks) move(src, dest, num int, maintain_order bool) {
    src_len := len(s[src])
    
    if maintain_order {
        for i := src_len - num; i < src_len; i++ {
            s[dest] = append(s[dest], s[src][i])
        }
    } else {
        for i := src_len - 1; i >= src_len - num; i-- {
            s[dest] = append(s[dest], s[src][i])
        }
    }

    s[src] = s[src][:src_len-num]
}

func (s CrateStacks) move_from_instr_str(instr_str string, maintain_order bool) error {
    instr_str_split := strings.Split(instr_str, " ")

    num, err := strconv.Atoi(instr_str_split[1])
    if err != nil {
        return err
    }

    src, err := strconv.Atoi(instr_str_split[3])
    if err != nil {
        return err
    }
    src--

    dest, err := strconv.Atoi(instr_str_split[5])
    if err != nil {
        return err
    }
    dest--

    s.move(src, dest, num, maintain_order)

    return nil
}
