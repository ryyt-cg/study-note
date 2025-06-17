use std::thread;
use std::time::Duration;

/// Spawns a thread
/// Join a thread with `handle.join().unwrap();`
/// Move data into the thread with `thread::spawn(move || {`
/// Communicate between threads with channels

pub fn do_it() {
    println!("\nIn demo_spawning_threads::do_it()");

    // Spawns a thread
    // thread::current().id() returns the thread id.
    // useful for debugging.
    thread::spawn(|| {
        for i in 1..=10 {
            println!("{:?} displaying {}", thread::current().id(), i);
            thread::sleep(Duration::from_secs(1));
        }
    });

    for i in 100..=105 {
        println!("{:?} displaying {}", thread::current().id(), i);
        thread::sleep(Duration::from_secs(4));
    }

    println!("That's all, folks!");
}
