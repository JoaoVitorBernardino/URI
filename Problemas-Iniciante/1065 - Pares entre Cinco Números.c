#include <stdio.h>

int main() {
    int num, i;
    int result = 0;

    for(i = 0; i < 5; i++) {
        scanf("%d", &num);
        
        if(num % 2 == 0) {
            result++;
        }
    }
    
    printf("%d valores pares\n", result);

    return 0;
}
