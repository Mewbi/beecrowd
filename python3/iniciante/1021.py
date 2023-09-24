def notas(v, n):
	print('NOTAS:')
	for i in n:
		print('{} nota(s) de R$ {:.2f}'.format(int(v / i), i))
		v = round(v % i, 2)
	return v

def moedas(v, m):
	print('MOEDAS:')
	for i in m:
		print('{} moeda(s) de R$ {:.2f}'.format(int(v / i), i))
		v = round(v % i, 2)


n = [100.00, 50.00, 20.00, 10.00, 5.00, 2.00] 
m = [1.00, 0.50, 0.25, 0.10, 0.05, 0.01]

v = float(input())


moedas(notas(v, n), m)