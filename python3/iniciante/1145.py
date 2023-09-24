x, y = map(int, input().split())
count = 0

for i in range(1, y + 1):
    count += 1
    print(i, end="")

    if count == x:
        count = 0
        print()
    elif i != y:
        print(" ", end="")
