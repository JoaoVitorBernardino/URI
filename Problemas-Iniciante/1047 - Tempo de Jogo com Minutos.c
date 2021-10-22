#include <stdio.h>

int main() {
    int horaInicial, minInicial, horaFinal, minFinal;
    int tempo, horaTotal ,minTotal;

    scanf("%d%d%d%d", &horaInicial, &minInicial, &horaFinal, &minFinal);
    
    minInicial += horaInicial * 60;
    minFinal += horaFinal * 60;

    tempo = minFinal - minInicial;

    if(tempo <= 0) {
        tempo += 24 * 60;
    }

    horaTotal = tempo / 60;
    minTotal = tempo % 60;

    printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", horaTotal, minTotal);

    return 0;
}
