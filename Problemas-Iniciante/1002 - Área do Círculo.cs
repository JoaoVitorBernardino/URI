using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 

        double raio = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
        
        double area = 3.14159 * Math.Pow(raio, 2.0);
        
        Console.WriteLine("A=" + area.ToString("F4", CultureInfo.InvariantCulture));
    }

}
