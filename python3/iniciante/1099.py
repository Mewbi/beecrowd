n = int(input())

for _ in range(n):
    x, y = list(map(int, input().split()))
    if y > x:
        maior = y
        menor = x
    else:
        maior = x
        menor = y

    sum = 0
    for j in range(menor + 1, maior):
        if j % 2 == 1:
            sum += j

    print(sum)
