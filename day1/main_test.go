package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func setup() []ElfInventory {
    input_bytes, err := ioutil.ReadFile("input")
    if err != nil {
        log.Fatal("Failed to read input file.")
    }

    input := string(input_bytes)
    return invs_from_string(input)
}


func setup2() []int {
    input_bytes, err := ioutil.ReadFile("input")
    if err != nil {
        log.Fatal("Failed to read input file.")
    }

    input := string(input_bytes)
    return invs_from_string2(input)
}

func BenchmarkPart1(b *testing.B) {
    inventories := setup()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        part1(inventories)
    }
}

func BenchmarkPart2a(b *testing.B) {
    inventories := setup()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        part2a(inventories)
    }
}

func BenchmarkPart2b(b *testing.B) {
    inventories := setup()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        part2b(inventories)
    }
}

func BenchmarkPart2c(b *testing.B) {
    inventories := setup()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        part2c(inventories)
    }
}

func BenchmarkPart1_cache_opt(b *testing.B) {
    inventories := setup2()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        part1_cache_opt(inventories)
    }
}

func BenchmarkPart2a_cache_opt(b *testing.B) {
    inventories := setup2()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        part2a_cache_opt(inventories)
    }
}

func BenchmarkPart2b_cache_opt(b *testing.B) {
    inventories := setup2()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        part2b_cache_opt(inventories)
    }
}

func BenchmarkPart2c_cache_opt(b *testing.B) {
    inventories := setup2()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        part2c_cache_opt(inventories)
    }
}
