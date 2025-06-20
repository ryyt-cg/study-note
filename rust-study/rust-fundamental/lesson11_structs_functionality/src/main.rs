#![allow(dead_code)]

mod demo_associated_data;
mod demo_associated_functions;
mod demo_modular_code;
mod demo_mutable_impl;
mod demo_simple_impl;
mod mytypes;

fn main() {
    // Uncomment one of these statements, to run a demo...

    // demo_simple_impl::do_it();
    // demo_mutable_impl::do_it();
    // demo_modular_code::do_it();
    // demo_associated_functions::do_it();
    demo_associated_data::do_it();
}
