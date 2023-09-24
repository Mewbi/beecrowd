arr = []
x = int(input())
n = 0

for i in range(1000):
    arr.append(n)
    n = n + 1
    if n >= x:
        n = 0

for i, v in enumerate(arr):
    print("N[{}] = {}".format(i, v))
