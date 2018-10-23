var desafio1 = function(){
var santa = 'Santa';
var rem = 'rem';
    for (var i = 1; i <= 200; i++){
        if (i % 5 == 0 && i % 6 == 0 ){
            console.log(i + ' ' + santa + rem);
        } else if (i % 6 == 0){
            console.log(i + ' ' + rem);
        } else if(i % 5 == 0){
            console.log(i + ' ' + santa );
        } else {
            console.log(i);
        }
    }  
}

desafio1();
