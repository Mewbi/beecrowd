def print_and_clean(arr, t):
    for i, v in enumerate(arr):
        print("{}[{}] = {}".format(t, i, v))
    return []

imp = []
par = []

for _ in range(15):
    n = int(input())

    if n % 2 == 0:
        par.append(n)
    else:
        imp.append(n)

    if len(imp) == 5:
        imp = print_and_clean(imp, "impar")

    if len(par) == 5:
        par = print_and_clean(par, "par")

print_and_clean(imp, "impar")
print_and_clean(par, "par")
