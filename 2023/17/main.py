import heapq

grid = {}
with open("inputs/17/in.txt", "r") as f:
    data = f.read().split("\n")
    lenX,lenY = len(data[0]), len(data)
    for y,line in enumerate(data):
        for x,c in enumerate(line):
            grid[(x,y)] = int(c)
        
D = [(1,0),(0,1),(-1,0),(0,-1)]

def solve(part2):
    ans = 0
    src,dir_,count = (0,0),0,0
    # dir_: 0: right, 1: down, 2: left, 3: up; count: number of consecutive lines or columns
    dest = (lenX-1,lenY-1)
    heap = []
    heapq.heappush(heap, (0,src,dir_,count)) # (dst, pos, prev_infos) 
    explored = set() # (X,Y,prev_infos) set of explored nodes 
    while heap:
        dst, pos, dir_, count = heapq.heappop(heap)
        if pos == dest:
            if part2 and count < 4:
                continue
            ans = dst
            break
        if (pos,dir_,count) in explored:
            continue
        explored.add((pos,dir_,count))
        for next_dir, (dx,dy) in enumerate(D): # next possible moves
            x,y = pos
            next_x, next_y = x+dx, y+dy
            if ((next_x,next_y),next_dir,count) in explored:
                continue
            # k: 0: right, 1: down, 2: left, 3: up
            if not part2 and (dir_, count) == (next_dir,3) or next_dir == (dir_+2)%4:
                # We can't go more than 3 times in the same direction
                # If we come from the opposite direction, we can't go back
                continue
            if part2 and (dir_ != next_dir and 0 < count < 4) or (dir_, count) == (next_dir,10) or next_dir == (dir_+2)%4:
                # We need to go at least 4 times in the same direction before turning, count>0 bcs count at src is 0 and we can turn at the first move
                # We can't go more than 10 times in the same direction
                # If we come from the opposite direction, we can't go back
                continue
            if 0 <= next_x < lenX and 0 <= next_y < lenY: # and (next_x,next_y,v): # what is  and (next_x,next_y,v) ? 
                if dir_ == next_dir:
                    new_count = count+1
                else:
                    new_count = 1
                new_dst = dst + grid[(next_x,next_y)]
                heapq.heappush(heap, (new_dst,(next_x,next_y),next_dir,new_count))
    return ans


print(solve(False)) # Part 1
print(solve(True)) # Part 2