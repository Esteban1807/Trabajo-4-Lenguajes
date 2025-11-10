//Esta clase es para las llantas medias
public class Medium extends Tyres{
    //Definimos los valores con este construcctor
    public Medium(String A, String B, String C, String D, int lapsMax, boolean isWet, String tyreColor){
        super(A, B, C, D, lapsMax, isWet, tyreColor);
    }

    //sobreescribimos el metodo durability y si lleva mas vueltas que la durabilidad maxima, llamamos al metodo de pits dentro de este
    @Override
    public String Durability(int laps){
        if (laps > durabilityMax){
            Pits();
            return "the tyres have been on track " + laps + " laps \nThe Mediums are been changes";}
        else{
            return "the tyres have been on track " + laps + " laps\nCan continue";
        }
    }

    //Sobreescribimos el metodo de la presion determinando la presion requerida para este neumatico
    @Override
    public void PSI(){
        System.out.println("The air pressure required is 40");
    }
}