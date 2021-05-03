using System; 

class URI {

    static void Main(string[] args) { 
        
        int dias = int.Parse(Console.ReadLine());
        int ano, mes;
        
        ano = dias / 365;
        dias = dias % 365;
        
        mes = dias /30;
        dias = dias % 30;
        
        Console.WriteLine(ano + " ano(s)");
        Console.WriteLine(mes + " mes(es)");
        Console.WriteLine(dias + " dia(s)");
    }

}
