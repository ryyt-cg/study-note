use std::thread;
use std::time::Duration;

/// This function spawns a new thread and prints a counter from both the main thread and the new thread.
/// The new thread will not finish before the main thread exits.
fn main() {
    thread::spawn(|| {
        for i in 1..11 {
            println!("counter from thread {}", i);
            thread::sleep(Duration::from_millis(2_000));
        }
    });

    for i in 1..11 {
        println!("counter from main {}", i);
        thread::sleep(Duration::from_millis(1_000));
    }
}