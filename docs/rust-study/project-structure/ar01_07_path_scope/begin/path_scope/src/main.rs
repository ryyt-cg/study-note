fn main() {
    greeting::formal::english();
}

mod greeting {
    pub mod formal {
        pub fn english() {
            println!("hello");
        }

        pub fn spanish() {
            println!("hola");
        }
    }

    pub mod casual {
        pub fn english() {
            println!("hey");
        }

        pub fn spanish() {
            println!("oye");
        }
    }
}