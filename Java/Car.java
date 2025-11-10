public class Car {
     //Aca defino los parametros del carro (sus llantas)
    public String rightFront;
    public String leftFront;
    public String rightRear;
    public String leftRear;
    //Este es el metodo constructor que le asigna valores a los parametros
    public Car(String A, String B, String C, String D){
        this.rightFront = A;
        this.leftFront = B;
        this.rightRear = C;
        this.leftRear = D;
    }

    //Aca defino el metodo de pits que sera necesario cuando las llantas esten desgastadas
    public void Pits(){
        System.out.println("The car is entering to pits and the tyres will be change");
    }

}
