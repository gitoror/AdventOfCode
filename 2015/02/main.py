with open('02/in.txt', 'r') as f:
  data = f.read().split('\n')

sum = 0
for i in range(len(data)):
  l,w,h = data[i].split('x')
  l,w,h = int(l),int(w),int(h)
  sum+= 2*l*w + 2*w*h + 2*h*l + min(l*w,w*h,h*l)

print(sum)

sum = 0
for i in range(len(data)):
  l,w,h = data[i].split('x')
  l,w,h = int(l),int(w),int(h)
  a,b,c = sorted([l,w,h])
  sum+= a*2 + b*2 + l*w*h

print(sum)