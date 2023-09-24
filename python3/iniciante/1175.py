arr = []

for _ in range(20):
    x = int(input())
    arr.append(x)

for i, _ in enumerate(arr):
    print("N[{}] = {}".format(i, arr[-(i + 1)]))
