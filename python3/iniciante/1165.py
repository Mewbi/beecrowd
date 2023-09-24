import math

def is_prime(n):
    if n <= 1:
        return False

    for i in range(2, int(math.sqrt(n)) + 1):
        if n % i == 0:
            return False

    return True

n = int(input())

for i in range(n):
    x = int(input())

    if is_prime(x):
        print(str(x) + " eh primo")
    else:
        print(str(x) + " nao eh primo")
