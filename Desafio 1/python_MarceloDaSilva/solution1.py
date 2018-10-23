

for ind in range(1,201):
    if(ind % 5 == 0 and ind % 6 == 0):
        result = f'{ind} Santarem'
    elif(ind % 5 == 0):
        result = f'{ind} Santa'
    elif(ind % 6 == 0):
        result = f'{ind} rem'
    else:
        result = ind

    print(result)