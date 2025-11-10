//esta clase representa un carro con 4 llantas
public class Car
{
    //Aca defino los parametros del carro (sus llantas)
    public string rightFront;
    public string leftFront;
    public string rightRear;
    public string leftRear;
    //Este es el metodo constructor que le asigna valores a los parametros
    public Car(string A, string B, string C, string D)
    {
        rightFront = A;
        leftFront = B;
        rightRear = C;
        leftRear = D;
    }

    //Aca defino el metodo de pits que sera necesario cuando las llantas esten desgastadas
    public void Pits()
    {
        Console.WriteLine("The car is entering to pits and the tyres will be change");
    }

}