use std::thread;
use std::time::Duration;

fn main() {
    let handle = thread::spawn(|| {
        for i in 1..11 {
            println!("counter from thread {}", i);
            thread::sleep(Duration::from_millis(2_000));
        }
    });

    // fn join(self) - Waits for the associated thread to finish
    // This will block the main thread until the spawned thread finishes
    handle.join().unwrap();
    println!("Thread has finished.");

    for i in 1..11 {
        println!("counter from main {}", i);
        thread::sleep(Duration::from_millis(1_000));
    }
}