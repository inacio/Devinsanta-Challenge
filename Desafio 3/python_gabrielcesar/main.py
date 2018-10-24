'''
	20181023
	Gabriel Cesar
	gabrielcesar2@gmail.com
'''
import sys
sys.setrecursionlimit(1000000)

def factorial(i):
    if i == 0 or i <= 1:
        return 1 
    else:
        return i * factorial(i - 1)

j = []

for i in range(1, 301):
    j.append(i ** 2)

for i in j:
	print(factorial(i))

