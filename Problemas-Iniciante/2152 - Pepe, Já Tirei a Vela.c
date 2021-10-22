#include <stdio.h>

int main() {
    int qtdTestes, i;
    int H, M, O;

    scanf("%d", &qtdTestes);

    for(i = 0; i < qtdTestes; i++) {
        scanf("%d %d %d", &H, &M, &O);

        if(H < 10 && M < 10) {
            if(O == 0) {
                printf("0%d:0%d - A porta fechou!\n", H, M);
            } else {
                printf("0%d:0%d - A porta abriu!\n", H, M);
            }
        } else if(H >= 10 && M < 10) {
            if(O == 0) {
                printf("%d:0%d - A porta fechou!\n", H, M);
            } else {
                printf("%d:0%d - A porta abriu!\n", H, M);
            }
        } else if(H < 10 && M >= 10) {
            if(O == 0) {
                printf("0%d:%d - A porta fechou!\n", H, M);
            } else {
                printf("0%d:%d - A porta abriu!\n", H, M);
            }
        } else {
            if(O == 0) {
                printf("%d:%d - A porta fechou!\n", H, M);
            } else {
                printf("%d:%d - A porta abriu!\n", H, M);
            }
        }
    }

    return 0;
}
