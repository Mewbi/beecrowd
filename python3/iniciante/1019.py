def convert(t):
	h = t / 3600
	t = t % 3600
	m = t / 60
	print('{}:{}:{}'.format(int(h),int(m),int(t % 60)))

convert(int(input()))