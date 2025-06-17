use std::sync::Mutex;
fn main() {
    // Create a mutex
    let mutex = Mutex::new(0);
    {
        // Lock the mutex
        let mut data = mutex.lock().unwrap();
        *data = 10;
    }
    // Unlock the mutex
    let data = mutex.lock().unwrap();
    println!("Data: {}", *data);
}