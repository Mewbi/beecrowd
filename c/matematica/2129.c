#include <stdio.h>

typedef unsigned long long ull;

ull nonzero(ull n){
    ull x, y, potX, lastY;

    if(n == 0 || n == 1)
        return 1;

    x = n / 5;
    y = n % 5;

    potX = x % 4;
    switch(potX){
        case 0:
            potX = 1;
            break;
        
        case 1:
            potX = 2;
            break;
        
        case 2:
            potX = 4;
            break;
        
        case 3:
            potX = 8;
            break;
    }

    switch(y){
        case 0:
            lastY = 1;
            break; 
        case 1:
            lastY = 1;
            break; 
        case 2:
            lastY = 2;
            break; 
        case 3:
            lastY = 6;
            break; 
        case 4:
            lastY = 4;
            break; 
        default:
            lastY = 0;
    }

    return (potX * nonzero(x) * lastY);
}

int main(){
    ull n, result, instancia = 1;

    while(scanf("%llu", &n) != EOF){
        printf("Instancia %llu\n", instancia); 
        instancia++;

        result = nonzero(n);
        while(result >= 10)
            result = result % 10;

        printf("%llu\n\n", result);
    }
    return 0;
}
