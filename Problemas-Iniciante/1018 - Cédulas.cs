using System; 

class URI {

    static void Main(string[] args) { 

        int valor = int.Parse(Console.ReadLine());
         
        Console.WriteLine(valor);
        Console.WriteLine(valor / 100 + " nota(s) de R$ 100,00");
        valor = valor % 100;
        Console.WriteLine(valor / 50 + " nota(s) de R$ 50,00");
        valor = valor % 50;
        Console.WriteLine(valor / 20 + " nota(s) de R$ 20,00");
        valor = valor % 20;
        Console.WriteLine(valor / 10 + " nota(s) de R$ 10,00");
        valor = valor % 10;
        Console.WriteLine(valor / 5 + " nota(s) de R$ 5,00");
        valor = valor % 5;
        Console.WriteLine(valor / 2 + " nota(s) de R$ 2,00");
        valor = valor % 2;
        Console.WriteLine(valor / 1 + " nota(s) de R$ 1,00"); 
    }

}
