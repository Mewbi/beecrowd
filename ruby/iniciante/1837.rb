def euclidean_division(a, b)
  limit = (b < 0) ? -b : b
  r = 0
  q = 0
  ret = 0

  (0..limit - 1).each do |r|
    ret = r
    if (a - r) % b == 0
      q = (a - r) / b
      break
    end
  end

  [q, ret]
end

a, b = gets.chomp.split(' ').map(&:to_i)
q, r = euclidean_division(a, b)
puts "#{q} #{r}"
