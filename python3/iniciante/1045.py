def sort(arr):
    for i in range(len(arr) - 1):
        for j in range(len(arr) - i - 1):
            if arr[j] < arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
    return arr

l = list(map(float, input().split()))

l = sort(l)
a, b, c = l[0], l[1], l[2]

if a >= b + c:
    print("NAO FORMA TRIANGULO")
else:
    if a ** 2 == b ** 2 + c ** 2:
        print("TRIANGULO RETANGULO")

    if a ** 2 > b ** 2 + c ** 2:
        print("TRIANGULO OBTUSANGULO")

    if a ** 2 < b ** 2 + c ** 2:
        print("TRIANGULO ACUTANGULO")

    if a == b == c:
        print("TRIANGULO EQUILATERO")

    if (a == b and a != c) or (a == c and a != b) or (b == c and b != a):
        print("TRIANGULO ISOSCELES")
