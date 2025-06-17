# Rust - Error Handling

Rust has a robust error handling system built into the language.
Rust uses the `Result` enum to handle errors.
The `Result` enum has two variants: `Ok` and `Err`.
The `Ok` variant is used to return a value when the operation is successful,
and the `Err` variant is used to return an error when the operation fails.

## Panic (Unrecoverable Errors)

Panic is a mechanism used to handle unrecoverable errors. When a panic occurs, the program will print an error message
and exit. Panics are used to handle errors that are not recoverable, such as out-of-bounds array access or division by
zero.

```rust
fn main() {
    panic!("This is a panic!");
}
```

## Recoverable Errors

Recoverable errors are errors that can be handled by the program. Rust uses the `Result` enum to handle recoverable
errors. The `Result` enum has two variants: `Ok` and `Err`. The `Ok` variant is used to return a value when the
operation is successful, and the `Err` variant is used to return an error when the operation fails.

```rust
fn main() {
    let result: Result<i32, String> = Ok(42);
    match result {
        Ok(value) => println!("Result: {}", value),
        Err(error) => println!("Error: {}", error),
    }
}
```



