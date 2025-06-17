# Data Type

## Scalar Types

A scalar type represents a single value. Rust has four primary scalar types: integers, floating-point numbers, Booleans,
and characters.

### Integer Types

| Length  | Signed | Unsigned |
|---------|--------|----------|
| 8-bit   | i8     | u8       |
| 16-bit  | i16    | u16      |
| 32-bit  | i32    | u32      |
| 64-bit  | i64    | u64      |
| 128-bit | i128   | u128     |
| arch    | isize  | usize    |

### Floating Point Types

| Length | Signed |
|--------|--------|
| 32-bit | f32    |
| 64-bit | f64    |

### Numeric Operations

```rust
fn main() {
    // addition
    let sum = 5 + 10;

    // subtraction
    let difference = 95.5 - 4.3;

    // multiplication
    let product = 4 * 30;

    // division
    let quotient = 56.7 / 32.2;

    // remainder
    let remainder = 43 % 5;
}
```

### Character Type

```rust
fn main() {
   let c = 'z';
   let z = 'ℤ';
   let heart_eyed_cat = '😻';
}
```

## Compound Types

### Tuple

Grouping Values into Tuples<br/>
A tuple is a general way of grouping together some number of other values with a variety of types into one compound
type.

```rust
fn main() {
    let tup: (i32, f64, u8) = (500, 6.4, 1);
}
```

```rust
fn main() {
    let tup = (500, 6.4, 1);
    let (x, y, z) = tup;

    println!("The value of y is: {}", y);
}
```

```asciidoc
The value of y is: 6.4
```

In addition to destructuring through pattern matching, we can also access a tuple element directly by using a period (.)
followed by the index of the value we want to access. For example:

```rust
fn main() {
    let x: (i32, f64, u8) = (500, 6.4, 1);
    let five_hundred = x.0;
    let six_point_four = x.1;

    let one = x.2;
}
```

### Arrays

Another way to have a collection of multiple values is with an array. Unlike a tuple, every element of an array
must-have the same type. Arrays in Rust are different than arrays in some other languages because arrays in Rust have a
fixed
length: once declared, they cannot grow or shrink in size.

```rust
fn main() {
    let a = [1, 2, 3, 4, 5];
    
    let first = a[0];
    let second = a[1];
}
```
