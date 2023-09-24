def check_numbers(a, b, c, d):
	if (b > c) and (d > a) and ((c + d) > (a + b)) and (c > 0) and (d > 0) and ((a % 2) == 0):
		print('Valores aceitos')
	else:
		print('Valores nao aceitos')

numbers = [int(n) for n in input().split()]

check_numbers(numbers[0], numbers[1], numbers[2], numbers[3])