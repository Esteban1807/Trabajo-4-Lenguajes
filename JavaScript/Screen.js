import Soft from "./Soft.js";

import Medium from "./Medium.js";

import Hard from "./Hard.js";

import Intermedium from "./Intermedium.js";

import FullWet from "./FullWet.js";

const soft = new Soft("SoftA", "SoftB", "SoftC", "SoftD", 15, false, "Red");
soft.ShowAll();
soft.PSI();
console.log("____________________________");
console.log(soft.Durability(16));
console.log("__________________\n");

const medium = new Medium("MediumA", "MediumB", "MediumC", "MediumD", 30, false, "Yellow");
medium.ShowAll();
medium.PSI();
console.log("____________________________");
console.log(medium.Durability(15));
console.log("__________________\n");

const hard = new Hard("HardA", "HardB", "HardC", "HardD", 40, false, "White");
hard.ShowAll();
hard.PSI();
console.log("____________________________");
console.log(hard.Durability(46));
console.log("__________________\n");

const full = new FullWet("FullA", "FullB", "FullC", "FullD", 25, true, "Blue");
full.ShowAll();
full.PSI();
console.log("____________________________");
console.log(full.Durability(23));
console.log("__________________\n");

const inter = new Intermedium("InterA", "InterB", "InterC", "InterD", 40, true, "Green");
inter.ShowAll();
inter.PSI();
console.log("____________________________");
console.log(inter.Durability(41));
console.log("__________________\n");