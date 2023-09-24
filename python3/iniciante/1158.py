n = int(input())

for _ in range(n):
    x, y = map(int, input().split())
    sum_val, count = 0, 0

    while count < y:
        if x % 2 != 0:
            count += 1
            sum_val += x

        x += 1

    print(sum_val)
