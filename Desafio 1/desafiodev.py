def desafio():
    contador = 0
    for contador in range(1, 201):
        if (contador == 0):
            print(contador)
        elif (contador % 6 == 0 and contador % 5 == 0):
            print("{} Santarem".format(contador))
        elif (contador % 5 == 0):
            print("{} Santa".format(contador))
        elif (contador % 6 == 0):
            print("{} rem".format(contador))
        else:
            print(contador)
desafio()