def price(t, q):
	print('Total: R$ {:.2f}'.format(p[t] * q))

p = {1: 4, 2: 4.5, 3: 5, 4: 2, 5: 1.5};

numbers = [int(n) for n in input().split()]

price(numbers[0], numbers[1])