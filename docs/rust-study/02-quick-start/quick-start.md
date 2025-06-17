# Generating a New Project

cargo new hello-rust

new directory called hello-rust with the following files:

```
hello-rust
|- Cargo.toml
|- src
  |- main.rs
```

## A small Rust application

Now let’s write a small application with our new dependency. In our main.rs, add the following code:

```rust
use ferris_says::say;
// from the previous step
use std::io::{stdout, BufWriter};

fn main() {
    let stdout = stdout();
    let message = String::from("Hello fellow Rustaceans!");
    let width = message.chars().count();

    let mut writer = BufWriter::new(stdout.lock());
    say(message.as_bytes(), width, &mut writer).unwrap();
}
```

cargo run

```
----------------------------
< Hello fellow Rustaceans! >
----------------------------
              \
               \
                 _~^~^~_
             \) /  o o  \ (/
               '_   -   _'
               / '-----' \
    
```

[Rust Language link!](https://www.rust-lang.org)

[Style Guidelines](https://doc.rust-lang.org/1.0.0/style/README.html)

## Run example and bin executable rust

Run example applications

1. create examples directory at same level of src directory
2. create any file name contains main function
3. cargo run --example file_name

```bash
cargo run --example file_name
```

Run bin applications

1. create bin directory within src directory
2. create any file name contains main function
3. cargo run --bin file_name

```bash
cargo run --bin file_name
```
