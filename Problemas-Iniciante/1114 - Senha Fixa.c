#include <stdio.h>
#include <stdbool.h>

int main() {
    int senha;
    bool isValida = false;

    do{
        scanf("%d", &senha);

        if(senha != 2002) {
            printf("Senha Invalida\n");
        } else {
            printf("Acesso Permitido\n");
            isValida = true;
        }
    } while(isValida != true);


    return 0;
}
