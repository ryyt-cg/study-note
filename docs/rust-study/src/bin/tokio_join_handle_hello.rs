async fn hello_world(i: i32) -> String {
    std::thread::sleep(std::time::Duration::from_secs(1));
    format!("Hello, World! {i}").to_string()
}

/// The #[tokio::main] attribute is used to define the entry point of a Tokio application.
#[tokio::main]
async fn main() {
    let now = std::time::Instant::now();
    let mut handles = Vec::new();

    for i in 0..10 {
        let handle = tokio::spawn(hello_world(i));
        handles.push(handle);
    }

    // The output will be:
    // Hello, World! 0
    // Hello, World! 1
    // Hello, World! 2
    // Hello, World! 3
    // Hello, World! 4
    // Hello, World! 5
    // Hello, World! 6
    // Hello, World! 7
    // Hello, World! 8
    // Hello, World! 9
    for handle in handles {
        // is blocked by await; therefore, the next handle will not be executed until the current handle is finished.
        let value = handle.await.unwrap();
        println!("{}", value);
    }

    println!("Elapsed: {:.2?}", now.elapsed());
}