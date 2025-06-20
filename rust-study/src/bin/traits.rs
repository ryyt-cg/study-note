trait Animal {
    fn speak(&self) -> String;
}

struct Dog {
    name: String,
}
struct Cat {
    name: String,
}

impl Animal for Dog {
    fn speak(&self) -> String {
        return "Woof!".to_string()
    }
}

impl Animal for Cat {
    fn speak(&self) -> String {
        "Meow!".to_string()
    }
}

fn main() {
    let dog = Dog{name: "Rusty".to_string()};
    let cat = Cat {name: "Misty".to_string()};

    println!("{} says {}", dog.name, dog.speak());
    println!("{} says {}", cat.name, cat.speak());
}