#include <stdio.h>
 
int main() {
 
    int x = 0, result = 0;

    scanf("%d", &x);

    while(result != 6) {
        if(x % 2 != 0) {
            printf("%d\n", x);
            result++;
        }
        x++;            
    }
 
    return 0;
}
