with open('03/in.txt', 'r') as f:
  data = f.read()

x,y=0,0
seen = set()
seen.add((x,y))
for s in data:
  if s == "^":
    y+=1
  elif s == "v":
    y-=1
  elif s == ">":
    x+=1
  else:
    x-=1
  seen.add((x,y))
print(len(seen)) # Part 1

xs,ys,xr,yr=0,0,0,0
seen = set()
seen.add((xs,ys))
for k,s in enumerate(data):
  if k%2 == 0:
    if s == "^":
      ys+=1
    elif s == "v":
      ys-=1
    elif s == ">":
      xs+=1
    else:
      xs-=1
    seen.add((xs,ys))
  else:
    if s == "^":
      yr+=1
    elif s == "v":
      yr-=1
    elif s == ">":
      xr+=1
    else:
      xr-=1
    seen.add((xr,yr))
print(len(seen)) # Part 2