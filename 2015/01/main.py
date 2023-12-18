with open('01/in.txt', 'r') as f:
  data = f.read()

r=0
for s in data:
  if s == "(":
    r+=1
  else:
    r-=1
print(r) # Part 1

step,r=0,0
for k,s in enumerate(data):
  if s == "(":
    r+=1
  else:
    r-=1
  if r == -1:
    step = k+1
    break
print(step) # Part 2