def calcularAreaTrianguloRetangulo(base, altura):
    area = (base * altura) / 2
    return area

def calcularAreaCirculo(raio):
    pi = 3.14159
    area = pi * raio**2
    return area

def calcularAreaTrapezio(base1, base2, altura):
    area = ((base1 + base2) * altura) / 2
    return area

def calcularAreaQuadrado(lado):
    area = lado**2
    return area

def calcularAreaRetangulo(lado1, lado2):
    area = lado1 * lado2
    return area

line = input().split()
A = float(line[0])
B = float(line[1])
C = float(line[2])

areaTriangulo = calcularAreaTrianguloRetangulo(A, C)
areaCirculo = calcularAreaCirculo(C)
areaTrapezio = calcularAreaTrapezio(A, B, C)
areaQuadrado = calcularAreaQuadrado(B)
areaRetangulo = calcularAreaRetangulo(A, B)

print("TRIANGULO: {:.3f}".format(areaTriangulo))
print("CIRCULO: {:.3f}".format(areaCirculo))
print("TRAPEZIO: {:.3f}".format(areaTrapezio))
print("QUADRADO: {:.3f}".format(areaQuadrado))
print("RETANGULO: {:.3f}".format(areaRetangulo))
