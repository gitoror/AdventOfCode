with open("../inputs/12/input.txt", "r") as f:
    data = f.read().split("\n")
 
records = []
for line in data:
    springs, sizes = line.split(" ")
    sizes = [int(x) for x in sizes.split(",")]
    records.append([springs, sizes])
 
memo_map = {}
 

def operational(springs, start, end):
    if (start - 1 >= 0 and springs[start - 1] == "#") or (
        end < len(springs) and springs[end] == "#"
    ):
        return False
    for c in springs[start:end]:
        if c == ".":
            return False
    return True
 

def arrangements(springs, sizes):
    if (springs, tuple(sizes)) in memo_map:  # memoization part 2
        return memo_map[(springs, tuple(sizes))]
    size = sizes[0]
    sum = 0
    for start in range(0, len(springs) - size + 1):
        end = start + size  # careful, does not include the last spring of the studied
        if operational(springs, start, end):
            if len(sizes) == 1:
                if not ("#" in springs[:start] + springs[end:]):
                    sum += 1
            else:
                if not ("#" in springs[: max(0, start - 1)]):
                    sum += arrangements(springs[end + 1 :], sizes[1:])
    memo_map[(springs, tuple(sizes))] = sum  # memoization part 2
    return sum
 

def solve(records):
    return sum(arrangements(springs, sizes) for springs, sizes in records)
 

print("PART 1", solve(records))
 
records2 = []
for springs, sizes in records:
    records2.append([(springs + "?") * 4 + springs, sizes * 5])
 
print("PART 2", solve(records2))