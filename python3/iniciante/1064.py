p = 0
n = 0
t = 0

for i in range(1, 7):
    n = float(input())

    if n > 0:
        p += 1
        t += n

print("{} valores positivos".format(p))

if p > 0:
    print("{:.1f}".format(t/p))
else:
    print("Nenhum valor positivo foi inserido")
