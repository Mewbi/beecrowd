count = 0
sum = 0.0

while True:
    v = float(input())

    if v < 0 or v > 10:
        print("nota invalida")
    else:
        sum += v
        count += 1

    if count == 2:
        break

print('media = {:.2f}'.format(sum / 2))
