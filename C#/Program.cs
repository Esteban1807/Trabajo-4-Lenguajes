public class Program
{
    public static void Main(string[] args)
    {
        Soft Soft = new Soft("SoftA", "SoftB", "SoftC", "SoftD", 15, false, "Red");
        Soft.ShowAll();
        Soft.PSI();
        Console.WriteLine("____________________________");
        Console.WriteLine(Soft.Durability(16));
        Console.WriteLine("__________________\n");

        Medium medium = new Medium("MediumA", "MediumB", "MediumC", "MediumD", 30, false, "Yellow");
        medium.ShowAll();
        medium.PSI();
        Console.WriteLine("____________________________");
        Console.WriteLine(medium.Durability(15));
        Console.WriteLine("__________________\n");

        Hard hard = new Hard("HardA", "HardB", "HardC", "HardD", 40, false, "White");
        hard.ShowAll();
        hard.PSI();
        Console.WriteLine("____________________________");
        Console.WriteLine(hard.Durability(46));
        Console.WriteLine("__________________\n");

        FullWet full = new FullWet("FullA", "FullB", "FullC", "FullD", 25, true, "Blue");
        full.ShowAll();
        full.PSI();
        Console.WriteLine("____________________________");
        Console.WriteLine(full.Durability(23));
        Console.WriteLine("__________________\n");

        Intermedium inter = new Intermedium("InterA", "InterB", "InterC", "InterD", 40, true, "Green");
        inter.ShowAll();
        inter.PSI();
        Console.WriteLine("____________________________");
        Console.WriteLine(inter.Durability(41));
        Console.WriteLine("__________________\n");

    }
}