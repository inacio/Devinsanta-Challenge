# Solution for the challenge 1 in python by Iago Lisboa

class Challenge1:
    y = "santa"
    z = "rem"
    for x in range(1, 201):
        if x % 5 == 0 and x % 6 != 0:
            print(x, y)
        elif x % 6 == 0 and x % 5 != 0:
            print(x, z)
        elif x % 5 == 0 and x % 6 == 0:
            print(x, y+z)
        else:
            print(x)
