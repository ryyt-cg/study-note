# Nested Functions and Closures

## Nested Functions

* You must specify full type info for a nested function, parameter type / return type
* A nested function cannot access variables in the outer scope - you must pass any variables you need as parameters into
  the function

* Closures overcome these constraints

```rust
pub fn do_it() {
    println!("\nIn demo_nested_functions::do_it()");

    fn sqr(i: i32) -> i32 {
        i * i
    }

    println!("Square of 5 is {}", sqr(5));
    println!("Square of 7 is {}", sqr(7));
}
```

## Closures

* Closures in Rust are the equivalent of lambdas in other programming languages
* A closure is similar to a nested function but more flexible in several respects:
    * a closure can infer parameter/return types
    * a closure can capture external variable

* The syntax for defining a closure is similar (but simpler) to defining a nested function:
    * define parameters inside ||
    * optionally specify parameter/return type
    * define the closure body inside {}

* You can assign a closure to a variable
    * then use the variable as if it were a function name, to invoke the closure

```rust
use chrono::{DateTime, Utc};
use std::time::Duration;
use std::thread::sleep;

pub fn do_it() {
    println!("\nIn demo_closures::do_it()");

    closure_no_params();
    closure_one_param();
    closure_many_params();
    closure_multiple_statements();
}

fn closure_no_params() {
    let get_timestamp = || -> DateTime<Utc> { Utc::now() };
    println!("Timestamp:   {}", get_timestamp());//.format("%T"));
}

fn closure_one_param() {
    let reciprocal = |n: f64| -> f64 { if n == 0.0 { 0.0 } else { 1.0 / n } };
    println!("Reciprocal:  {}", reciprocal(5.0));
}

fn closure_many_params() {
    let prod = |a: i32, b: i32| -> i32 { a * b };
    println!("Product:     {}", prod(20, 5));
}

fn closure_multiple_statements() {
    let get_timestamp_after_delay = |seconds: u64| -> DateTime<Utc> {
        sleep(Duration::new(seconds, 0));
        Utc::now()
    };
    println!("Timestamp:   {}", get_timestamp_after_delay(5).format("%T"));
}
```

## Using type inference with closures

```rust
use chrono::Utc;
use std::time::Duration;
use std::thread::sleep;

pub fn do_it() {
    println!("\nIn demo_closures_inferred_types::do_it()");

    closure_no_params();
    closure_one_param();
    closure_many_params();
    closure_multiple_statements();
}

fn closure_no_params() {
    let get_timestamp = || Utc::now();
    println!("Timestamp:   {}", get_timestamp().format("%T"));
}

fn closure_one_param() {
    let reciprocal = |n| if n == 0.0 { 0.0 } else { 1.0 / n };

    // This statement binds the reciprocal() parameter to floating-point.
    println!("Reciprocal:  {}", reciprocal(5.0));

    // So we can't call reciprocal() with any other type thereafter.
    // println!("Reciprocal:  {}", reciprocal(5));
}

fn closure_many_params() {
    let prod = |a, b| a * b;
    println!("Product:     {}", prod(20, 5));
}

fn closure_multiple_statements() {
    let get_timestamp_after_delay = |seconds| {
        sleep(Duration::new(seconds, 0));
        Utc::now()
    };
    println!("Timestamp:   {}", get_timestamp_after_delay(5).format("%T"));
}
```

## Capturing variables by reference

```rust
pub fn do_it() {
    println!("\nIn demo_closures_capture_reference::do_it()");

    capture_immutable_reference();
    capture_mutable_reference();
}

fn capture_immutable_reference() {
    let b1 = String::from("┌─────────────────┐");
    let b2 = String::from("└─────────────────┘");

    let display_heading = |s| {
        println!("{}", b1);
        println!("│ {:<15} │", s);
        println!("{}", b2);
    };

    display_heading(String::from("hello"));
    display_heading(String::from("goodbye!"));

    println!("b1: {}\nb2: {}\n", b1, b2);
}

fn capture_mutable_reference() {
    let mut b1 = String::from("┌─────────────────┐");
    let mut b2 = String::from("└─────────────────┘");

    let mut display_heading = |s| {
        b1.push_str("✅");
        b2.push_str("✅");
        println!("{}", b1);
        println!("│ {:<15} │", s);
        println!("{}\n", b2);
    };

    display_heading(String::from("hello"));
    display_heading(String::from("goodbye!"));

    println!("b1: {}\nb2: {}\n", b1, b2);
}
```

## Capturing variables by value

* You can force a closure to capture variables by value, prefix the closure with the move keyword
* It is useful when you spawn another thread
    * The thread executes the code in a closure
    * The thread might outlive the outer function
    * So the closure must capture variables by value

```rust
use std::time::Duration;

pub fn do_it() {
    println!("\nIn demo_closures_capture_value::do_it()");

    capture_value_automatically();
    capture_value_forcibly();

    // Let's wait a bit before we quit... 
    std::thread::sleep(Duration::new(10, 0))
}

fn capture_value_automatically() {
    let message = String::from("hello");

    println!("Message initially:  {}", message);

    // consume_message moves 'message' into closure here.
    let consume_message = || {
        println!("Message in closure: {}", message);
        std::mem::drop(message);
    };

    // Can't use 'message' here, it's owned by the closure.
    // println!("{}", message);  // Nope!

    // We can call consume_message() once.
    consume_message();

    // But we can't call consume_message() again.
    // consume_message();        // Nope!
}

fn capture_value_forcibly() {
    let message = String::from("HELLO");

    println!("Start of method...");

    // We must forcibly move captured values into closure (otherwise it won't compile).
    std::thread::spawn(move || {
        println!("Message at start of closure: {}", message);
        std::thread::sleep(Duration::new(5, 0));
        println!("Message at end of closure:   {}", message);
    });

    println!("End of method...");
}
```

## Closure Iterator

* Rust has a standard type named Iterator
    * Enables you to iterate over a collection
    * You can perform an operation on each element
    * You specify the operation as a closure

```rust
pub fn do_it() {
    println!("\nIn demo_closures_iteration::do_it()");

    demo_simple_iteration();
    demo_unused_closure_variable();
    demo_filtering_maping();
    demo_collecting_result();
}

fn demo_simple_iteration() {
    let v = vec!["donald", "huey", "louie", "dewey"];

    println!("All ducks:");
    v.iter()
        .for_each(|e| println!("   {}", e));
}

fn demo_unused_closure_variable() {
    let v = vec!["donald", "huey", "louie", "dewey"];

    println!("Redacted ducks:");
    v.iter()
        .for_each(|_| println!("   xxx"));
}

fn demo_filtering_maping() {
    let v = vec!["donald", "huey", "louie", "dewey"];

    println!("Uppercase 'd' ducks:");
    v.iter()
        .filter(|e| e.starts_with('d'))
        .map(|e| e.to_uppercase())
        .for_each(|e| println!("   {}", e));
}

fn demo_collecting_result() {
    let v = vec!["donald", "huey", "louie", "dewey"];

    let upper_y_ducks =
        v.iter()
            .filter(|e| e.ends_with('y'))
            .map(|e| e.to_uppercase())
            .collect::<Vec<String>>();

    println!("There are {} ducks ending with 'y':", upper_y_ducks.len());
    upper_y_ducks.iter()
        .for_each(|e| println!("   {}", e));
}
```
