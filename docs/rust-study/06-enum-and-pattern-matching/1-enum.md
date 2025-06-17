# Defining an Enum

We can express this concept in code by defining an IpAddrKind enumeration and listing the possible kinds an IP address
can be, V4 and V6. These are known as the variants of the enum:

```rust
enum IpAddrKind {
    V4,
    V6,
}
```

```rust
use std::net::IpAddr;

enum IpAddr {
    V4(String),
    V6(String),
}

let home = IpAddr::V4(String::from("127.0.0.1"));
let loopback = IpAddr::V6(String::from("::1"));
```

```rust
enum IpAddr {
    V4(u8, u8, u8, u8),
    V6(String),
}

let home = IpAddr::V4(127, 0, 0, 1);
let loopback = IpAddr::V6(String::from("::1"));
```

```rust
enum Option<T> {
    Some(T),
    None,
}
```

# The match Control Flow Operator

```rust
enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

fn value_in_cents(coin: Coin) -> u32 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter => 25,
    }
}
```

## Matching with Option<T>

```rust
fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        None => None,
        Some(i) => Some(i + 1),
    }
}

let five = Some(5);
let six = plus_one(five);
let none = plus_one(None);
```

## The _ Placeholder

Rust also has a pattern we can use in situations when we don’t want to list all possible values. For example, a u8 can
have valid values of 0 through 255. If we only care about the values 1, 3, 5, and 7, we don’t want to have to list out
0, 2, 4, 6, 8, 9 all the way up to 255. Fortunately, we don’t have to: we can use the special pattern _ instead:

```rust
let some_u8_value = 0u8;
match some_u8_value {
    1 => println!("one"),
    3 => println!("three"),
    5 => println!("five"),
    7 => println!("seven"),
    _ => (),
}
```

## Concise Control Flow with if let

The if let syntax lets you combine if and let into a less verbose way to handle values that match one pattern and ignore
the rest. Consider the program in Listing 6-6 that matches on an Option<u8> value but only wants to execute code if the
value is three:

```rust
let some_u8_value = Some(0u8);
match some_u8_value {
    Some(3) => println!("three"),
    _ => (),
}
```

We want to do something with the Some(3) match but do nothing with any other Some<u8> value or the None value. To
satisfy the match expression, we have to add _ => () after processing just one variant, which is a lot of boilerplate
code to add.

Instead, we could write this in a shorter way using if let. The following code behaves the same as the match in Listing
6-6:

```rust
let some_u8_value = Some(0u8);
if let Some(3) = some_u8_value {
    println!("three");
}
```

