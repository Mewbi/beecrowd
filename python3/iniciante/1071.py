a = int(input())
b = int(input())

start, endVal, sum_val = a, b, 0

if a > b:
    start, endVal = b, a

for i in range(start + 1, endVal):
    if abs(i % 2) == 1:
        sum_val += i

print(sum_val)
