import copy

with open("in.txt", "r") as f:
    data = f.read().split(",")


def calc_hash(code):
    a = 0
    for s in code:
        a += ord(s)
        a *= 17
        a = a % 256
    return a


def part1(data):
    sum = 0
    for code in data:
        sum += calc_hash(code)
    return sum


def part2(data):
    boxes = [[] for k in range(256)]
    for code in data:
        box_info = code.split("=")
        if len(box_info) == 2:
            box_id = calc_hash(box_info[0])
            operation = "="
            v = code[-1]
        else:
            box_id = calc_hash(box_info[0][:-1])
            operation = "-"
        if operation == "=":
            already_modified = False
            for len_id, lens in enumerate(boxes[box_id]):
                if lens[0] == box_info[0]:
                    boxes[box_id][len_id][1] = int(v)
                    already_modified = True
            if not already_modified:
                boxes[box_id].append([box_info[0], int(v)])
        else:
            for len_id, lens in enumerate(boxes[box_id]):
                if lens[0] == box_info[0][:-1]:
                    # print(boxes[box_id], len_id)
                    boxes[box_id].pop(len_id)
    sum = 0
    for box_id, box in enumerate(boxes):
        for len_id, lent in enumerate(box):
            sum += (1 + box_id) * (1 + len_id) * lent[1]
    print("PRINT")
    for box in boxes:
        print(box)
    return sum


print(part1(data))
print(part2(data))
