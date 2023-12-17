import heapq

grid = {}
with open("inputs/17/in.txt", "r") as f:
    data = f.read().splitlines()
    lenX,lenY = len(data[0]), len(data)
    for y,line in enumerate(data):
        for x,c in enumerate(line):
            grid[(x,y)] = int(c)
        

def part1(grid):
    ans = 0
    D = [(1,0),(0,1),(-1,0),(0,-1)]
    src = (0,0)
    dest = (lenX-1,lenY-1)
    heap = []
    heapq.heappush(heap, (0,src,(0,0))) # (dst, pos, prev_infos) 
    explored = set() # (X,Y,prev_infos) set of explored nodes 
    while heap:
        dst, pos, prev_infos = heapq.heappop(heap)
        prev_dir, count = prev_infos # prev_dir: 0: right, 1: down, 2: left, 3: up; count: number of consecutive lines or columns
        if pos == dest:
            ans = dst
            break
        if (pos,prev_dir,count) in explored:
            continue
        explored.add((pos,prev_dir,count))
        for next_dir, (dx,dy) in enumerate(D): # next possible moves
            # k: 0: right, 1: down, 2: left, 3: up
            if (prev_dir, count) == (next_dir,3) or next_dir == (prev_dir+2)%4:
                # We can't go more than 3 times in the same direction
                # If we come from the opposite direction, we can't go back
                continue
            x,y = pos
            next_x, next_y = x+dx, y+dy
            if 0 <= next_x < lenX and 0 <= next_y < lenY: # and (next_x,next_y,v): # what is  and (next_x,next_y,v) ? 
                if prev_dir == next_dir:
                    new_count = count+1
                else:
                    new_count = 1
                new_dst = dst + grid[(next_x,next_y)]
                heapq.heappush(heap, (new_dst, (next_x,next_y), (next_dir, new_count)))
    return ans

def part2(grid):
    ans = 0
    D = [(1,0),(0,1),(-1,0),(0,-1)]
    src = (0,0)
    dest = (lenX-1,lenY-1)
    heap = []
    heapq.heappush(heap, (0,src,(0,0))) # (dst, pos, prev_infos) 
    explored = set() # (X,Y,prev_infos) set of explored nodes 
    while heap:
        dst, pos, prev_infos = heapq.heappop(heap)
        prev_dir, count = prev_infos # prev_dir: 0: right, 1: down, 2: left, 3: up; count: number of consecutive lines or columns
        if pos == dest:
            if count < 4:
                continue
            ans = dst
            break
        if (pos,prev_dir,count) in explored:
            continue
        explored.add((pos,prev_dir,count))
        for next_dir, (dx,dy) in enumerate(D): # next possible moves
            # k: 0: right, 1: down, 2: left, 3: up
            if (prev_dir != next_dir and 0 < count < 4) or (prev_dir, count) == (next_dir,10) or next_dir == (prev_dir+2)%4:
                # We need to go at least 4 times in the same direction before turning
                # We can't go more than 10 times in the same direction
                # If we come from the opposite direction, we can't go back
                continue
            x,y = pos
            next_x, next_y = x+dx, y+dy
            if 0 <= next_x < lenX and 0 <= next_y < lenY: # and (next_x,next_y,v): # what is  and (next_x,next_y,v) ? 
                if prev_dir == next_dir:
                    new_count = count+1
                else:
                    new_count = 1
                new_dst = dst + grid[(next_x,next_y)]
                heapq.heappush(heap, (new_dst, (next_x,next_y), (next_dir, new_count)))
    return ans

# print(part1(grid))
print(part2(grid))