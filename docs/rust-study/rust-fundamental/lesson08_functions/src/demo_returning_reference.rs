pub fn do_it() {
    println!("\nIn demo_returning_reference::do_it()");

    let s = String::from("hello world");

    let r1 = get_first_word(&s); // Receives &str (implicit typing).
    println!("r1: {}", r1);

    let r2: &str = get_first_word(&s); // Receives &str (explicit typing).
    println!("r2: {}", r2);

    let message: &str = get_message(99);
    println!("message: {}", message);
}

// The get_first_word() function takes a string slice (&str)
// as a parameter and returns a string slice.
// It iterates over the characters in the string until
// it finds a space, and then returns a slice of the string up to
// that point. This demonstrates how a function can return a reference to a part of the data
// it was passed, without taking ownership of the data.
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

// The get_message() function takes an integer as a parameter and returns a
// static string slice (&'static str). This function demonstrates how a function can
// return a reference to a string literal, which is stored in the program's binary
// and has a static lifetime.
fn get_message(mark: i32) -> &'static str {
    if mark >= 50 {
        "PASS😃"
    } else {
        "FAIL😢"
    }
}

// The code also includes two commented-out functions,
// bad_func_1() and bad_func_2(), which demonstrate what happens when a function tries
// to return a reference to data that goes out of scope when the function ends.
// These functions would not compile because they would result in a dangling reference.
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
