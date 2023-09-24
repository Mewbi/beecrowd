x = int(input())

while True:
    z = int(input())
    if z > x:
        break

sum_val = x
count = 1

while True:
    sum_val = sum_val + x + count
    count = count + 1

    if sum_val > z:
        break

print(count)
