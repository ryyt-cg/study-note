pub fn do_it() {
    println!("\nIn demo_passing_mutable_references::do_it()");

    let mut n = 42;
    let mut s = String::from("hello");

    some_func(&mut n, &mut s); // Mutably borrows n and s.

    // After some_func() returns, do_it() continues execution.
    // It increments n by 1_000_000 and appends a string to s.
    // This is allowed because n and s are still owned by do_it(),
    // even though their references were passed to some_func().
    n += 1_000_000; // OK. We still own n.
    s.push_str("👍👍👍"); // OK. We still own s.

    // Finally, it prints the modified values of n and s.
    println!("n: {}", n);
    println!("s: {}", s);
}

// The some_func() function takes two mutable references as parameters, iparam and sparam.
// It first prints the initial values of these parameters.
// Then, it modifies the values that iparam and sparam point to.
// This is done using the dereference operator *, which is necessary
// because iparam and sparam are references to the values, not the values themselves.
fn some_func(iparam: &mut i32, sparam: &mut String) {
    println!("Values initially: {}, {}", iparam, sparam);
    *iparam += 10;
    sparam.push_str(" world");
    println!("Values afterward: {}, {}", iparam, sparam);
}
