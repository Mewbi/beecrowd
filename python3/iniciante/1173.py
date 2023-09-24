arr = []
x = int(input())

for _ in range(10):
    arr.append(x)
    x = x * 2

for i, v in enumerate(arr):
    print("N[{}] = {}".format(i, v))
