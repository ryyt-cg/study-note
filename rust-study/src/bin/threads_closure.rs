use std::thread;
use std::thread::JoinHandle;

fn main() {
    // Create a vector to hold the child-threads
    let mut handles: Vec<JoinHandle<()>> = vec![];

    for i in 1..11 {
        // need to move i into the closure to avoid the error
        let handle: JoinHandle<()> = thread::spawn( move || {
            println!("counter from thread {}", i);
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }
}