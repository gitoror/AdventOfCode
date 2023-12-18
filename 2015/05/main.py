import itertools


with open('05/in.txt', 'r') as f:
  data = f.read().split('\n')

# data = ["ugknbfddgicrmopn","aaa","jchzalrnumimnmhp","haegwjzuvuyypxyu","dvszwmarrgswjxmb"]
sum = 0
for word in data:
  vowels = 0
  twice = False
  last_condition = True
  for i in range(len(word)):
    letter = word[i]
    if vowels != 3 and letter in "aeiou":
      vowels += 1
    if i != 0:
      if word[i-1] == letter:
        twice = True
      if word[i-1:i+1] in ["ab","cd","pq","xy"]:
        last_condition = False
        break
  if vowels == 3 and twice and last_condition:
    sum += 1

print(sum)

sum = 0
for word in data:
  dic = {}
  cond1 = False
  cond2 = False
  for i, l in enumerate(word):
    if l in dic:
      dic[l].append(i)
    else: 
      dic[l] = [i]
  for k in dic:
    poss = dic[k]
    if len(poss) > 1:
      for i,j in itertools.combinations(poss,2):
        if abs(j-i) > 1 and word[i:i+2] == word[j:j+2]:
          cond1 = True
          break
      if cond1:
        break
  for i, l in enumerate(word):
    if i > 1 and l == word[i-2]:
      cond2 = True
      break
  if cond1 and cond2:
    sum += 1
  
print(sum)
        
