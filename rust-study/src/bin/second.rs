fn main() {
    println!("Hello, world!");

    let mut a = String::from("test");
    let b = &mut a;
    // b = &mut String::from("lls");

    let c = &a;
}
