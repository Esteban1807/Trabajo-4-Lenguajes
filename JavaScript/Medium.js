import Tyres from "./Tyres.js";

// Esta clase es para las llantas medias (Medium)
export default class Medium extends Tyres {
  // Constructor que delega en Tyres
  constructor(A, B, C, D, lapsMax, isWet, tyreColor) {
    super(A, B, C, D, lapsMax, isWet, tyreColor);
  }

  // Si supera la durabilidad máxima, entra a pits
  Durability(laps) {
    if (laps > this.durabilityMax) {
      this.Pits();
      return (
        `the tyres have been on track ${laps} laps\n` +
        `The Mediums are being changed`
      );
    }
    return `the tyres have been on track ${laps} laps\nCan continue`;
  }

  // Presión recomendada para Medium
  PSI() {
    console.log("The air pressure required is 40");
  }
}