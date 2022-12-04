package main

import (
	"strconv"
	"strings"
)

type Range struct {
    start int
    end int
}

func range_from_str(range_str string) (Range, error) {
    var result Range
    split_strs := strings.Split(range_str, "-")

    start, err := strconv.Atoi(split_strs[0])
    if err != nil {
        return result, err
    }

    end, err := strconv.Atoi(split_strs[1])
    
    if err != nil {
        return result, err
    }

    result.start = start
    result.end = end
    
    return result, nil
}

func is_contained_in(a, b Range) bool {
    return b.start <= a.start && a.end <= b.end 
}

func covers(a, b Range) bool {
    return is_contained_in(a, b) || is_contained_in(b, a)
}

func overlaps(a, b Range) bool {
    return (a.start <= b.end) && (b.start <= a.end)
}
