arr = []

for _ in range(10):
    x = int(input())
    if x <= 0:
        x = 1

    arr.append(x)

for i, v in enumerate(arr):
    print("X[{}] = {}".format(i, v))
