# IoC in Rust

Inversion of Control (IoC) is a design principle in which the flow of control of a system is inverted. In IoC, the
control of the program is transferred to a container or framework. The container or framework manages the dependencies
of the application. The IoC principle is implemented in Rust using traits and generics.

In service.rs,

```rust
// Write me a sample trait and struct that uses the trait as a field.
pub trait Repoer {
    fn get(&self) -> String;
}

pub struct Repo {
    data: String,
}

impl Repo {
    pub fn new() -> Self {
        Repo {
            data: "Hello, World!".to_string(),
        }
    }
}

impl Repoer for Repo {
    fn get(&self) -> String {
        self.data.clone()
    }
}

pub trait Servicer {
    fn get(&self) -> String;
}

pub(crate) struct Service {
    repo: Box<dyn Repoer + 'static>,
}

impl Service {
    pub fn new(repo: Box<dyn Repoer + 'static>) -> Self {
        Service {
            repo,
        }
    }

}
impl Servicer for Service {

    fn get(&self) -> String {
        self.repo.get()
    }
}
```

main.rs

```rust
mod service;
use crate::service::{Servicer};

#[tokio::main]
async fn main() {
    let repo = Box::new(crate::service::Repo::new());
    let service = crate::service::Service::new(repo);

    println!("{}", service.get());
}
```

This code won't work if implement async traits to 'async fn get(&self)'. The error message is:

```bash
error[E0038]: the trait `Repoer` cannot be made into an object
  --> src/service.rs:29:15
   |
29 |     repo: Box<dyn Repoer + 'static>,
   |               ^^^^^^^^^^^^^^^^^^^^ `Repoer` cannot be made into an object
   |
note: for a trait to be "object safe" it needs to allow building a vtable to allow the call to be resolvable dynamically; for more information visit <https://doc.rust-lang.org/reference/items/traits.html#object-safety>
  --> src/service.rs:3:14
   |
2  | pub trait Repoer {
   |           ------ this trait cannot be made into an object...
3  |     async fn get(&self) -> String;
   |              ^^^ ...because method `get` is `async`
   = help: consider moving `get` to another trait
   = help: only type `service::Repo` implements the trait, consider using it directly instead
```

To fix this, we need to split the trait into two traits, one for async and one for sync. The code will look like this:
in service.rs

```rust
// Write me a sample trait and struct that uses the trait as a field.
pub trait Repoer {
    async fn get(&self) -> String;
}

pub struct Repo {
    data: String,
}

// Inject constructor - inject the dependency (DB) into the struct
impl Repo {
    pub fn new() -> Self {
        Repo {
            data: "Hello, DB!".to_string(),
        }
    }
}

impl Repoer for Repo {
    async fn get(&self) -> String {
        self.data.clone()
    }
}
    
pub trait Servicer {
    async fn get(&self) -> String;
}

pub struct Service<R> {
    repo: R,
}

// Constructor injection - inject the dependency (repo) into the struct
impl<R> Service<R> {
    pub fn new(repo: R) -> Self{
        Service {
            repo,
        }
    }

}
impl<R> Servicer for Service<R> where R: Repoer {
    async fn get(&self) -> String {
        self.repo.get().await
    }
}
```

in main.rs

```rust
#[tokio::main]
async fn main() {
    let repo = service::Repo::new();
    let service = crate::service::Service::new(repo);

    println!("{}", service.get().await);
}
```

This example demonstrates how to implement IoC in Rust using traits and generics.</p>
Service has two dependencies, Repo and Gateway. The Repo and Gateway are injected into the Service struct using
constructor injection. The Service struct implements the Servicer trait, which has a get method that returns a string.
The get method calls the get method of the Repo and Gateway dependencies and concatenates the results.

```rust
// Write me a sample trait and struct that uses the trait as a field.
pub trait Repoer {
    async fn get(&self) -> String;
}

pub struct Repo {
    data: String,
}

// Inject constructor - inject the dependency (DB) into the struct
impl Repo {
    pub fn new() -> Self {
        Repo {
            data: "Hello, Data!".to_string(),
        }
    }
}

impl Repoer for Repo {
    async fn get(&self) -> String {
        self.data.clone()
    }
}

trait Gatewayer {
    async fn get(&self) -> String;
}

pub struct Gateway {}

impl Gateway {
    pub fn new() -> Self {
        Gateway {}
    }
}

impl Gatewayer for Gateway {
    async fn get(&self) -> String {
        "Hello, Gateway".to_string()
    }
}

pub trait Servicer {
    async fn get(&self) -> String;
}

pub struct Service<R, G> {
    repo: R,
    gateway: G
}

// Constructor injection - inject the dependency (repo) into the struct
impl<R, G> Service<R, G> {
    pub fn new(repo: R, gateway: G) -> Self{
        Service {
            repo,
            gateway,
        }
    }

}
impl<R, G> Servicer for Service<R, G> where R: Repoer, G: Gatewayer {
    async fn get(&self) -> String {
        self.repo.get().await + "  " + self.gateway.get().await.as_str()
    }
}

#[tokio::main]
async fn main() {
    let repo = service::Repo::new();
    let gateway = service::Gateway::new();
    let service = crate::service::Service::new(repo, gateway);

    println!("{}", service.get().await);
}
```