public class Main {
    public static void main(String[] args) {
        Soft Soft = new Soft("SoftA", "SoftB", "SoftC", "SoftD", 15, false, "Red");
        Soft.ShowAll();
        Soft.PSI();
        System.out.println("____________________________");
        System.out.println(Soft.Durability(16));
        System.out.println("__________________\n");

        Medium medium = new Medium("MediumA", "MediumB", "MediumC", "MediumD", 30, false, "Yellow");
        medium.ShowAll();
        medium.PSI();
        System.out.println("____________________________");
        System.out.println(medium.Durability(15));
        System.out.println("__________________\n");

        Hard hard = new Hard("HardA", "HardB", "HardC", "HardD", 40, false, "White");
        hard.ShowAll();
        hard.PSI();
        System.out.println("____________________________");
        System.out.println(hard.Durability(46));
        System.out.println("__________________\n");

        FullWet full = new FullWet("FullA", "FullB", "FullC", "FullD", 25, true, "Blue");
        full.ShowAll();
        full.PSI();
        System.out.println("____________________________");
        System.out.println(full.Durability(23));
        System.out.println("__________________\n");

        Intermedium inter = new Intermedium("InterA", "InterB", "InterC", "InterD", 40, true, "Green");
        inter.ShowAll();
        inter.PSI();
        System.out.println("____________________________");
        System.out.println(inter.Durability(41));
        System.out.println("__________________\n");
    }
}
