def seqS(n):
    if n == 1:
        return n
    return 1/n + seqS(n - 1)

result = seqS(100)
print("{:.2f}".format(result))
