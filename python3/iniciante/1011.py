def calcularVolumeEsfera(raio):
    pi = 3.14159
    volume = (4.0/3) * pi * raio**3
    return volume

raio = float(input())

volume = calcularVolumeEsfera(raio)
volume = "{:.3f}".format(volume)

print("VOLUME = " + volume)
