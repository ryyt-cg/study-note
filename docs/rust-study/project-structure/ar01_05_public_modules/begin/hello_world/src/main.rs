fn main() {
    hello::english();
}

mod hello {
    fn english() {
        println!("hello");
    }

    fn spanish() {
        println!("hola");
    }

    mod casual {
        fn english() {
            println!("hey");
        }
    }
}