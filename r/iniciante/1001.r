# Read multiple lines from stdin (must use 'r')
input <- file('stdin', 'r')

a <- as.integer(readLines(input, n=1))
b <- as.integer(readLines(input, n=1))

soma = a + b

# Write to stdout (must use '')
write(paste("X =", soma), '')
