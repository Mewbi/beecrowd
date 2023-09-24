number = int(input())
hours = int(input())
value = float(input())

salary = hours * value
salary = "{:.2f}".format(salary)

print("NUMBER =", number)
print("SALARY = U$ " + salary)
