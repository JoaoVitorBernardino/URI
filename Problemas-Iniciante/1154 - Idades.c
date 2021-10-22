#include <stdio.h>

int main() {

    int idades = 0, soma = 0, div = 0;
    float media;

    while(idades >= 0) {
        scanf("%d", &idades);

        if(idades >= 0){
            soma += idades;
            div++;
        }
    }

    media = (float)soma / div;

    printf("%.2f\n", media);

    return 0;
}
