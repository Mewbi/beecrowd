name = input()
s = float(input())
v = float(input())

total = s + 0.15 * v
total = "{:.2f}".format(total)

print("TOTAL = R$ " + total)
