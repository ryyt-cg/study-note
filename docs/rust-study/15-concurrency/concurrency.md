# Rust Concurrency

* Threads
* Message Channels
* Mutexes and Atomic Variables

## Threads

### Common Concurrency Errors

Race Conditions: Threads interact with data in an unpredictable way. This can lead to data corruption and other issues.

Deadlocks: Threads are waiting for each other to release a resource that they need. This can lead to a situation where
no thread can make progress.

### Creating Threads

Use the `thread::spawn` function to create a new thread.

```rust
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
```

### Joining Threads

Use the [`join`](https://doc.rust-lang.org/std/thread/struct.JoinHandle.html#method.join) method to wait for a thread to
finish.

```rust
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
    // join() is a blocking operation, meaning that it will block the current thread until the associated thread has finished.
    // In this case, it will block the main thread until the spawned thread finishes
    handle.join().unwrap();
    println!("Thread has finished.");

    for i in 1..11 {
        println!("counter from main {}", i);
        thread::sleep(Duration::from_millis(1_000));
    }
}
```

### Closure and Threads

When you create a closure, it captures the environment in which it is created. This means that it can access variables
from the surrounding scope.

```rust
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
```

## Message Channels

Channels are a way to send messages between threads. They are a way to communicate between threads without sharing
memory.

### Creating a Channel

Use the `channel` function to create a channel. The function returns a tuple with a `Sender` (tx) and a `Receiver` (rx).

```rust
use std::sync::mpsc;
```

Common Currency Errors:

* Race Conditions - Threads interact with data in an unpredictable way. This can lead to data corruption and other
  issues.
* Deadlocks - Threads are waiting for each other to release a resource that they need. This can lead to a situation
  where no thread can make progress.
* Starvation - A thread is unable to make progress because it is waiting for a resource that is being held by another
  thread.

```rust
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
```

### Unbounded Channels

The `channel` function creates an unbounded channel. This means that the channel can hold an unlimited number of
messages.

```rust
use std::sync::mpsc::channel;

fn main() {
    // Create a bounded channel
    let (tx, rx) = sync_channel();
}
```

### Bounded Channels

The `sync::mpsc::channel` function can be used to create a bounded channel. This means that the channel can hold a
limited number of messages.

```rust
use std::sync::mpsc::{channel, sync_channel};

fn main() {
    // Create a bounded channel
    let (tx, rx) = sync_channel(3);
}
```

### Blocking and Non-Blocking Channels

The `send` and `recv` methods are blocking. This means that they will block the thread until the operation is complete.
The `try_send` and `try_recv` methods are non-blocking. This means that they will return an error if the operation
cannot be completed immediately.

```rust
use std::sync::mpsc::channel;
```

## Mutexes and Atomic Variables

Mutexes are a way to share data between threads. They are a way to protect data.

### Creating a Mutex

Use the `Mutex` struct to create a mutex. The `lock` method is used to acquire the lock. The `unlock` method is used to
release the lock.

```rust
use std::sync::Mutex;
```

### Mutexes and Threads

Mutexes can be used to share data between threads. The `lock` method is used to acquire the lock. The `unlock` method is
used to release the lock.

```rust
use std::sync::Mutex;
use std::thread::{JoinHandle, spawn};
```