//Esta clase es para las llantas intermedias
public class Intermedium : Tyres
{
    //Definimos los valores con este constuctor
    public Intermedium(string A, string B, string C, string D, int lapsMax, bool isWet, string tyreColor) : base(A, B, C, D, lapsMax, isWet, tyreColor)
    {
    }

    //Con este metodo se estima el clima ya que estos neumaticos son para dicho clima
    public void AreWet()
    {
        Console.WriteLine("The weather is medium");
    }

    //Sobreescribimos el metodo de la presion determinando la presion requerida para este neumatico
    public override void PSI()
    {
        Console.WriteLine("The air pressure required is 37");
    }

    //sobreescribimos el metodo durability y si lleva mas vueltas que la durabilidad maxima, llamamos al metodo de pits dentro de este
    public override string Durability(int laps)
    {
        if (laps > durabilityMax)
        {
            Pits();
            return "the tyres have been on track " + laps + " laps \nThe Intermediums are been changes";
        }
        else
        {
            return "the tyres have been on track " + laps + " laps\nCan continue";
        }
    }
}