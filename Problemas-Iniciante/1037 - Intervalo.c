#include <stdio.h>

int main() {
    float intervalo;

    scanf("%f", &intervalo);

    if(intervalo >= 0.0 && intervalo <= 25.0) {
        printf("Intervalo [0,25]\n");
    } else if(intervalo > 25.0 && intervalo <= 50.0) {
        printf("Intervalo (25,50]\n");
    } else if(intervalo > 50.0 && intervalo <= 75.0) {
        printf("Intervalo (50,75]\n");
    } else if(intervalo > 75.0 && intervalo <= 100.0) {
        printf("Intervalo (75,100]\n");
    } else {
        printf("Fora de intervalo\n");
    }

    return 0;
}
