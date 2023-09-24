n = int(input())

for i in range(n):
    v = int(input())

    if v == 0:
        print("NULL")
    elif v % 2 == 0:
        print("EVEN", end=" ")
    elif abs(v % 2) == 1:
        print("ODD", end=" ")

    if v > 0:
        print("POSITIVE")
    elif v < 0:
        print("NEGATIVE")
