import hashlib

data = "iwrupvqb"

for i in range(10000000):
  if hashlib.md5((data+str(i)).encode()).hexdigest()[:5] == "00000":
    print(i) # Part 1
    break