# Borrowing

* Simple borrowing
* Rust borrow checker
* String slice techniques
* Array slices
* Array slices techniques

## Simple borrowing

Instead of transferring/moving ownership by assign one variable to another. ie let a = b assumption that a does not have
Copy trait. Rust has feature called borrow.

* Facilitated via references
* A reference can borrow a value, without claiming ownership
* To borrow a value by prefix the value with &
* A value can be immutably borrowed any number of times
* Borrowed values are immutable
* Even if the borrowed value is mutable, that borrow can't be used to mutate the value

```rust
fn main() {
    let s = String::from("huey");
    let r = &s; // r borrows value from s
    // another way to borrow, multiple immutable borrowers is OK.
    let q: &String = &s;

    println!("{}, {}, {}", s, r, q);
}
```

## Mutable borrower

* Mutable borrows allow values to be changed without transferring ownership
* Only one mutable borrow can exist at once

```rust
fn main() {
    let mut s = String::from("dewey");
    let b = &mut a;

    // Not allow multiple mutable borrowers because of race conditions. 
    //let c = &mut a;
}
```

## Borrow Checker / Rules

The borrow checker enforces these rules:

* Single writer (mutable borrow)
* Many readers (immutable borrow)
* No mixing of mutable and immutable borrows to prevent data races

```rust
pub fn do_it() {
    println!("\nIn demo_borrow_checker::do_it()");

    defining_many_immutable_references();
    restrictions_after_defining_mutable_reference();
    restrictions_after_defining_immutable_reference();
}

fn defining_many_immutable_references() {
    let s = String::from("hello");

    // You can define any number of immutable borrows.
    let r1 = &s;
    let r2 = &s;
    let r3 = &s;

    println!("r1: {}, r2: {}, r3: {}", r1, r2, r3);
}

fn restrictions_after_defining_mutable_reference() {
    let mut s = String::from("huey");
    s.push_str(" louie");

    // If you have a mutable borrow...
    let r1 = &mut s;

    // Then you can't have any other borrows at all.
    // let r2 = &mut s;       // Nope!
    // let r3 = &s;           // Nope!

    // And you can't do this either (println! tries to borrow s).
    // println!("s: {}", s);  // Nope!

    r1.push_str(" dewey");

    println!("r1: {}", r1);
}

fn restrictions_after_defining_immutable_reference() {
    let mut s = String::from("huey");
    s.push_str(" louie");

    // If you have immutable borrow(s)...
    let r1 = &s;
    let r2 = &s;

    // Then you can't have any mutable borrows.
    // let r3 = &mut s;       // Nope!

    // And you also can't modify the original object either.
    // s.push_str(" dewey");  // Nope!

    println!("s: {}, r1: {}, r2: {}", s, r1, r2);
}
```

## String Slice Intro

```rust
pub fn do_it() {
    println!("\nIn demo_string_slice_intro::do_it()");

    slice_string_object();
    slice_string_literal();
}

fn slice_string_object() {
    let obj = String::from("hello");

    // Define a &String variable (implicitly typed).
    let s1 = &obj;

    // Define a &str variable (explicitly typed).
    let s2: &str = &obj;

    println!("s1: {}, s2: {}", s1, s2);
}

fn slice_string_literal() {

    // Define a &str variable (implicitly typed).
    let s3 = "hello";

    // Define a &str variable (implicitly typed)
    let s4: &'static str = "world";

    println!("s3: {}, s4: {}", s3, s4);
}
```

## String Slice Techniques

```rust
pub fn do_it() {
    println!("\nIn demo_string_slice_techniques::do_it()");

    slice_usage();
    slice_iteration();
    slice_part_of_string();
    slice_mutability();
}

fn slice_usage() {

    // Declare a string slice.
    let s1 = "hello😃";

    // Note: This examples would also work with String objects.
    // let s1 = String::from("hello😃");

    println!("\ns1 ptr: {:p}, len: {}, text: {}", s1.as_ptr(), s1.len(), s1);
}

fn slice_iteration() {

    // Declare a string slice.
    let s2 = "hello😃";

    // Note: This examples would also work with String objects.
    // let s2 = String::from("hello😃");

    println!("\nRaw bytes in s2 (in decimal, hex, and octal):");
    for b in s2.bytes() {
        println!("  {} {:x} {:o}", b, b, b);
    }

    println!("Characters in s2:");
    for ch in s2.chars() {
        println!("  {}", ch);
    }
}

fn slice_part_of_string() {
    let message = "howdy😎";

    // Create slices as a portion of string.
    let s3 = &message[0..3];
    let s4 = &message[..3];
    let s5 = &message[2..5];
    let s6 = &message[2..];

    println!("\ns3 ptr: {:p}, len: {}, text: {}", s3.as_ptr(), s3.len(), s3);
    println!("s4 ptr: {:p}, len: {}, text: {}", s4.as_ptr(), s4.len(), s4);
    println!("s5 ptr: {:p}, len: {}, text: {}", s5.as_ptr(), s5.len(), s5);
    println!("s6 ptr: {:p}, len: {}, text: {}", s6.as_ptr(), s6.len(), s6);
}

fn slice_mutability() {
    let mut message = String::from("croeso");
    message.push_str(" o gymru");

    if true {
        let s: &mut str = &mut message[9..];
        s.make_ascii_uppercase();
    }

    println!("\nmessage: {}", message);
}
```
