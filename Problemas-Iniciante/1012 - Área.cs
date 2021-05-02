using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 

        string[] v = Console.ReadLine().Split(' ');
            double A = double.Parse(v[0], CultureInfo.InvariantCulture);
            double B = double.Parse(v[1], CultureInfo.InvariantCulture);
            double C = double.Parse(v[2], CultureInfo.InvariantCulture);

            double areaTR = (A * C) / 2;
            double areaC = Math.Pow(C, 2.0) * 3.14159;
            double areaTra = (A + B) * C / 2;
            double areaQ = Math.Pow(B, 2.0);
            double areaR = A * B;

            Console.WriteLine("TRIANGULO: " + areaTR.ToString("F3", CultureInfo.InvariantCulture));
            Console.WriteLine("CIRCULO: " + areaC.ToString("F3", CultureInfo.InvariantCulture));
            Console.WriteLine("TRAPEZIO: " + areaTra.ToString("F3", CultureInfo.InvariantCulture));
            Console.WriteLine("QUADRADO: " + areaQ.ToString("F3", CultureInfo.InvariantCulture));
            Console.WriteLine("RETANGULO: " + areaR.ToString("F3", CultureInfo.InvariantCulture));

    }

}
