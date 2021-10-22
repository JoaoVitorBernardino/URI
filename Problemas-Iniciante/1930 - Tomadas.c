#include <stdio.h>
 
int main() {
 
    int T1, T2, T3, T4;
    int maxAparelhos;

    scanf("%d%d%d%d", &T1, &T2, &T3, &T4);

    maxAparelhos = (T1 + T2 + T3 + T4) - 3;

    printf("%d\n", maxAparelhos);
 
    return 0;
}
