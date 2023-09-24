n = 0
pos = 0
maior = 0

for i in range(1, 101):
    n = int(input())
    if n > maior:
        maior = n
        pos = i

print("{}\n{}".format(maior, pos))
