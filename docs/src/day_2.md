# Day 2: Rock Paper Scissors

[Link](https://adventofcode.com/2022/day/2)

Nothing special going on here other than simply simulating the rules of the
game.

### Also cool observation:
we know:
```
      beats         beats
rock <------ paper <------ scissors
 |                            ^
 |                            |
 l-----------------------------
         beats
```

So if we assign to each shape the following values:
```
Rock = 0,
Paper = 1,
Scissors = 2
```

- Then to find a shape that beats shape `x`, we just calculate `(x + 1) mod 3`
- Then to find a shape that looses to shape `x`, we just calculate `(x + 2) mod 3`
