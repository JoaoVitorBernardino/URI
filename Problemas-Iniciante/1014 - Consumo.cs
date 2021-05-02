using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 

      int X = int.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
      double Y = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
      
      double media = X / Y;
      
      Console.WriteLine(media.ToString("F3", CultureInfo.InvariantCulture) + " km/l");

    }

}
