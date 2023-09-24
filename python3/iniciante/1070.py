n = int(input())
tot = 0

while True:
    if n % 2 == 1:
        print(n)
        tot += 1
    n += 1

    if tot == 6:
        break
