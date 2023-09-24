l1 = input().split()
l2 = input().split()
l3 = input().split()
l4 = input().split()

d1 = int(l1[1])
h1 = int(l2[0])
m1 = int(l2[2])
s1 = int(l2[4])

d2 = int(l3[1])
h2 = int(l4[0])
m2 = int(l4[2])
s2 = int(l4[4])

dias = d2 - d1
horas = h2 - h1
minutos = m2 - m1
segundos = s2 - s1

if segundos < 0:
    segundos += 60
    minutos -= 1

if minutos < 0:
    minutos += 60
    horas -= 1

if horas < 0:
    horas += 24
    dias -= 1

print("{} dia(s)\n{} hora(s)\n{} minuto(s)\n{} segundo(s)".format(dias, horas, minutos, segundos))
