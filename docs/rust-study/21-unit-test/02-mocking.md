# Mocking in Rust

To test a component in isolation, we need to mock the dependencies of the component. Mocking is a technique used in unit
testing when a component is dependent on another component that is difficult to test. Mocking is used to isolate the
component under test from the rest of the system. Mocking is used to simulate the behavior of the dependent component.

- Mocking is a technique used in unit testing when a component is dependent on another component that is difficult to
  test.
- Mocking is used to isolate the component under test from the rest of the system.
- Mocking is used to simulate the behavior of the dependent component.

In a unit test, mock object can simulate teh behavior of complex, real objects and are therefore useful when a read
object is impractical or impossible to incorporate into a unit test.

Mon-Deterministic Resources

* External APIs
* Databases
* Resources using global state
* Time-dependent resources
* Environment variables

Mocking would control the behavior of these resources and make the test deterministic.

## Reason for Mocking

* Isolation of Components for Unit Testing
* Improved Test Stability & Reliability
* Efficiency and Speed

## Mockall

There are several libraries available in Rust for mocking.
One of the popular libraries is Mockall.
Mockall is a Rust library for mocking traits.
Mockall generates mock objects for traits at compile time.

add mockall crate to [Cargo.toml](https://doc.rust-lang.org/cargo/reference/specifying-dependencies.html)

```toml
mockall = {version = "0.12.1"}
```

* Must refactor to traits implementation in order to mock
* Use `Mockall` crate to generate mock objects
* Use Mock<trait_name> to create a mock object

for example:

```rust
use std::arch::{is_arm_feature_detected, is_mips_feature_detected};
use mockall::automock;

#[automock]
pub trait ProductServicer {
    fn get_product_by_id(&self, id: i32) -> Result<Product, Error>;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_product_by_id() {
        let mut product_service = ProductServicer::new();
        product_service.expect_run().times(1).returning(|| println!("ProductServicer::run"));
        
        product_service.get_product_by_id(1);
    }
}

fn main() {
    let product_service = MockProductServicer::new();
    let product = product_service.get_product_by_id(1);
}
```










