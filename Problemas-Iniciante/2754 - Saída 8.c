#include <stdio.h>

int main() {
   double x, y;
   
   x = 234.345;
   y = 45.698;

   printf("%.6f - %.6f\n", x, y);
   printf("%.0f - %.0f\n", x, y);
   printf("%.1f - %.1f\n", x, y);
   printf("%.2f - %.2f\n", x, y);
   printf("%.3f - %.3f\n", x, y);
   printf("%e - %e\n", x, y);
   printf("%E - %E\n", x, y);
   printf("%G - %G\n", x, y);
   printf("%g - %g\n", x, y);
    
    return 0;
}
