'''
  Solução em Python - Desafio 1 - Inácio Régis Neto :::: https://onlinegdb.com/SyGTDhno7
'''

[print(cont , 'Santarem' if cont % 5 == 0 and cont % 6 == 0 else 'Santa' if cont % 5 == 0 else 'rem' if cont % 6 == 0 else '') for cont in range(1,201) ]
