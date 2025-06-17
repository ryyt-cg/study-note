# Generics

- Generics are a way to define a type or function without specifying the exact type it will work with.
- Reusable code can be written with generics.
- Reduces code duplication.

## Defining a Generic Function

```rust
fn print_type<T>(value: T) {
    println!("{}", std::any::type_name::<T>());
}
```

## Defining a Generic Struct

```rust
struct Point<T> {
    x: T,
    y: T,
}
```

## Implementing a Generic Method

```rust
impl<T> Point<T> {
    fn x(&self) -> &T {
        &self.x
    }
}
```
