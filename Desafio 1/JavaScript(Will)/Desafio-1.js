let numbers = []
let mult_5 = ' Santa'
let mult_6 = 'rem'

// Cria o array de 200 números
for(var i = 1; i <= 200; i++){
	numbers.push(i)
}

// Verifica o multiplo e concatena as strings
numbers.map(async(i) => {
	if(i % 5 == 0 & i % 5 == 0){
		console.log(i + mult_5 + mult_6)
	}
	else if(i % 5 == 0){
		console.log(i + mult_5)
	}
	else if (i % 6 == 0) {
		console.log( i + ' ' + mult_6)
	}
	else if (i % 5 !== 0 & i % 5 !== 0) {
		console.log(i)
	}
})
