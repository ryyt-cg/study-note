fn main() {
    trait Animal {
        // Default implementation for the speak method
        fn speak(&self) -> String {
            "I don't know what to say".to_string()
        }
    }

    struct Dog {
        name: String,
    }
    struct Cat {
        name: String,
    }

    struct Jellyfish {
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

    // Jellyfish does not implement the speak method
    impl Animal for Jellyfish{}

    let dog = Dog { name: "Rusty".to_string() };
    let cat = Cat { name: "Misty".to_string() };
    let jellyfish =  Jellyfish{ name: "Witty".to_string() };

    println!("{} says {}", dog.name, dog.speak());
    println!("{} says {}", cat.name, cat.speak());
    println!("{} says {}", jellyfish.name, jellyfish.speak());

}