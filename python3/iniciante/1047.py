nums = input().split(' ')
h1 = int(nums[0])
m1 = int(nums[1])
h2 = int(nums[2])
m2 = int(nums[3])

hours = h2 - h1
if hours < 0:
    hours = 24 - abs(hours)

mins = m2 - m1
if mins < 0:
    hours -= 1
    mins = 60 - abs(mins)

if hours < 0:
    hours = 24 + hours

if hours == 0 and mins == 0:
    hours = 24

print("O JOGO DUROU {} HORA(S) E {} MINUTO(S)".format(hours, mins))
