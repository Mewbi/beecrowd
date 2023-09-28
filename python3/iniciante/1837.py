def euclideanDivision(a, b):
    limit = -b if b < 0 else b
    r = 0
    q = 0

    for r in range(limit):
        if (a - r) % b == 0:
            q = (a - r) // b
            break

    return q, r

a, b = map(int, input().split())
q, r = euclideanDivision(a, b)
print(q, r)
