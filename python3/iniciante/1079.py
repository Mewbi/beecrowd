n = int(input())

for _ in range(n):
    n1, n2, n3 = list(map(float, input().split()))

    m = (n1 * 2 + n2 * 3 + n3 * 5) / 10
    print("{:.1f}".format(m))
