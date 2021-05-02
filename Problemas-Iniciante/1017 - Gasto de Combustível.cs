using System; 
using System.Globalization;

class URI {

    static void Main(string[] args) { 

        int horas = int.Parse(Console.ReadLine());
        int media = int.Parse(Console.ReadLine());
        
        double mult = horas * media;
        double consumo =  mult / 12;
        
        Console.WriteLine(consumo.ToString("F3"), CultureInfo.InvariantCulture);
        
    }

}
