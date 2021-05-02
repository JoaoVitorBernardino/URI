using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 

            string[] p1 = Console.ReadLine().Split(' ');
            int codigo1 = int.Parse(p1[0]);
            int num1 = int.Parse(p1[1]);
            double vUnit1 = double.Parse(p1[2], CultureInfo.InvariantCulture);

            double VP1 = num1 * vUnit1;

            string[] p2 = Console.ReadLine().Split(' ');
            int codigo2 = int.Parse(p2[0]);
            int num2 = int.Parse(p2[1]);
            double vUnit2 = double.Parse(p2[2], CultureInfo.InvariantCulture);

            double VP2 = num2 * vUnit2;

            double total = VP1 + VP2;

            Console.WriteLine("VALOR A PAGAR: R$ " + total.ToString("F2", CultureInfo.InvariantCulture));

    }

}
