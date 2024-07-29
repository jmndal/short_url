// FUNDAMENTALS

/* The ANY Type */
let sales = 123_456_789; // Large number (eg. 1,000 up) can use `_` to make code readable
let course = 'TypeScript';
let is_published = true;
let level; // Better use annotation

/* ARRAY */
let numbers: number[] = []; // For numbers
let books: string[] = []; // For strings

/* TUPLES */
let user: [number, string] = [1, 'Test']; // Useful if only have two values, the key & value pairs

/* ENUMS */
// If we add `const` before `enum`, the compiler will compile a more optimized code
enum Size { Small = 0, Medium, Large }; // By default tsc compiler assign first value as 0 and others as 1, 2, and so on, but you can set the value.
let mySize: Size = Size.Medium;         // But if you want to set the value as a string you need to set all the value as string
console.log(mySize);

/* FUNCTIONS */
function calculateTax(income: number, taxYear = 2024): number { // Best practice: Always ANNOTATE functions (number, string, void)
  if (taxYear < 2024)
    return income * 1.2;
  return income * 1.3;
};

calculateTax(10_000, 2024);

/* OBJECT */ 
let employee: Employee = { 
  id: 1, 
  name: 'Test', 
  retire: (date: Date) => { // `date` is method, 
    console.log(date);
  }
}; // Can add `readonly` before and `?` after a property

/* ---------------------------------------------------------------------- */

// ADVANCE TYPES

/* TYPE ALIASES */
type Employee = {
  readonly id: number;
  name: string;
  retire: (date: Date) => void;
};

/* UNION */
function kgToLbs(weight: number | string): number {
  // Narrowing
  if (typeof weight === 'number')
    return weight * 2.2;
  else
    return parseInt(weight) * 2.2;
};

kgToLbs(10);
kgToLbs('10kg');

/* INTERSECTION */
type Draggable = {
  drag: () => void
};

type Resizable = {
  resize: () => void
};

type UIWidget = Draggable & Resizable;

let textBox: UIWidget = {
  drag: () => {},
  resize: () => {},
};

/* LITERAL */
// exact, specific value
type Quantity = 50 | 100;
type Metric = 'cm' | 'inch';
let quantity: Quantity = 100;
let metric: Metric = 'cm';

/* NULLABLE */ 
function greet(name: string | null | undefined) {
  if (name)
    console.log(name.toUpperCase());
  else
    console.log('Annyeong!');
};

greet(null);

/* OPTIONAL CHAINING */
type Customer = {
  birthday?: Date,
};

function getCustomer(id: number): Customer | null | undefined {
  return id === 0 ? null : {
    birthday: new Date()
  };
};

let customer = getCustomer(0);

// Optional property access operator `?.`
console.log(customer?.birthday?.getFullYear);

// Optional element access operator
// customers?.[0];

// Optional call
let log: any = null;
log?.('a');