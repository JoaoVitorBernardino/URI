using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 

        int NUMBER = int.Parse(Console.ReadLine());
        int horas = int.Parse(Console.ReadLine());
        double hTrab = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
        
        double SALARY = horas * hTrab;
        
        Console.WriteLine("NUMBER = " + NUMBER);
        Console.WriteLine("SALARY = U$ " + SALARY.ToString("F2", CultureInfo.InvariantCulture));
    }

}
