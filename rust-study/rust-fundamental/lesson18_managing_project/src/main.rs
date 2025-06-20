fn main() {
    // absolute path
    crate::hello::english();
    crate::hello::spanish();
    crate::hello::casual::english();

    // relative path
    hello::english();
    hello::casual::english();
}

// inline module
mod hello {
    pub fn english() {
        println!("Hello");
    }

    pub fn spanish() {
        println!("hola");
    }

    // Inner nested modules have access to outer module,
    // but outer module does not have access to inner modules.
    pub mod casual {
        pub fn english() {
            println!("hey");
        }
    }
}

//  cargo check
//     Checking lesson18_managing_project v0.1.0                                                                                                                                                                  ─╯
// error[E0603]: function `english` is private
//  --> src/main.rs:2:19
//   |
// 2 |     crate::hello::english();
//   |                   ^^^^^^^ private function
//   |

// 6  |     crate::hello::casual::english();
//    |                   ^^^^^^  ------- function `english` is not publicly re-exported
//    |                   |
//    |                   private module
