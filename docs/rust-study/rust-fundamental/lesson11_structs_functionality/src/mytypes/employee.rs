use std::sync::atomic::{AtomicI32, Ordering};

pub struct Employee {
    id: i32,
    name: String,
    full_time: bool,
    salary: u64,
}

static NEXT_ID: AtomicI32 = AtomicI32::new(0);

impl Employee {
    const MAX_SALARY: u64 = 99_000;

    pub fn to_string(&self) -> String {
        format!(
            "[{}], {} earns {}, full_time status: {}",
            self.id, self.name, self.salary, self.full_time
        )
    }

    pub fn payrise(&mut self, amount: u64) {
        self.salary += amount;
        if self.salary > Employee::MAX_SALARY {
            // Can also say Self::MAX_SALARY
            self.salary = Employee::MAX_SALARY;
        }
    }

    pub fn new(name: String, salary: u64, full_time: bool) -> Employee {
        let id = NEXT_ID.fetch_add(1, Ordering::Relaxed);
        Employee {
            id,
            name,
            salary,
            full_time,
        }
    }
}
