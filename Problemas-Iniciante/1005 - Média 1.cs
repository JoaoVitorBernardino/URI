using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 

        double A = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
        double B = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
        
        double nota1 = A * 3.5;
        double nota2 = B * 7.5;
        
        double MEDIA = (nota1 + nota2) / 11;
        
        Console.WriteLine("MEDIA = " + MEDIA.ToString("F5", CultureInfo.InvariantCulture));
    }

}
