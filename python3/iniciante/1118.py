close = False

while not close:
    count = 0
    sum = 0.0
    choose = 0

    while count < 2:
        v = float(input())

        if v < 0 or v > 10:
            print("nota invalida")
        else:
            sum += v
            count += 1

    print('media = {:.2f}'.format(sum / 2))
    count = 0
    sum = 0

    while True:
        print("novo calculo (1-sim 2-nao)")
        choose = int(input())

        if choose == 2 or choose == 1:
            if choose == 2:
                close = True

            break  # Exit the inner loop
