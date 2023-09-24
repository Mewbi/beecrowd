def fast_fibonacci(N):
    if N <= 1:
        return N

    a, b = 0, 1
    for i in range(2, N + 1):
        a, b = b, a + b
    return b

n = int(input())

for _ in range(n):
    x = int(input())
    print("Fib({}) = {}".format(x, fast_fibonacci(x)))
