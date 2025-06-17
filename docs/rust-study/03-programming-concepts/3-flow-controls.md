## if-test

The parentheses around if test is not required. Will get compile warming if present.
Must have curly brackets to enclose statement(s).

```rust
fn main() {
    let height = 1.67;

    if height > 1.8 {
        println!("You are tall");
    }
}
```

## if-else-test

```rust
fn main() {
    let height = 1.67;

    if height > 1.8 {
        println!("You are tall");
    } else {
        println!("You are not so tall");
    }
}
```

Writing an if-else expression

```rust
fn main() {
    let greet = "Hola";
    let spoken_language = if greet == "Hola" { "Spanish" } else { "Not Spanish" };

    println!("Language is {}", spokenLanguage);
}
```

## Matching

Rust has a match keyword, which you can to match:

* integers
* Booleans
* Enums
* Arrays
* Tuples
* Structs

Match requires "_" for default

Range

```rust
fn main() {
    let num = 75;

    match num {
        0..=50 => println!("from 0-50"),
        51..=100 => println!("from 51 - 100"),
        _ => println!("No matching")
    }
}
```

Specify multiple patterns using the | syntax

```rust
fn main() {
    let num = 75;

    match num {
        25 | 50 => println!("25, or 50"),
        75 | 100 => println!("75 or100"),
        _ => println!("No matching")
    }
}
```

infinity loop

```rust
fn main() {
    loop {}
}
```

## Using break and continue

```rust
fn main() {
    let some_condition = true;
    loop {
        if some_condition {
            break;
        }
    }
}
```

Continue

```rust
fn main() {
    let some_condition = true;
    loop {
        if some_condition {
            continue;
        }
    }
}
```

## Controlling Nested Loops

Label the out loop For example

outer: loop

break out of the outer loop

break out
