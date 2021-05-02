using System; 

class URI {

    static void Main(string[] args) { 

        string[] nums = Console.ReadLine().Split(' ');
        
        decimal A = decimal.Parse(nums[0]);
        decimal B = decimal.Parse(nums[1]);
        decimal C = decimal.Parse(nums[2]);
        
        decimal MaiorAB;
        decimal MaiorC;
        MaiorAB = (A + B + Math.Abs(A - B)) / 2;
        MaiorC = (MaiorAB + C + Math.Abs(MaiorAB - C)) / 2;
        
        Console.WriteLine(MaiorC + " eh o maior");
    }

}
