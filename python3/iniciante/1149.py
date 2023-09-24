line = list(map(int, input().split()))
x = line[0]
rest = line[1:]
y = 0
for r in rest:
    if r > 0:
        y = r
        break

sum = 0

for i in range(y):
    sum += x + i

print(sum)
