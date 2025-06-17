# Rust Test Functions

In his 1972 essay “The Humble Programmer,” Edsger W. Dijkstra said that “Program testing can be a very effective way to
show the presence of bugs, but it is hopelessly inadequate for showing their absence.”

* Define using the #[test] attribute
* Run with cargo test command

## assert! Macro

* Evaluates a single Boolean argument
* if true, do nothing
* if false, call the panic! macro

assert_eq!
assert_ne!

## Should_panic

* add #[should_panic(expected="message")] if the code expects panic. This functionality is similar to Java @Exception

## Test Filter

```bash
cargo test test_name
```

## Ignoring Tests

```rust
#[ignore]
```

