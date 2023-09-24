def is_perfect(n):
    sum_val = 0
    for i in range(1, n):
        if sum_val > n:
            break
        if n % i == 0:
            sum_val += i
    return sum_val == n

n = int(input())

for i in range(n):
    x = int(input())

    if is_perfect(x):
        print(str(x) + " eh perfeito")
    else:
        print(str(x) + " nao eh perfeito")
