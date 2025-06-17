# [Store a trait in the struct](https://users.rust-lang.org/t/how-to-store-a-trait-as-field-of-a-struct/87762)

## Purpose

Demonstrate how to store a trait in a struct.

https://users.rust-lang.org/t/how-to-store-a-trait-as-field-of-a-struct/87762/2?u=rhtran

```rust
use std::{io::Read, fs::File};

struct Test<'r> {
    readable: &'r mut dyn Read,
}

impl<'r> Test<'r> {
    pub fn new(readable: &'r mut dyn Read) -> Self {
        Self { readable }
    }

    pub fn read(&mut self, buf: &mut [u8]) {
        self.readable.read(buf); // <-- ERROR: cannot borrow data in an `Arc` as mutable
    }
}

fn main() {
    let mut file = File::open("foo.txt").unwrap();
    let test = Test::new(&mut file);
}
```