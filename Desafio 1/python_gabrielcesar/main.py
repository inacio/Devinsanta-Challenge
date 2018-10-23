'''
    20181022
    Gabriel Cesar
    gabrielcesar2@gmail.com
'''

for i in range(1, 201):
    if i % 5 == 0 and i % 6 == 0:
        print(i, 'Santarem')
    elif i % 5 == 0:
        print(i, 'Santa')
    elif i % 6 == 0:
        print(i, 'rem')
    else:
        print(i)

