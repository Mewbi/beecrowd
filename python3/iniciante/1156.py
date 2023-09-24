def seqSII(x, y):
    if x == 39:
        return x / y

    return (x / y) + seqSII(x + 2, y * 2)

result = seqSII(1, 1)
print("{:.2f}".format(result))
