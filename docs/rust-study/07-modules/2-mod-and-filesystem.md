# mod and the Filesystem

```bash
cargo new --lib communicator
cd communicator
```

library crate communicator

```
communicator
|- Cargo.toml
|- src
  |- lib.rs
```

cargo new --bin hello-rust</br>
binary crate hello-rust

```
hello-rust
|- Cargo.toml
|- src
  |- main.rs
```

cargo new communicator - creates a crate called communicator.

* repo could contain 1 or more crates
* crate contains 1 or more modules
* module could contain modules, functions, enums and structs

Filename: src/lib.rs

```
mod network {
    fn connect() {
    }
}

mod client {
    fn connect() {
    }
    mod server {
        fn connect() {
        }
    }
}
```

```asciidoc
communicator
├── client
└── network
    └── server
```

If these modules had many functions, and those functions were becoming lengthy, it would be difficult to scroll through
this file to find the code we wanted to work with. Because the functions are nested inside one or more mod blocks, the
lines of code inside the functions will start getting lengthy as well. These would be good reasons to separate the
client, network, and server modules from src/lib.rs and place them into their own files.

```asciidoc
mod client;

mod network {
    fn connect() {
    }

    mod server {
        fn connect() {
        }
    }
}
```

src/client.rs

```rust
fn connect()  {
}
```

src/lib.rs

```rust
mod client;
mod network;
```

src/network.rs

```rust
fn connect() {
}

mod server {
    fn connect() {
    }
}
```

Instead of continuing to follow the same file naming pattern we used previously, we can do what the note suggests:

1. Make a new directory named network, the parent module’s name.
2. Move the src/network.rs file into the new network directory, and rename it to src/network/mod.rs.
3. Move the submodule file src/server.rs into the network directory.

Here are commands to carry out these steps:

```bash
$ mkdir src/network
$ mv src/network.rs src/network/mod.rs
$ mv src/server.rs src/network
```

The corresponding file layout now looks like this:

```asciidoc
├── src
│   ├── client.rs
│   ├── lib.rs
│   └── network
│       ├── mod.rs
│       └── server.rs
```

### Rules of Module Filesystems

Let’s summarize the rules of modules with regard to files:

* If a module named foo has no submodules, you should put the declarations for foo in a file named foo.rs.
* If a module named foo does have submodules, you should put the declarations for foo in a file named foo/mod.rs.


