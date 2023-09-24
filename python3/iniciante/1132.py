x = int(input())
y = int(input())
sum = 0

if y < x:
    x, y = y, x

for i in range(x, y + 1):
    if i % 13 != 0:
        sum += i

print(sum)
