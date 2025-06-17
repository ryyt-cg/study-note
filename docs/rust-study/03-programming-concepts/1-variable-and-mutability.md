# Variable & Data Type

Rust has several signed integer types of known size:

* i8
* i16
* i32
* i64
* i128

unsigned integer types of know size:

* u8
* u16
* u32
* u64
* u128

Rust has platform-specific integer types

* isize
* usize

these types are optimized for the natural word size of your platform

* for maximum efficiency

## Rust naming convention

For details click [here](https://rust-lang.github.io/api-guidelines/naming.html)

In general, Rust tends to use UpperCamelCase for "type-level" constructs (types and traits) and snake_case for "
value-level" constructs. More precisely:

| Item                    | Convention                                                 |
|-------------------------|------------------------------------------------------------|
| Crates                  | unclear                                                    |
| Modules                 | snake_case                                                 |
| Types                   | UpperCamelCase                                             |
| Traits                  | UpperCamelCase                                             |
| Enum variants           | UpperCamelCase                                             |
| Functions               | snake_case                                                 |
| Methods                 | snake_case                                                 |
| General constructors    | new or with_more_details                                   |
| Conversion constructors | from_some_other_type                                       |
| Macros                  | snake_case                                                 |
| Local variables         | snake_case                                                 |
| Statics                 | SCREAMING_SNAKE_CASE                                       |
| Constants               | SCREAMING_SNAKE_CASE                                       |
| Type parameters         | concise UpperCamelCase, usually single uppercase letter: T |
| Lifetimes               | short lowercase, usually a single letter: 'a, 'de, 'src    |
| Features                | unclear but see C-FEATURE                                  |

In Rust Language, **by default variables are immutable.** This is one of many nudges in Rust that encourages you to
write
your code in a way that takes advantage of the safety and easy concurrency that Rust offers. However, you can still make
your variables mutable. Let’s explore how and why Rust encourages you to favor immutability, and why you
might want to opt out.

```rust
fn main() {
    let x = 5;
    println!("The value of x is: {}", x);
    x = 6;
    println!("The value of x is: {}", x);
}
```

**_x variable can not re-assign._**

Immutable variable can not be reassigned but ti can be re-declare or called shadowing.

```rust
fn main() {
    let num = "12345";
    println!("The value of x is: {}", num);
    let num = 12345;
    println!("The value of x is: {}", num);
}
```

But mutability can be very useful. Variables are immutable only by default; we can make them mutable by adding mut in
front of the variable name.

```rust
fn main() {
    let mut x = 5;
    println!("The value of x is: {}", x);
    x = 6;
    println!("The value of x is: {}", x);
}
```

## Constants

Here’s an example of a constant declaration where the constant’s name is MAX_POINTS and its value is set to 100,000.
Rust constant naming convention is to use all upper case with underscores between words. Compiler will give warning and
suggest to upper case.
The underscore in the integer value is for eligible purpose such as in English 100,000 is more eligible than 100000
which most computer languages lack.

```rust
    const MAX_POINTS: u32 = 100_000;
```

## Shadowing

we can declare a new variable with the same name as a previous variable, and the new variable shadows the previous
variable. Rustaceans say that the first variable is shadowed by the second, which means that the second variable’s value
is what we’ll see when we use the variable. We can shadow a variable by using the same variable’s name and repeating the
use of the let keyword as follows:

```rust
fn main() {
    let x = 5;
    let x = x + 1;
    let x = x * 2;

    println!("The value of x is: {}", x);
}
```

```
The value of x is: 12
```

The other difference between mut and shadowing is that because we’re effectively creating a new variable when we use the
let keyword again, we can change the type of the value, but reuse the same name. For example, say our program asks a
user to show how many spaces they want between some text by inputting space characters, but we really want to store that
input as a number:

```rust
fn main() {
    let spaces = "   ";
    let spaces = spaces.len();
}
```

However, if we try to use mut for this, as shown here, we’ll get a compile-time error:

```rust
let mut spaces = "   ";
spaces = spaces.len();
```

The error says we’re not allowed to mutate a variable’s type:

```bash
error[E0308]: mismatched types
--> src/main.rs:3:14
|
3 |     spaces = spaces.len();
|              ^^^^^^^^^^^^ expected &str, found usize
|
= note: expected type `&str`
found type `usize`

```

