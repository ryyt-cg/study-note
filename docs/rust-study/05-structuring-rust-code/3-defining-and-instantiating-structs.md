# Defining and Instantiating Structs

## Defining struct

* Use the struct keyword
* Give the struct name, starting with a capital
* Define field names and types

```rust
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}
```

Instantiating struct

```rust
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

fn main() {
    let user1 = User {
        email: String::from("someone@examples.com"),
        username: String::from("someusername123"),
        active: true,
        sign_in_count: 1,
    };
}

```

To get a specific value from a struct, we can use dot notation. If we wanted just this user’s email address, we can use
user1.email wherever we want to use this value. If the instance is mutable, we can change a value by using the dot
notation and assigning into a particular field. Listing 5-3 shows how to change the value in the email field of a
mutable User instance:

```rust
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

fn main() {
    let mut user1 = User {
        email: String::from("someone@examples.com"),
        username: String::from("someusername123"),
        active: true,
        sign_in_count: 1,
    };
    user1.email = String::from("anotheremail@examples.com");
    z
}

```

build_user function that returns a User instance with the given email and username. The active field gets the value of
true, and the sign_in_count gets a value of 1.

```rust
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

fn build_user(email: String, username: String) -> User {
    User {
        email: email,
        username: username,
        active: true,
        sign_in_count: 1,
    }
}
```

## Passing struct parameters by value

It depends on struct implementing Copy trait or not

* if implementing Copy trait
    * a bitwise copy of the struct is passed to the called function
    * The caller retains ownership and can use the struct afterwards
* if not implementing Copy trait
    * Ownership of the struct is moved to the called function
    * The caller loses ownership and can't use the struct afterwards

```rust
use crate::my_types::Employee;

pub fn do_it() {
    println!("\nIn demo_struct_pass_value::do_it()");

    let e1 = Employee {
        name: String::from("John"),
        salary: 1000,
        fulltime: false,
    };

    // Pass struct by value (and move ownership into the function).
    consume_employee(e1);

    // Can't use e1 now...
    // println!("{}", e1.name);
}

fn consume_employee(e: Employee) {
    println!(
        "{} earns {}, fulltime status: {}",
        e.name, e.salary, e.fulltime
    );
} // Employee object dropped here.
```

## Passing struct parameters as by reference

```rust
use crate::mytypes::Employee;

pub fn do_it() {
    println!("\nIn demo_struct_pass_reference::do_it()");

    let mut e1 = Employee {
        name: String::from("Jane"),
        salary: 1000,
        fulltime: true,
    };

    // Pass struct by reference.
    print_employee(&e1);

    // Pass struct by mutable reference.
    reward_employee(&mut e1);
    print_employee(&e1);
}

fn print_employee(e: &Employee) {
    println!(
        "Using explicit dereferencing: {} earns {}, fulltime status: {}",
        (*e).name,
        (*e).salary,
        (*e).fulltime
    );
    println!(
        "Using implicit dereferencing: {} earns {}, fulltime status: {}",
        e.name, e.salary, e.fulltime
    );
}

fn reward_employee(e: &mut Employee) {
    // Using explicit dereferencing.
    (*e).salary += 500;

    // Using implicit dereferencing.
    e.salary += 250;
}
```

## Return struct as value

```rust
use crate::mytypes::Employee;

pub fn do_it() {
    println!("\nIn demo_struct_return_value::do_it()");

    let e1 = build_employee(String::from("Jane"), 1000, true);
    print_employee(&e1);

    let mut e2 = build_employee_v2(String::from("John"), 1000, false);
    e2.salary += 750;
    print_employee(&e2);
}

fn build_employee(name: String, salary: u64, fulltime: bool) -> Employee {
    Employee {
        name: name,
        salary: salary,
        fulltime: fulltime,
    }
}

fn build_employee_v2(name: String, salary: u64, fulltime: bool) -> Employee {
    Employee {
        name,
        salary,
        fulltime,
    }
}

fn print_employee(emp: &Employee) {
    println!(
        "{} earns {}, fulltime status: {}",
        emp.name, emp.salary, emp.fulltime
    );
}
```

## Return struct as reference

* <'a> is a lifetime annotation, letter a can be anything name, like b, c or z.
* We apply 'a to each reference param to indicate that they have the same lifetime.

```rust
use crate::mytypes::Employee;

pub fn do_it() {
    println!("\nIn demo_struct_return_reference::do_it()");

    let mut e1 = build_employee(String::from("Jane"), 1001, true);
    let mut e2 = build_employee(String::from("John"), 1000, false);

    let ri = choose_employee(&e1, &e2);
    print_employee(ri);

    let rm = choose_mut_employee(&mut e1, &mut e2);
    rm.salary *= 2;
    print_employee(rm);
}

fn build_employee(name: String, salary: u64, fulltime: bool) -> Employee {
    Employee {
        name,
        salary,
        fulltime,
    }
}

fn choose_employee<'a>(e1: &'a Employee, e2: &'a Employee) -> &'a Employee {
    if e1.salary > e2.salary {
        e1
    } else {
        e2
    }
}

fn choose_mut_employee<'a>(e1: &'a mut Employee, e2: &'a mut Employee) -> &'a mut Employee {
    if e1.salary > e2.salary {
        e1
    } else {
        e2
    }
}

fn print_employee(e: &Employee) {
    println!(
        "{} earns {}, fulltime status: {}",
        e.name, e.salary, e.fulltime
    );
} 
```



