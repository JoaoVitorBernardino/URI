#include <stdio.h>
#include <math.h>

int main() {
    double delta, raiz;
    double A, B, C;
    double R1, R2;

    scanf("%lf %lf %lf", &A, &B, &C);
    
    delta = (B * B)-(4 * A * C);
    
	R1 = (-B + sqrt(delta)) / (2 * A);
    R2 = (-B - sqrt(delta))/ (2 * A);
        
    if((R1 > 0 && R2 > 0) || (R1 < 0 && R2 < 0)) {
        printf("R1 = %.5f\n", R1);
        printf("R2 = %.5f\n", R2);
    
    } else {
        
        printf("Impossivel calcular\n");
       
    }

    

    return 0;
}
