p1 = input().split()
p2 = input().split()

total = float(p1[1]) * float(p1[2]) + float(p2[1]) * float(p2[2])

total = "{:.2f}".format(total)

print("VALOR A PAGAR: R$ " + total)
