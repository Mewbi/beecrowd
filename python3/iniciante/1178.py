arr = []
x = float(input())
arr.append(x)

for i in range(2, 101):
    x = x / 2
    arr.append(x)

for i, v in enumerate(arr):
    print("N[{}] = {:.4f}".format(i, v))
