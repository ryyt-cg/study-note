Note: as a general recommendation, don't put trait bounds directly on a struct when it is avoidable (because such bounds
are forced to appear all over the code). Instead, put it on the functions or impls which require that trait.

```rust
pub struct Bar<'a, T> { // Note: No mention of SomeTrait here
    abc: &'a str,
    def: T,
}

impl<'a, T> Bar<'a> 
    where T: SomeTrait   // <-- this is where we do it
    {
    pub fn some_function(&self)  {
    self.def.foo();
    }
}
```

There are only three instances where you are required to put a bound in the struct itself:

when the object has a field of an associated type
(see std::borrow::Cow 6 which has a ToOwned bound because it actually may contain the associated type <B as ToOwned>::
Owned)
(similarly Peekable 2 holds a <T as Iterator>::Item)
when that trait is used in the destructor
(examples are rare; I could only find one 1 in the standard library; that type needs Ord because its drop impl calls a
method that requires Ord)
when you carry another struct that has a bound. (see? It's infectous!)

## Rust Dependency Injection implementation of trait into another object

```rust
pub trait Downloader {
    fn run(&self);
}

pub struct DownloaderFoo;
impl Downloader for DownloaderFoo {
    fn run(&self)  {
        println!("DownloaderFoo::run");
    }
}

pub struct DownloaderBar;
impl Downloader for DownloaderBar {
    fn run(&self)  {
        println!("DownloaderBar::run");
    }
}

pub struct Installer {
    downloader: Box<dyn Downloader>,
}

impl Installer {
    pub fn run(&self) {
        self.downloader.run();
        println!("Installer::run");
    }
}

fn main() {
    Installer{downloader: Box::new(DownloaderFoo{})}.run();
    Installer{downloader: Box::new(DownloaderBar{})}.run();
}
```

