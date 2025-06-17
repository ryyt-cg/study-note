# Naming Conventions

In general, Rust tends to use CamelCase for "type-level" constructs (types and traits) and snake_case for "value-level"
constructs. More precisely:

| Item	                   | Convention                                            |
|-------------------------|-------------------------------------------------------|
| Crates                  | snake_case (but prefer single word)                   |
| Modules                 | snake_case                                            |
| Types                   | CamelCase                                             |
| Traits                  | CamelCase                                             |
| Enum variants           | CamelCase                                             |
| Functions               | snake_case                                            |
| Methods                 | snake_case                                            |
| General constructors    | new or with_more_details                              |
| Conversion constructors | from_some_other_type                                  |
| Local variables         | snake_case                                            |
| Static variables	       | SCREAMING_SNAKE_CASE                                  |
| Constant variables      | SCREAMING_SNAKE_CASE                                  |
| Type parameters         | concise CamelCase, usually single uppercase letter: T |
| Lifetimes               | short, lowercase: 'a                                  |

## Passing Parameters by Value

* When passing a copyable value (Copy trait)
    * Rust bit-copies the value into the parameter
    * The original function retains ownership of the value
* When passing a non-copyable value (e.g. String)
    * Rust moves ownership of the value into the parameter
    * The original function loses ownership of the value!

```rust
pub fn do_it() {
    println!("\nIn demo_passing_values::do_it()");

    let n = 42;
    let s = String::from("hello");

    some_func(n, s);            // Copies n, but moves ownership of s.

    println!("n: {}", n);       // OK. We still own n.
    // println!("s: {}", s);    // NO! We don't own s.
}

fn some_func(iparam: i32, sparam: String) {
    println!("In some_func, iparam is {}", iparam);
    println!("In some_func, sparam is {}", sparam);
}
```

## Module

Define a module using the mod keyword

```rust
mod foo {
    pub struct FooError {}
}
```

## Functions

## Enum

Enum is essential in Rust. It is used to define a type by enumerating its possible values.

## Defining and using simple enums

* Specify the enum type, starting with a capital
* Specify allowed values (variants), also starting with capitals

## Using an Enum Type

* Use the enum type, followed by ::, followed by a variant
* you often use match to test variant values

```rust

enum Color {
    Red,
    Green,
    Blue
}

fn main() {
    let c: Color = Color::Red;

    match c {
        Color::Blue => println!("Blue"),
        Color::Green => println!("Green"),
        Color::Red => println!("Red"),
    }
}

```

## Avoid Dead Code

Use some variants but no all. Will get warnings.

* To allow dead code, you can use the following attribute:
    * #[allow(dead_code)]

## Understanding Enum

Associate data with variants in enum type

```rust
enum HouseLocation {
    Number(i32),
    Name(String),
    Unknown
}

fn main() {
    let h1 = HouseLocation::Number(32);
    let h2 = HouseLocation::Name("Cartref");
}

```

## Using the Option enum

```rust
enum Option<T> {
    Some(T),
    None
}
```

Option is typically used as a function to return type

* Indicates the return value might be empty not dealing with null value\

```rust
fn demo_using_option_enum() {
    println!("\nDemo using the Option<T> enum");

    let sec: Option<u32>;

    // Uncomment one of the following statements.
    sec = sec_of_day(23, 59, 59);
    // sec = sec_of_day(1234, 59, 59);

    match sec {
        Some(s) => println!("Second of day: {}", s),
        None => println!("Second of day: no value available")
    }

    // You can use unwrap_or() to extract Some value from an Option, or use a fallback value if None.
    println!("Unwrapped sec: {}", sec.unwrap_or(0));
}

fn sec_of_day(h: u32, m: u32, s: u32) -> Option<u32> {
    return if h <= 23 && m <= 59 && s <= 59 {
        let secs = h * 3600 + m * 60 + s;
        Option::Some(secs)
    } else {
        Option::None
    };
}

```

## Using the Result enum

