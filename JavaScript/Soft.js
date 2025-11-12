import Tyres from "./Tyres.js";

// Esta clase es para las llantas blandas (Soft)
export default class Soft extends Tyres {
  // Constructor que delega en Tyres
  constructor(A, B, C, D, lapsMax, isWet, tyreColor) {
    super(A, B, C, D, lapsMax, isWet, tyreColor);
  }

  // Si supera la durabilidad máxima, entra a pits
  Durability(laps) {
    if (laps > this.durabilityMax) {
      this.Pits();
      return (
        `the tyres have been on track ${laps} laps\nThe Softs are being changed`
      );
    }
    return `the tyres have been on track ${laps} laps\nCan continue`;
  }

  // Presión recomendada para Soft
  PSI() {
    console.log("The air pressure required is 34");
  }
}