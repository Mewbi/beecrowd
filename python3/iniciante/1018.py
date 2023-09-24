def qnt(v):
	print(v)
	for i in notas:
		print('{} nota(s) de R$ {},00'.format(int(v / i), i))
		v = v % i

notas = [100, 50, 20, 10, 5, 2, 1]

qnt(int(input()))