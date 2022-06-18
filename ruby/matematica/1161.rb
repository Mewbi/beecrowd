def fatorial (n)
	(1..n).inject(:*) || 1
end

while num = gets
	num = num.split(' ')
	puts "#{fatorial(num[0].to_i) + fatorial(num[1].to_i)}"
end