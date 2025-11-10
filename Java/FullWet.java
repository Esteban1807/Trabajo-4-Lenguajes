//Esta clase es para las llantas blandas
public class FullWet extends Tyres{
    //Definimos los valores con este construcctor
    public FullWet(String A, String B, String C, String D, int lapsMax, boolean isWet, String tyreColor){
        super(A, B, C, D, lapsMax, isWet, tyreColor);
    }

    //sobreescribimos el metodo durability y si lleva mas vueltas que la durabilidad maxima, llamamos al metodo de pits dentro de este
    @Override
    public String Durability(int laps){
        if (laps > durabilityMax){
            Pits();
            return "the tyres have been on track " + laps + " laps \nThe Full Wets are been changes";}
        else{
            return "the tyres have been on track " + laps + " laps\nCan continue";
        }
    }

    //Sobreescribimos el metodo de la presion determinando la presion requerida para este neumatico
    @Override
    public void PSI(){
        System.out.println("The air pressure required is 33");
    }

    //Con este metodo se estima el clima ya que estos neumaticos son para dicho clima
    public void AreWet(){
        System.out.println("The weather is heavy");
    }
}