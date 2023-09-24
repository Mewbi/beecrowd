a, b = list(map(int, input().split()))

if a >= b:
    b += 24

print("O JOGO DUROU {} HORA(S)".format(b - a))
