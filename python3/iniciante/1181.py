n = int(input())
calc = input()

sum_val = 0.0

for i in range(12):
    for j in range(12):
        x = float(input())
        if n == i:
            sum_val += x

if calc == "M":
    sum_val = sum_val / 12.0

print("{:.1f}".format(sum_val))
