def nonzero(n):
    if n == 0 or n == 1:
        return 1

    x = n // 5
    y = n % 5

    potX = x % 4
    if potX == 0:
        potX = 1
    elif potX == 1:
        potX = 2
    elif potX == 2:
        potX = 4
    elif potX == 3:
        potX = 8

    lastY = 0
    if y == 0:
        lastY = 1
    elif y == 1:
        lastY = 1
    elif y == 2:
        lastY = 2
    elif y == 3:
        lastY = 6
    elif y == 4:
        lastY = 4

    return potX * nonzero(x) * lastY

instancia = 1
while True:
    try:
        n = int(input())
        print("Instancia {}".format(instancia))
        instancia += 1

        result = nonzero(n)
        while result >= 10:
            result = result % 10

        print(result)
        print()
    except EOFError:
        break
