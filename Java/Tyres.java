//Esta clase representa los neumaticos del carro
public abstract class Tyres extends Car{
    //Defino los parametros de las llantas
    public int durabilityMax;
    public boolean wet;
    public String color;

    //En este constructor defino los valores de los parametros del carro y de las llantas
    public Tyres(String A, String B, String C, String D, int lapsMax, boolean isWet, String tyreColor){
        super(A, B, C, D);
        this.durabilityMax = lapsMax;
        this.wet = isWet;
        this.color = tyreColor;
    }

    //Creo el metodo de la durabilidad de los neumaticos
    public String Durability(int laps){
        return "the tyres have been on track " + laps + " laps";
    }

    //En este metodo veré toda la informacion del neumatico
    public void ShowAll(){
        System.out.println("The tyres are: "+rightFront+ leftFront+ rightRear+ leftRear +"\nTheir durability is "+durabilityMax+ " laps\nThe tyres are wet? "+wet+" \nTyres color "+color);
    }

    //Metodo para decir la presion adecuada de los neumaticos
    public abstract void PSI();
}

