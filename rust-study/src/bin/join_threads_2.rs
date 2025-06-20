use std::thread;
use std::time::Duration;

fn main() {
    let handle = thread::spawn(|| {
        for i in 1..11 {
            println!("counter from thread {}", i);
            thread::sleep(Duration::from_millis(2_000));
        }
    });

    for i in 1..11 {
        println!("counter from main {}", i);
        thread::sleep(Duration::from_millis(1_000));
    }

    // Place the join call here; after the main thread has started,
    // This will not block the main thread
    // Wait for the thread to finish
    handle.join().unwrap();
    println!("Thread has finished.");
}