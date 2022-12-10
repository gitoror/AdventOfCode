# Advent Of Code 2022

Language used : JavaScript

## Day 1

I used sort(), but I could have used a simple specific algo to remain in O(n)

## Day 2

I used `modulo` to avoid listing all the cases

## Day 3

A simple but bad idea is to use nested `for` loops. To reduce the complexity, it's better to use another data structure : `Set`

## Day 4

Very simple

## Day 5

More data to format but it is still simple. I factor the appropriate code to solve part 1 and 2 because they are more about 95% similar

## Day 6

## Day 7

I struggled a lot on that one because I didn't know how to visit every element of an array that has an unknown depth. I manage to do it with recursivity but I think it is a constraint to improve the complexity in this case because we can't take shortcuts to stop the exploration. Difficulties to choose the appropriate data structure.

## Day 8

I tried to make a clever efficient solution but I had problems. The best is always to start by the obvious solution, don't search to find shortcuts.

## Day 9

Difficult problem to understand, part 1 was ok, but for part 2 I needed a long time to understand how the knots of the rope were moving !!! It was not compatible with my approach for part 1 where I made 2 moves to move diagonnaly, and that have been a disaster ... Also, it's important to put `else if`and not just `if` because with recursivity the programm can change the coordiantes of a knot that has already been treated

## Day 10

Cool and relaxing problem, the more diffcult was to read the letters on the CRT ... and I'm serious about that :`)
