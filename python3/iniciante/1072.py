n = int(input())

inCount, outCount = 0, 0

for i in range(n):
    v = int(input())

    if 10 <= v <= 20:
        inCount += 1
    else:
        outCount += 1

print("{} in".format(inCount))
print("{} out".format(outCount))
