use std::{thread, time};

use tokio::time::{sleep, Duration};

fn blocking_call() -> String {
    thread::sleep(time::Duration::from_secs(1));
    println!("Blocking call");
    "Finally done".to_string()
}

async fn async_call(id: i32) {
    sleep(Duration::from_secs(2)).await;
    println!("Async Call: ID {}", id);
}

/// The function will spawn a blocking call in a separate thread and 10 asynchronous calls in separate threads.
/// It will then wait for all the asynchronous calls to finish, and finally print the result of the blocking call.
#[tokio::main]
async fn main() {
    // Spawn a blocking call in a separate thread with blocking_call_handle
    let blocking_code_handle = tokio::task::spawn_blocking(blocking_call);

    // Spawn 10 async calls in separate threads with async_handles
    let mut async_handles = Vec::new();
    for id in 0..10 {
        async_handles.push(tokio::spawn(async_call(id)));
    }

    // Have two options to wait for the async calls to finish or the blocking call to finish
    // In this case; we will wait for the async calls to finish first.
    for handle in async_handles {
        handle.await.unwrap();
    }

    let result = blocking_code_handle.await.unwrap();
    println!("Blocking call: {}", result);
}
