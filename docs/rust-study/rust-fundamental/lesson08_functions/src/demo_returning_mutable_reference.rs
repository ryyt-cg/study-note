pub fn do_it() {
    println!("\nIn demo_returning_mutable_reference::do_it()");

    let mut s = String::from("hello");
    let r = some_func(&mut s); // Receives mutable reference to a String.

    // After some_func() returns, do_it() continues execution.
    // It appends the string " and goodbye" to the string
    // that r points to, and then prints r.
    // This is allowed because r is a mutable reference to s, so do_it()
    // can still modify s through r.
    r.push_str(" and goodbye");
    println!("r: {}", r);
    // println!("s: {}", s);
}

/*
 The some_func() function takes a mutable reference to a String as a parameter
 and returns a mutable reference to a String. It appends the string " world"
 to the string that x points to, and then returns x. This demonstrates how a
 function can modify data it is passed a reference to, and then returns a reference to the same data.
 */
fn some_func(x: &mut String) -> &mut String {
    x.push_str(" world");
    x
}
