//Esta clase representa los neumaticos del carro
public abstract class Tyres : Car
{
    //Defino los parametros de las llantas
    public int durabilityMax;
    public bool wet;
    public string color;

    //En este constructor defino los valores de los parametros del carro y de las llantas
    public Tyres(string A, string B, string C, string D, int lapsMax, bool isWet, string tyreColor) : base(A, B, C, D)
    {
        durabilityMax = lapsMax;
        wet = isWet;
        color = tyreColor;
    }

    //Creo el metodo de la durabilidad de los neumaticos
    public virtual string Durability(int laps)
    {
        return "the tyres have been on track " + laps + " laps";
    }

    //En este metodo veré toda la informacion del neumatico
    public void ShowAll()
    {
        Console.WriteLine($"The tyres are: {rightFront}, {leftFront}, {rightRear}, {leftRear}\nTheir durability is {durabilityMax} laps\nThe tyres are wet? {wet}\nTyres color {color}");
    }

    //Metodo para decir la presion adecuada de los neumaticos
    public abstract void PSI();
}

