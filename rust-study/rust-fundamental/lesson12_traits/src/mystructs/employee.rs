use crate::mytraits::log::Log;
use crate::mytraits::print::Print;

pub struct Employee {
    name: String,
    salary: u64,
    full_time: bool,
}

impl Employee {
    pub fn payrise(&mut self, amount: u64) {
        self.salary += amount
    }

    pub fn new(name: String, salary: u64, full_time: bool) -> Employee {
        Employee {
            name,
            salary,
            full_time,
        }
    }
}

impl Print for Employee {
    fn print(&self) {
        println!(
            "{} earns {}, full_time: {}",
            self.name, self.salary, self.full_time
        );
    }
}

impl Log for Employee {
    const LOG_TIMESTAMP: bool = true;

    fn log(&self) {
        println!("{}, {}, {}", self.name, self.salary, self.full_time);
    }
}
