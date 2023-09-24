n = int(input())
qtd, tot = 0, 0
cob = {
    "C": 0,
    "R": 0,
    "S": 0,
}

for _ in range(n):
    line = input().split()

    qtd = int(line[0])
    animal = line[1]

    cob[animal] += qtd
    tot += qtd

print("Total: {} cobaias".format(tot))
print("Total de coelhos: {}".format(cob['C']))
print("Total de ratos: {}".format(cob['R']))
print("Total de sapos: {}".format(cob['S']))
print("Percentual de coelhos: {:.2f} %".format(cob['C'] / tot * 100))
print("Percentual de ratos: {:.2f} %".format(cob['R'] / tot * 100))
print("Percentual de sapos: {:.2f} %".format(cob['S'] / tot * 100))
