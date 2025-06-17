use std::sync::mpsc;
use std::thread;

pub fn do_it() {
    println!("\nIn demo_channels_single_message::do_it()");

    // Create a channel to which we can send messages, and from which we can receive those messages.
    let (tx, rx) = mpsc::channel();

    // Spawn a thread to send a message to the channel.
    // Must include move keyword to transfer ownership of tx to the spawned thread.
    // The `move` to make sure the closure has a long enough lifetime of the thread.
    // Can `clone()` be used instead of `move`?
    // Yes, but it is not recommended.
    // I would think that `clone()` would be used to create a new reference to the channel.
    let handle = thread::spawn(move || {
        tx.send(String::from("Hei hei")).unwrap();
    });

    // On the main thread, receive the message from the channel.
    let received = rx.recv().unwrap();
    println!("Received: {}", received);

    handle.join().unwrap();
}
