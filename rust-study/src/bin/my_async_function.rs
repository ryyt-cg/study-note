use std::time::{Duration, Instant};

use tokio::time::{sleep};

#[tokio::main]
async fn main() {
    let now = Instant::now();

    // so far this still runs synchronously
    let f = my_function(100);
    println!("Let's Get Rusty");
    f.await;

    // Concurrency
    let mut handles = vec![];

    // This will spawn 5 tasks that will run concurrently
    // The order in which they finish is not guaranteed
    for i in 0..5 {
        let handle = tokio::spawn(async move {
            my_function(i).await;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.await.unwrap();
    }

    println!("Elapsed: {:.2?}", now.elapsed());
}

async fn my_function(i: i32) {
    println!("{i} I'm an async function");
    let s1 = read_from_database().await;
    println!("{i} First result: {s1}");
    let s2 = read_from_database().await;
    println!("{i} Second result: {s2}");
    // let s3 = read_from_database().await;
    // println!("{i} Third result: {s3}");
}

async fn read_from_database() -> String {
    sleep(Duration::from_millis(50)).await;
    "DB Result".to_owned()
}
