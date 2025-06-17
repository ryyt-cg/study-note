

/// The async keyword is used to define an asynchronous function.
async fn hello_world() -> String {
    "Hello, World!".to_string()
}

/// The #[tokio::main] attribute is used to define the entry point of a Tokio application.
#[tokio::main]
async fn main() {
    // The await keyword is used to wait for the result of an asynchronous function.
    let result = hello_world().await;
    println!("{}", result);
}