while True:
    n, m, p = map(int, input().split())
    if n == 0 and m == 0 and p == 0:
        break

    vetor = [[0] * (n + 1) for _ in range(m + 1)]
    for i in range(1, m + 1):
        vetor[i] = [0] + list(map(int, input().split()))

    resp = 0
    for i in range(1, n + 1):
        jaforam = 0
        for j in range(1, m + 1):
            if vetor[j][i] == 1:
                jaforam += 1
            else:
                if jaforam >= p:
                    resp += 1
                jaforam = 0
        if jaforam >= p:
            resp += 1

    print(resp)
