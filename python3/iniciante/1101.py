while True:
    m, n = list(map(int, input().split()))
    if n <= 0 or m <= 0:
        break

    maior, menor = m, n
    if n > m:
        maior, menor = n, m

    sum = 0
    for i in range(menor, maior + 1):
        sum += i
        print(i, end=' ')

    print("Sum={}".format(sum))
