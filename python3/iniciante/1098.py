for i in range(0, 21, 2):
    for j in range(10 + i, i + 31, 10):
        if i % 10 != 0:
            print("I={:.1f} J={:.1f}".format(i/10,j/10))
        else:
            print("I={} J={}".format(i//10, j//10))
