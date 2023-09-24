def calculate_years(pa, pb, g1, g2):
    anos = 0
    while anos <= 100 and pa <= pb:
        anos += 1
        pa += int(pa * (g1 / 100))
        pb += int(pb * (g2 / 100))

    if anos > 100:
        return "Mais de 1 seculo."
    else:
        return str(anos) + " anos."

n = int(input())

for _ in range(n):
    pa, pb, g1, g2 = map(float, input().split())
    result = calculate_years(pa, pb, g1, g2)
    print(result)
