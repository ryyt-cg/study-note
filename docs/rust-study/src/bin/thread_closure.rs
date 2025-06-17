use std::thread;

fn main() {
    let message = String::from("Hello World");

    thread::spawn( move || {
        println!("{}", message);
    }).join().unwrap();

    // the message is moved to the closure above
    // drop(message);
}