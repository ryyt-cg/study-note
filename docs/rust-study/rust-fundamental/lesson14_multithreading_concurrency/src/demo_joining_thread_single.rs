use std::thread;
use std::time::Duration;

pub fn do_it() {
    println!("\nIn demo_joining_thread_single::do_it()");

    // Spawns a thread with a handle for joining.
    let handle = thread::spawn(|| {
        println!("{:?} starting", thread::current().id());

        thread::sleep(Duration::from_secs(10));
        println!("{:?} ending", thread::current().id());

        // Uncomment the following code, to deliberately panic.
        // panic!("Deliberate panicking, dude!");
    });

    for i in 100..=105 {
        println!("{:?} displaying {}", thread::current().id(), i);
        thread::sleep(Duration::from_millis(500));
    }

    println!(
        "{:?} waiting for other thread to end",
        thread::current().id()
    );

    // Uncomment the following code, to unwrap the join() Result. This could panic!
    // join() method is a blocking call.
    // It waits for the spawn thread to finish.
    handle.join().unwrap();

    // should not use unwrap() in production code.
    // unwrap() will panic if the Result is an Err.
    // it will prematurely end the program.

    // Or uncomment the following code, to test the join() Result explicitly.
    // use match to handle the Result explicitly.
    // match handle.join() {
    //     Ok(_)  => println!("join() result is Ok"),
    //     Err(_) => println!("join() result is Err")
    // }

    println!("That's all, folks!");
}
