using System; 

class URI {

    static void Main(string[] args) { 

       int seg = int.Parse(Console.ReadLine());
       int hora, min;
       
       hora = seg / 3600;
       seg = seg % 3600;
       
       min = seg / 60;
       seg = seg % 60;

        Console.WriteLine(hora + ":" + min + ":" + seg);
    }

}
