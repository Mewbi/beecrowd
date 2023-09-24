comb = [0, 0, 0]

while True:
    c = int(input())

    if 1 <= c <= 4:
        if c == 4:
            break
        comb[c - 1] += 1

print("MUITO OBRIGADO")
print("Alcool: {}".format(comb[0]))
print("Gasolina: {}".format(comb[1]))
print("Diesel: {}".format(comb[2]))
