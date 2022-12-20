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

Cool and relaxing problem, the more diffcult was to read the letters on the CRT ... and I'm serious about that :`(

## Day 11

Part 1 okay but struggle to find the idea that completely unlocked part 2 in a blink of an eye

## Day 12

Hopefully I heard my classmates talking about Dijkstra algorithm to find the shortest path of in a graph. Otherwise I would have tried to make recursivity to explore paths for each neighbor. I struggle to really have an intuition of what is doing Dijkstra algorithm..

## Day 13

I have exploited JavaScript JSON parsing to parse the array string input. I had some logic issues fo `if` conditions and for the recursive calls in the main `isOrdered` function but I manage to debug it. I have exploited the `sort` function in JS with a callBack function corresponding to the function made in part 1.

## Day 14

I used a matrix but it would have been more efficient to use maps.

## Day 15

Very hard Part 2 because of time execution. I had sevaral ideas to reduce the complexity but it is still over 1 sec ... I think the order of the sensors in the input is important because when I order my sensors by x rising the time is exceeding 1 min ! I have 2 algos for part 2, The one that is commented have a time of 15 sec and the presented one 2 sec. In the commented one I stayed one the same idea as part 1 with intervals. On the other version, which is more efficient, I focus on the possible zone of 1 box width around sensors where a beacon can be. And because of the unicity of the possible undetected beacon, the answer is the point of these zones that can't be detected by any sensor.

## Day 16

Hard problem, didn't do it

## Day 17

Goal : Mesure the size of a tower made of rocks after 2022 rocks have felt

Part 1 : Easy, even if my solution is a bit messy and base on copy paste which is not the point of programming... I know, I know

Part 2 : time excution pb, have to find cycles but don't understand how this can work bcs we cannot predict anything ...?

## Day 18

Goal : Calculate the area of a lava droplet

Part 1 okay , but in part 2 we don't want to include the area due to air bubbles in the droplet. The trick is to build a box that entirely contains the droplet, start from a point in that box and draw the real contour of the droplet.

## Day 19

No up to date at this point... It's holidays :) (lol but it's only work until now)