* Rust defines a standard enum type name Result:
* Result represents a value or an error
    * Might contain a value of some type T
    * Or might contain an error of some type E
* This is how Rust indicates errors rather than via exceptions

```rust
enum Result<T, E> {
    OK(T),
    Err(E)
}

```

```rust
fn demo_using_result_enum() {
    println!("\nDemo using the Result<T, E> enum");

    let res: Result<i32, std::num::ParseIntError>;

    // Uncomment one of the following statements.
    res = i32::from_str_radix("FF", 16);
    // res = i32::from_str_radix("wibble", 16);

    match res {
        Ok(n) => println!("Parsed str as i32: {}", n),
        Err(e) => println!("Error occurred: {}", e)
    }

    // You can use unwrap_or() to extract Ok value from a Result, or use a fallback value if Err.
    let res2 = i32::from_str_radix("FF", 16);
    println!("Unwrapped result: {}", res2.unwrap_or(-1));
}
```

## Passing reference parameter

* The called function receives a reference
    * The called function borrows the value
* The calling function retains ownership
    * The calling function can continue to use the value afterward
* To pass a reference parameter into a function
    * Precede the parameter with &
* When a function receives a reference parameter
    * Use * to de-reference to obtain the underlying value
    * Rust allow a more direct syntax fo invoking methods, no need to de-reference it.

* Declare a &str parameter
    * You can pass in a #String
    * You can pass in a &str. i.e. some_func("Jo"). &String can not receive string literals ("Jp")

```rust
pub fn do_it() {
    println!("\nIn demo_passing_references::do_it()");

    let n = 42;
    let s = String::from("Ola Nordmann");

    some_func1(&n, &s);       // Borrows n and s, i.e. passes references.
    // some_func2(&n, "Jo");  // NO! Can't pass &str into &String.

    some_func2(&n, &s);
    some_func2(&n, "Siv Nordmann");

    some_func3(&n, &s);
    some_func3(&n, "Per Nordmann");

    println!("n: {}", n);     // OK. We still own s.
    println!("s: {}", s);     // OK. We still own n.
}

fn some_func1(iparam: &i32, sparam: &String) {
    if *iparam >= 50 {
        println!("{}, {}, PASS 😃", *iparam, (*sparam).to_uppercase());
    } else {
        println!("{}, {}, FAIL 😢", *iparam, (*sparam).to_lowercase());
    }

    // Note Rust allows you to simplify method calls. Both the following are OK:
    //   (*sparam).to_uppercase()
    //   sparam.to_uppercase()
}

fn some_func2(iparam: &i32, sparam: &str) {
    if *iparam >= 50 {
        println!("{}, {}, PASS 😃", *iparam, sparam.to_uppercase());
    } else {
        println!("{}, {}, FAIL 😢", *iparam, sparam.to_lowercase());
    }
}

fn some_func3(iparam: &i32, sparam: &str) {
    println!("Values {0} and {1}, addresses {0:p} and {1:p}", iparam, sparam);
}
```

## Passing mutable reference parameters

* To pass a mutable reference parameter into a function, precede the parameter with &mut

```rust
pub fn do_it() {
    println!("\nIn demo_returning_mutable_reference::do_it()");

    let mut s = String::from("hello");
    let r = some_func(&mut s);    // Receives mutable reference to a String.

    r.push_str(" and goodbye");
    println!("r: {}", r);
}

fn some_func(s: &mut String) -> &mut String {
    s.push_str(" world");
    s
}
```

## Returning a value

syntax: fn some_func(params...) -> some_type {}

