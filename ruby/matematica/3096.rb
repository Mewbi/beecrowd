def kamenetsky(n)
  return 0 if n < 0
  return 1 if n <= 1

  x = ((n * Math.log10(n/Math::E)) +
      (Math.log10(2*Math::PI*n) / 2.0))

  return x.floor + 1
end

n = gets.chomp.to_f
puts kamenetsky(n)
