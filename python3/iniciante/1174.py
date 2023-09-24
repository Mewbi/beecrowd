arr = []

for _ in range(100):
    x = float(input())
    arr.append(x)

for i, v in enumerate(arr):
    if v <= 10:
        print("A[{}] = {:.1f}".format(i, v))
