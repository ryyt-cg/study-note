use std::sync::Arc;

use tokio::sync::Semaphore;
use tokio::time::{sleep, Duration};

/// This function simulates a person waiting in line to be served by a teller.
///
/// # Arguments
///
/// * `semaphore` - An Arc-wrapped Semaphore that controls access to the tellers.
/// * `name` - The name of the person.
///
/// # Remarks
///
/// The function will print a message when the person starts waiting and when they are being served.
async fn person(semaphore: Arc<Semaphore>, name: String) {
    println!("{} is waiting in line", name);

    teller(semaphore, name).await;
}

/// This function simulates a teller serving a customer.
///
/// # Arguments
///
/// * `semaphore` - An Arc-wrapped Semaphore that controls access to the tellers.
/// * `customer` - The name of the customer being served.
///
/// # Remarks
///
/// The function will print a message when the customer starts being served and when they leave.
async fn teller(semaphore: Arc<Semaphore>, customer: String) {
    let permit = semaphore.acquire().await.unwrap();

    sleep(Duration::from_secs(2)).await;
    println!("\n{} is being served by the teller", customer);
    sleep(Duration::from_secs(5)).await;
    println!("{} is now leaving the teller", customer);

    drop(permit);
}


/// This is the main function that simulates a queue of people waiting to be served by tellers.
///
/// # Remarks
///
/// The function will spawn a number of tasks equal to the number of people, each task representing a person waiting in line.
/// The Semaphore controls access to the tellers, ensuring that only a certain number of people can be served at the same time.
#[tokio::main]
async fn main() {
    let num_of_tellers = 4;
    let semaphore = Semaphore::new(num_of_tellers);
    let semaphore_arc = Arc::new(semaphore);

    let mut people_handles = Vec::new();
    for num in 0..10 {
        people_handles.push(tokio::spawn(person(
            semaphore_arc.clone(),
            format!("Person_{num}"),
        )));
    }

    for handle in people_handles {
        let _ = handle.await.unwrap();
    }
}