```rust
pub fn do_it() {
    println!("\nIn demo_returning_value::do_it()");

    let n = func_returning_copyable_type();         // Receives copy of value.
    println!("n: {}", n);

    let s = func_returning_non_copyable_type();     // Receives ownership of value.
    println!("s: {}", s);
}

fn func_returning_copyable_type() -> i32 {
    let n = 42;
    return n;       // Copies value back to caller.
}

fn func_returning_non_copyable_type() -> String {
    let s = String::from("hello");
    return s;       // Moves ownership back to caller.
}

// This function doesn't compile. 
// It should return a String object.
// Instead it returns a string literal (i.e. &str).
// fix it change to String::from("PASS")
/*
fn bad_func_returning_string_literal(mark: i32) -> String {
 	return if mark >= 50 {"PASS"} else {"FAIL"};
}
*/
```

## Return a reference

### Lifetime Management

Lifetime annotations don't change how long references live. They just describe the relationships of the lifetimes of
multiple references to each other without affecting the lifetimes.

* Rust doesn't allow to return a dangling reference
    * You can't return a reference to a local-based object
    *

```rust
pub fn do_it() {
    println!("\nIn demo_returning_reference::do_it()");

    let s = String::from("hello world");

    let r1 = get_first_word(&s);        // Receives &str (implicit typing).
    println!("r1: {}", r1);

    let r2: &str = get_first_word(&s);  // Receives &str (explicit typing).
    println!("r2: {}", r2);

    let message: &str = get_message(99);
    println!("message: {}", message);
}

fn get_first_word(s: &str) -> &str {
    let mut pos = 0;
    for ch in s.chars() {
        if ch == ' ' {
            break;
        }
        pos += 1
    }
    &s[..pos]
}

fn get_message(mark: i32) -> &'static str {
    if mark >= 50 { "PASS😃" } else { "FAIL😢" }
}

// This won't compile, because it returns a dangling reference.
/*
fn bad_func_1() -> &str {
   let s = String::from("hello");
   &s                   
}
*/

// This won't compile either, because it also returns a dangling reference.
/*
fn bad_func_2(s: String) -> &str {
	let mut pos = 0;
	for ch in s.chars() {
		if ch == ' ' {
			break;
		} 
		pos += 1
	}           
	&s[..pos]
}
*/
```

## Returning a mutable reference

```rust
pub fn do_it() {
    println!("\nIn demo_returning_mutable_reference::do_it()");

    let mut s = String::from("hello");
    let r = some_func(&mut s);    // Receives mutable reference to a String.

    r.push_str(" and goodbye");
    println!("r: {}", r);
}

fn some_func(s: &mut String) -> &mut String {
    s.push_str(" world");
    s
}
```

## Comments

All programmers strive to make their code easy to understand, but sometimes extra explanation is warranted. In these
cases, programmers leave notes, or comments, in their source code that the compiler will ignore but people reading the
source code may find useful.

Simple comment:

```asciidoc
// Hello, world.
```

In Rust, comments must start with two slashes and continue until the end of the line. For comments that extend beyond a
single line, you’ll need to include // on each line, like this:

```asciidoc
// So we’re doing something complicated here, long enough that we need
// multiple lines of comments to do it! Whew! Hopefully, this comment will
// explain what’s going on.
```

### Doc Comments

Doc comments are prefixed by three slashes (///) and indicate documentation that you would like to be included in
Rustdoc's output. They support Markdown syntax and are the main way of documenting your public APIs.

### Sentence structure

All doc comments, including the summary line, should begin with a capital letter and end with a period, question mark,
or exclamation point. Prefer full sentences to fragments.

The summary line should be written in third-person singular present indicative form. Basically, this means write "
Returns" instead of "Return".

```rust
/// Sets up a default runtime configuration, given compiler-supplied arguments.
///
/// This function will block until the entire pool of M:N schedulers has
/// exited. This function also requires a local thread to be available.
///
/// # Arguments
///
/// * `argc` & `argv` - The argument vector. On Unix this information is used
///                     by `os::args`.
/// * `main` - The initial procedure to run inside of the M:N scheduling pool.
///            Once this procedure exits, the scheduling pool will begin to shut
///            down. The entire pool (and this function) will only return once
///            all child threads have finished executing.
///
/// # Return value
///
/// The return value is used as the process return code. 0 on success, 101 on
/// error.
```

