#include <stdio.h>

int main() {
    float N1, N2, N3, N4, notaExame, media, mediaFinal;

    scanf("%f%f%f%f", &N1, &N2, &N3, &N4);

    media = ((N1 * 2) + (N2 * 3) + (N3 * 4) + N4) / 10;

    if(media >= 7.0) {
        printf("Media: %.1f\n", media);
        printf("Aluno aprovado.\n");
    } else if(media < 5.0) {
        printf("Media: %.1f\n", media);
        printf("Aluno reprovado.\n");
    } else {
        printf("Media: %.1f\n", media);
        printf("Aluno em exame.\n");
        scanf("%f", &notaExame);
        printf("Nota do exame: %.1f\n", notaExame);

        mediaFinal = (media + notaExame) / 2;

        if(mediaFinal >= 5.0) {
            printf("Aluno aprovado.\n");
            printf("Media final: %.1f\n", mediaFinal);
        } else {
            printf("Aluno reprovado.\n");   
            printf("Media final: %.1f\n", mediaFinal);
        }
    }

    return 0;
}
