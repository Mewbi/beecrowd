def bubble_sort(arr):
    n = len(arr)
    for i in range(n - 1):
        for j in range(n - 1 - i):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]

l = []
old = []

l = list(map(int, input().split()))

old = l.copy()

bubble_sort(l)

for i in l:
    print(i)

print()

for i in old:
    print(i)
