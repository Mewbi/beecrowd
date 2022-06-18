def media(a, b, c, d):
	m = (a+b+c+d) / 10
	return m

def checar(media, aprov, reprov=None):
	if (reprov==None):
		r = 'Aluno aprovado.' if media >= aprov else 'Aluno reprovado.'
		print(r)
	else:
		if (media >= aprov):
			print('Aluno aprovado.')
		elif (media < reprov):
			print('Aluno reprovado.')


n = [float(nums) for nums in input().split()]

m = media(n[0]*2, n[1]*3, n[2]*4, n[3]*1)
print('Media: {:.1f}'.format(m))

if (5 <= m <= 6.9):
	print('Aluno em exame.')

	e = float(input())
	print('Nota do exame: {:.1f}'.format(e))

	checar((m + e) / 2, 5)
	print('Media final: {:.1f}'.format((m + e) / 2))

else:
	checar(m, 7, 5)