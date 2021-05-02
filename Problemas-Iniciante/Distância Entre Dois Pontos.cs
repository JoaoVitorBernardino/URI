using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 
        
        string[] x1y1 = Console.ReadLine().Split(' ');
        
        float x1 = float.Parse(x1y1[0], CultureInfo.InvariantCulture);
        float y1 = float.Parse(x1y1[1], CultureInfo.InvariantCulture);
        
        string[] x2y2 = Console.ReadLine().Split(' ');
        
        float x2 = float.Parse(x2y2[0]);
        float y2 = float.Parse(x2y2[1]);
        
        double r1, r2;
        r1 = Math.Pow((x2 - x1), 2);
        r2 = Math.Pow((y2 - y1), 2);
        
        double Distancia;
        Distancia = Math.Sqrt(r1 + r2);
        Console.WriteLine(Distancia.ToString("F4", CultureInfo.InvariantCulture));
        
    }

}
