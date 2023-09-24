grenais, golInt, golGre, inter, gremio, emp, choose = 0, 0, 0, 0, 0, 0, 0
winner = ""

while True:
    golInt, golGre = list(map(int, input().split()))

    if golInt > golGre:
        inter += 1
    elif golGre > golInt:
        gremio += 1
    else:
        emp += 1

    grenais += 1

    print("Novo grenal (1-sim 2-nao)")
    choose = int(input())

    if choose == 2:
        break

print("{} grenais".format(grenais))
print("Inter:{}".format(inter))
print("Gremio:{}".format(gremio))
print("Empates:{}".format(emp))

if inter == gremio:
    print("Nao houve vencedor")
else:
    winner = "Inter"
    if gremio > inter:
        winner = "Gremio"

    print("{} venceu mais".format(winner))
