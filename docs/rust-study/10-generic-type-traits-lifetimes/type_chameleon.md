# [A Type Chameleon](https://www.youtube.com/watch?v=CWiz_RtA1Hw)

In Rust, we can use a trait object to create a type that can hold any value that implements a specific trait. This is
useful when we want to write functions that can accept any type that implements a trait.

```rust
fn(impl Trait)
fn() -> impl Trait
type = impl Trait
trait{fn() -> impl Trait
```

## fn() -> impl Trait

```rust
fn only_true<I>(iter: I) -> impl Iterator<Item = bool>
where
    I: Iterator<Item = bool>,
{
    iter.filter(|&x| x)
}
```

## Struct, Traits, and Impl

```rust
struct DougsData {
    some_boot: bool,
    some_float: f32,
    some_int: i32,
}

```