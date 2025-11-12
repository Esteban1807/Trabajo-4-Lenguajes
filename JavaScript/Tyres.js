import Car from "./Car.js";

// Esta clase representa los neumáticos del carro
export default class Tyres extends Car {
  // Constructor de los neumáticos
  constructor(A, B, C, D, lapsMax, isWet, tyreColor) {
    super(A, B, C, D);
    this.durabilityMax = lapsMax;
    this.wet = isWet;
    this.color = tyreColor;
  }

  // Método de durabilidad base (puede sobreescribirse)
  Durability(laps) {
    return `the tyres have been on track ${laps} laps`;
  }

  // Mostrar toda la información del neumático
  ShowAll() {
    console.log(
      `The tyres are: ${this.rightFront}, ${this.leftFront}, ${this.rightRear}, ${this.leftRear}\nTheir durability is ${this.durabilityMax} laps\nThe tyres are wet? ${this.wet}\nTyres color: ${this.color}`
    );
  }

  // Método para la presión (a implementar por cada compuesto)
  PSI() {
    console.log("NON");
  }
}