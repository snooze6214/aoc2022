# Day 4: Camp Cleanup

[Link](https://adventofcode.com/2022/day/4)

Ok I think the most important lesson I learnt from this problem is detecting 
overlapping ranges. So here we go

## Overlapping ranges
The problem at hand is to find a sufficient and necessary condition to check
whether two ranges are overlapping. So I formally define a range as a set of
numbers that are inbetween a start `x` and a stop `y`, inclusive. For notational
convinience, I denote such a range as `[x:y]`

Now, I define what I mean by two ranges overlapping. Two ranges `A` and `B` are
said to overlap if there exist some number `a` such that `a` is in `A` and `a`
is in `B` as well. Thus these two inequalities must hold when ranges `A` and `B`
overlap.

```
A.start <= a <= A.end
B.start <= a <= B.end
```
From the above two inequalities it should be obvious that a necessary condition
that should hold when `A` and `B` overlap are:
```
(A.start <= B.end) ^ (B.start <= A.end)
```
Now to show that this is a sufficient condition, we assume the condition holds.
Then suppose we choose some `a` in the range `[max(A.start, B.start):B.end]`
then it'll be in both `A` and `B`.
