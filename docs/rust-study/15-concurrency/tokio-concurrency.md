# Concurrency with Tokio

Asynchronous programming in Tokio is based on the `Future` trait. A `Future` is a computation that may not have finished
yet. Tokio provides a runtime that can execute these futures concurrently. This allows you to write asynchronous code
that is non-blocking and efficient. In this guide, we will explore how to write asynchronous code using Tokio.

```rust
/// The async keyword is used to define an asynchronous function.
async fn hello_world() -> String {
    "Hello, World!".to_string()
}

/// The #[tokio::main] attribute is used to define the entry point of a Tokio application.
#[tokio::main]
async fn main() {
    // the hello_world function returns a future that will be executed by the Tokio runtime
    let result = hello_world().await;
    println!("{}", result);
}
```

tokio::main is a macro that sets up the Tokio runtime and runs the async main function. The async keyword is used to
define an asynchronous function. Its functionality looks like this:

```rust
fn main() {
    let result = tokio::runtime::Runtime::Builder::new_current_thread()
        .enable_all()
        .build()
        .unwrap()
        .block_on(async { println!("Hello World"); });
}
```

## Testing Asynchronous Code

Testing asynchronous code can be challenging. Tokio provides a testing framework that makes it easy to test asynchronous
code. The `tokio::test` attribute is used to define an asynchronous test function. The `tokio::test` attribute sets up
the Tokio runtime and runs the test function asynchronously. Here is an example of how to test an asynchronous function:

```rust
#[tokio::test]
async fn test_hello_world() {
    let result = hello_world().await;
    assert_eq!(result, "Hello, World!");
}
```

## Channels

Channels are a way to communicate between threads. Tokio provides a channel implementation that allows you to send
messages between asynchronous tasks. The `tokio::sync::mpsc` module provides a channel implementation that is similar to
the standard library's `std::sync::mpsc` module. Here is an example of how to use a channel:

### oneshot Channel

The `tokio::sync::oneshot` module provides a channel implementation that allows you to send a single message between
asynchronous tasks. The `tokio::sync::oneshot::channel` function creates a oneshot channel that returns a `Sender` and
a `Receiver`. Here is an example of how to use a oneshot channel:

```rust
use tokio::sync::oneshot;
 
fn main() {
    let (tx, rx) = oneshot::channel();
 
    tokio::spawn(async move {
        let _ = tx.send("Hello, World!");
    });
 
    tokio::spawn(async move {
        let result = rx.await.unwrap();
        println!("{}", result);
    });
}
```

### mpsc Channel

The `tokio::sync::mpsc` module provides a channel implementation that allows you to send messages between asynchronous
tasks. The `tokio::sync::mpsc::channel` function creates a mpsc channel that returns a `Sender` and a `Receiver`. Here
is an example of how to use a mpsc channel:

```rust
use tokio::sync::mpsc;

fn main() {
    let (tx, mut rx) = mpsc::channel(32);
 
    tokio::spawn(async move {
        tx.send("Hello, World!").await.unwrap();
    });
 
    tokio::spawn(async move {
        while let Some(message) = rx.recv().await {
            println!("{}", message);
        }
    });
}
```