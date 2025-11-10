export default class Car {
  // Constructor que asigna los valores a las llantas
  constructor(A, B, C, D) {
    this.rightFront = A;
    this.leftFront = B;
    this.rightRear = C;
    this.leftRear = D;
  }

  // Método para entrar a pits
  Pits() {
    console.log("The car is entering to pits and the tyres will be changed");
  }
}