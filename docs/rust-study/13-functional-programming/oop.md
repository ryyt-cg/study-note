# Object Oriented Programming in Rust

## Create a Rectangle Class

```rust
struct Rectangle {
    width: u32,
    height: u32,
}
```

## Create Rectangle Methods/Functions

```rust
impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
    
    fn perimeter(&self) -> u32 {
        2 * (self.width + self.height)
    }
}
```

## Create a Rectangle Instance

```rust
fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    
    println!("The area of the rectangle is {} square pixels.", rect1.area());
    println!("The perimeter of the rectangle is {} pixels.", rect1.perimeter());
}
```

## Create Polymorphism in Rust

```rust
trait Shape {
    fn area(&self) -> u32;
    fn perimeter(&self) -> u32;
}

impl Shape for Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
    
    fn perimeter(&self) -> u32 {
        2 * (self.width + self.height)
    }
}

impl Shape for Circle {
    fn area(&self) -> u32 {
        (self.radius * self.radius) * 3
    }
    
    fn perimeter(&self) -> u32 {
        2 * self.radius * 3
    }
}

fn print_shape(shape: &dyn Shape) {
    println!("The area of the shape is {} square pixels.", shape.area());
    println!("The perimeter of the shape is {} pixels.", shape.perimeter());
}
```