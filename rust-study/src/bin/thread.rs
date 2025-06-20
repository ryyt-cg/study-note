use std::thread;

/// This function spawns a new thread and prints a message from the main thread and the new thread.
/// However, the new thread will not finish before the main thread exits.
fn main() {
    // Spawn a new thread
    thread::spawn(|| {
        // Print a message from the new thread
        println!("hello from thread");
    });

    println!("hello from main")
}