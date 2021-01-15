using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 

       string nome = Console.ReadLine();
       double SaFixo = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
       double vendas = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
       
       double com = vendas * 0.15;
       double sTotal = SaFixo + com;
       
       Console.WriteLine("TOTAL = R$ " + sTotal.ToString("F2", CultureInfo.InvariantCulture));
    }

}
