t = gets.chomp.to_i
t.times {
	num = gets.chomp.split(' ')
	n1 = num[0]
	n2 = num[1]
	if "#{n1[((n1.size - n2.size)),n1.size]}" == "#{n2}"
		puts "encaixa"
	else
		puts "nao encaixa"
	end
}