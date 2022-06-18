i = gets.chomp.to_i
puts "#{i / 365} ano(s)\n#{(i % 365) / 30} mes(es)\n#{(i % 365) % 30} dia(s)"