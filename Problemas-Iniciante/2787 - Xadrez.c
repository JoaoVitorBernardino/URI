#include <stdio.h>

int main() {
    int L, C;

    scanf("%d", &L);
    scanf("%d", &C);

    if(L % 2 == 0 && C % 2 == 0) {
        printf("1\n");
    } else if(L % 2 == 0 && C % 2 != 0) {
        printf("0\n");
    } else if(L % 2 != 0 && C % 2 == 0) {
        printf("0\n");
    } else {
        printf("1\n");
    }


    return 0;
}
