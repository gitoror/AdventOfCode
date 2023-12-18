import copy

with open("in.txt", "r") as f:
    data = f.read().split("\n")


def display(rocks):
    for line in rocks:
        l = ""
        for c in line:
            l += c
        print(l)
    print()


input = []
for line in data:
    input.append(list(line))


def tilt(rocks):
    for y in range(1, len(rocks)):
        for x in range(len(rocks[0])):
            if rocks[y][x] == "O":
                c = y
                while c - 1 >= 0 and rocks[c - 1][x] == ".":
                    rocks[c - 1][x], rocks[c][x] = "O", "."
                    c -= 1


def calc_load(rocks):
    sum = 0
    for y in range(len(rocks)):
        for x in range(len(rocks[0])):
            if rocks[y][x] == "O":
                sum += len(rocks) - y
    return sum


def part1(input):
    rocks = copy.deepcopy(input)
    tilt(rocks)
    return calc_load(rocks)


print(part1(input))


def spin_cycle(rocks):
    # north
    for y in range(1, len(rocks)):
        for x in range(len(rocks[0])):
            if rocks[y][x] == "O":
                c = y
                while c - 1 >= 0 and rocks[c - 1][x] == ".":
                    rocks[c - 1][x], rocks[c][x] = "O", "."
                    c -= 1
    # west
    for x in range(1, len(rocks[0])):
        for y in range(len(rocks)):
            if rocks[y][x] == "O":
                c = x
                while c - 1 >= 0 and rocks[y][c - 1] == ".":
                    rocks[y][c - 1], rocks[y][c] = "O", "."
                    c -= 1
    # south
    for y in range(len(rocks) - 2, -1, -1):
        for x in range(len(rocks[0])):
            if rocks[y][x] == "O":
                c = y
                while c + 1 < len(rocks) and rocks[c + 1][x] == ".":
                    rocks[c + 1][x], rocks[c][x] = "O", "."
                    c += 1
    # east
    for x in range(len(rocks[0]) - 2, -1, -1):
        for y in range(len(rocks)):
            if rocks[y][x] == "O":
                c = x
                while c + 1 < len(rocks[0]) and rocks[y][c + 1] == ".":
                    rocks[y][c + 1], rocks[y][c] = "O", "."
                    c += 1


def part2(input):
    N = 1000000000
    rocks = copy.deepcopy(input)
    prev_rocks_list = []
    for k in range(N):
        prev_rocks_list.append(copy.deepcopy(rocks))
        spin_cycle(rocks)
        for j, prev_rocks in enumerate(prev_rocks_list):
            if prev_rocks == rocks:
                return calc_load(prev_rocks_list[(N - j) % (k + 1 - j) + j])
    return calc_load(rocks)


print(part2(input))
