'''
	20181023
	Gabriel Cesar
	gabrielcesar2@gmail.com
'''

def factorial(i):
    if i == 0 or i == 1:
        return 1 
    else:
        return i * factorial(i - 1)

print(factorial(5))
