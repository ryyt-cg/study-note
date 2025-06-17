        # Error Handling

## Resources

Some links to help you understand error handling in Rust.

| Topic                     | Link                                                 |
|---------------------------|------------------------------------------------------|
| Deep Dive                 | https://www.lpalmieri.com/posts/error-handling-rust/ |
| Error Handle Patterns     | https://www.youtube.com/watch?v=f82wn-1DPas          |
| Get Rusty - Error Handler | https://www.youtube.com/watch?v=wM6o70NAWUI&t=549s   |
|                           |                                                      |

## Error Handle Patterns

### [Option](https://doc.rust-lang.org/std/option/enum.Option.html)

Type Option represents an optional value: every Option is either Some and contains a value, or None, and does not.
Option types are very common in Rust code, as they have a number of uses:

```rust
pub enum Option<T> {
    None,
    Some(T)
}
```

### [Result](https://doc.rust-lang.org/std/result/enum.Result.html)

Result is a type that represents either success (Ok) or failure (Err). It is often used to return errors to the caller.

```rust
pub enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

1. a function that takes a vector of strings
2. parse each string into a integer
3. computes the sum of all those integers
4. converts the result into a string
5. returns the string to the caller

```rust
fn sum_of_strings(strs: Vec<&str>) -> String {
    let mut sum = 0;
    for s in strs {
        sum += to_int(&s);
    }

    sum.to_string()
}

fn to_int(s: &str) -> i32 {
    s.parse::<i32>().unwrap()
}

fn main() {
    let strs = vec!["1", "2", "3"];
    let sum = sum_of_strings(strs);
    println!("sum: {:?}", sum);
}
``` 

* Should not use unwrap in production code. It will panic (end the function prematurely) if the string is not a valid
  integer. Using unwrap is meant to defer error handling to the caller.
* Use expect to provide a more informative error message. However, it still panics if the string is not a valid integer.

```rust
fn to_int(s: &str) -> i32 {
    s.parse::<i32>().expect("Failed to parse string to integer")
}
```

* use unwrap_or to provide a default value if the string is not a valid integer.

```rust
fn to_int(s: &str) -> i32 {
    s.parse::<i32>().unwrap_or(0)
}
```

* use Option to handle the error. It returns None if the string is not a valid integer.

```rust
fn to_int(s: &str) -> Option<i32> {
    s.parse().ok()
}

fn sum_of_strings(strs: Vec<&str>) -> String {
    let mut sum = 0;
    for s in strs {
        sum += match to_int(&s) {
            Some(i) => i,
            None => 0,
        };
    }

    sum.to_string()
}
```

```rust
fn sum_of_strings(strs: Vec<&str>) -> String {
    let mut sum = 0;
    for s in strs {
        if let Some(val) = to_int(&s) {
            sum += val;
        };
    }

    sum.to_string()
}
```

* use question mark operator to propagate the error to the caller.

```rust
fn sum_of_strings(strs: Vec<&str>) -> Option<String> {
    let mut sum = 0;
    for s in strs {
        let val =  to_int(&s)?;
        sum += val;
    }

    Ok(sum.to_string())
}
```

* use Result to handle the error. It returns an error if the string is not a valid integer.

```rust
#[derive(Debug)]
struct SummationError;
fn sum_of_strings(strs: Vec<&str>) -> Result<String, SummationError> {
    let mut sum = 0;
    for s in strs {
        let val =  to_int(&s).ok_or(SummationError)?;
        sum += val;
    }

    Ok(sum.to_string())
}
```

## Summary

* unwrap - extracts a value from a container or panics if the value is None or Err.
* ok - converts a Result to an Option.
* map - transforms the value inside a container. but with a different type.
* ? - propagates the error to the caller.