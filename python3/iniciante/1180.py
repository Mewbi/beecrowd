n = int(input())
line = list(map(int, input().split()))
v = line[0]
nums = line[1:]
p = 0

for i, x in enumerate(nums):
    if x < v:
        v = x
        p = i + 1

print("Menor valor:", v)
print("Posicao:", p)
