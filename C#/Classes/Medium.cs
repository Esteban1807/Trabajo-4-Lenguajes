//Esta clase es para las llantas medias
public class Medium : Tyres
{
    //Definimos los valores con este constuctor
    public Medium(string A, string B, string C, string D, int lapsMax, bool isWet, string tyreColor) : base(A, B, C, D, lapsMax, isWet, tyreColor)
    {
    }

    //sobreescribimos el metodo durability y si lleva mas vueltas que la durabilidad maxima, llamamos al metodo de pits dentro de este
    public override string Durability(int laps)
    {
        if (laps > durabilityMax)
        {
            Pits();
            return "the tyres have been on track " + laps + " laps \nThe Mediums are been changes";
        }
        else
        {
            return "the tyres have been on track " + laps + " laps\nCan continue";
        }
    }

    //Sobreescribimos el metodo de la presion determinando la presion requerida para este neumatico
    public override void PSI()
    {
        Console.WriteLine("The air pressure required is 40");
    }
}