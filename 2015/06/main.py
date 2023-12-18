with open('06/in.txt', 'r') as f:
  data = f.read().split('\n')

def part1():
  grid = [[0 for i in range(1000)] for j in range(1000)]
  for line in data:
    line = line.split()
    if len(line) ==5:
      r1,c1=line[2].split(",")
      r2,c2=line[4].split(",")
      if line[1] == "on":
        for r in range(int(r1),int(r2)+1):
          for c in range(int(c1),int(c2)+1):
            grid[r][c] = 1
      else:
        for r in range(int(r1),int(r2)+1):
          for c in range(int(c1),int(c2)+1):
            grid[r][c] = 0
    else:
      r1,c1=line[1].split(",")
      r2,c2=line[3].split(",")
      for r in range(int(r1),int(r2)+1):
        for c in range(int(c1),int(c2)+1):
          if grid[r][c] == 0:
            grid[r][c] = 1
          else:
            grid[r][c] = 0
  sum=0
  for r in range(len(grid)):
    for c in range(len(grid[0])):
      sum+=grid[r][c] 
  return sum
        
print(part1())

def part2():
  grid = [[0 for i in range(1000)] for j in range(1000)]
  for line in data:
    line = line.split()
    if len(line) ==5:
      r1,c1=line[2].split(",")
      r2,c2=line[4].split(",")
      if line[1] == "on":
        for r in range(int(r1),int(r2)+1):
          for c in range(int(c1),int(c2)+1):
            grid[r][c] = grid[r][c] + 1
      else:
        for r in range(int(r1),int(r2)+1):
          for c in range(int(c1),int(c2)+1):
            grid[r][c] = max(0, grid[r][c]-1)
    else:
      r1,c1=line[1].split(",")
      r2,c2=line[3].split(",")
      for r in range(int(r1),int(r2)+1):
        for c in range(int(c1),int(c2)+1):
          grid[r][c] += 2
  sum=0
  for r in range(len(grid)):
    for c in range(len(grid[0])):
      sum+=grid[r][c]
  return sum

# Rewrite part2 using dictionary
def part2_bis():
  grid = {}
  for line in data:
    line = line.split()
    if len(line) ==5:
      r1,c1=line[2].split(",")
      r2,c2=line[4].split(",")
      if line[1] == "on":
        for r in range(int(r1),int(r2)+1):
          for c in range(int(c1),int(c2)+1):
            if (r,c) not in grid:
              grid[(r,c)] = 1
            else:
              grid[(r,c)] += 1
      else:
        for r in range(int(r1),int(r2)+1):
          for c in range(int(c1),int(c2)+1):
            if (r,c) not in grid:
              grid[(r,c)] = 0
            else:
              grid[(r,c)] = max(0, grid[(r,c)]-1)
    else:
      r1,c1=line[1].split(",")
      r2,c2=line[3].split(",")
      for r in range(int(r1),int(r2)+1):
        for c in range(int(c1),int(c2)+1):
          if (r,c) not in grid:
            grid[(r,c)] = 2
          else:
            grid[(r,c)] += 2
  sum=0
  for r in range(1000):
    for c in range(1000):
      if (r,c) in grid:
        sum+=grid[(r,c)]
      
  return sum

  
# print(part2())
print(part2_bis())
