fn main() {
    let rect = Rectangle::new(1.2, 3.4);
    println!("rect area is {}", rect.get_area());
    println!("rect width is {}", rect.width);
}

struct Rectangle {
    width: f64,
    height: f64
}

impl Rectangle {
    fn new(width: f64, height: f64) -> Rectangle {
        Rectangle {
            width,
            height
        }
    }

    fn get_area(&self) -> f64 {
        self.width * self.height
    }
}