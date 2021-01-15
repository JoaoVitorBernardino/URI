using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 

       double R = double.Parse(Console.ReadLine(), CultureInfo.InvariantCulture);
       double volume = (4.0/3) * 3.14159 * Math.Pow(R, 3.0);
       
       Console.WriteLine("VOLUME = " + volume.ToString("F3", CultureInfo.InvariantCulture));
    }

}
