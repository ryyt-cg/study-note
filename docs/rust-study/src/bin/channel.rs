use std::sync::mpsc::channel;
use std::thread::{JoinHandle, spawn};

fn main() {
    // Create a channel
    let (tx, rx) = channel();

    // Create a vector to hold the child-threads
    let mut handles: Vec<std::thread::JoinHandle<()>> = vec![];

    // Create 10 threads and send the counter-value to the channel
    for i in 1..11 {
        // need to move i into the closure to avoid the error
        let tx_clone = tx.clone();
        let handle: JoinHandle<()> = spawn( move || {
            tx_clone.send(i).unwrap();
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    drop(tx);

    // Receive the counter-values from the channel
    // Receiver will block until all the senders are done
    for received in rx {
        println!("Got: {}", received);
    }
}