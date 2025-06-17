use std::sync::mpsc;
use std::thread;
use std::time::Duration;

pub fn do_it() {
    println!("\nIn demo_channels_multiple_messages::do_it()");

    // Create a channel to which we can send messages, and from which we can receive those messages.
    // mpsc stands for multiple producer, single consumer.
    let (tx, rx) = mpsc::channel();

    // Spawn a thread to send some messages to the channel.
    let handle = thread::spawn(move || {
        for i in 1..=10 {
            let s = std::format!("Message {}", i);
            tx.send(s).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    // On the main thread, receive the messages from the channel.
    // recv() is a blocking call.
    for received in rx {
        println!("Received: {}", received);
    }

    handle.join().unwrap();
}

// Copilot generates these unit tests.
#[cfg(test)]
mod tests {
    use super::*;
    use std::sync::mpsc;

    #[test]
    fn sends_correct_messages() {
        let (tx, rx) = mpsc::channel();
        let handle = thread::spawn(move || {
            send_messages(tx);
        });

        for i in 1..=10 {
            assert_eq!(rx.recv().unwrap(), format!("Message {}", i));
        }

        handle.join().unwrap();
    }

    #[test]
    fn sends_no_more_than_expected_messages() {
        let (tx, rx) = mpsc::channel();
        let handle = thread::spawn(move || {
            send_messages(tx);
        });

        for _ in 1..=10 {
            rx.recv().unwrap();
        }

        assert!(rx.recv().is_err());

        handle.join().unwrap();
    }
}