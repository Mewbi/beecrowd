import math

def calcularDistancia(x1, y1, x2, y2):
    return math.sqrt((x2 - x1)**2 + (y2 - y1)**2)

x1, y1 = list(map(float, input().split()))
x2, y2 = list(map(float, input().split()))

distancia = calcularDistancia(x1, y1, x2, y2)
print("{:.4f}".format(distancia))
