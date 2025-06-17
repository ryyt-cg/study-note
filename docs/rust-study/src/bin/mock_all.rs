#[cfg(test)]
use mockall::{automock, predicate::*};

#[cfg_attr(test, automock)]
pub trait Database {
    fn execute_query(&self, query: String);
}

pub fn get_user(db: Box<dyn Database>, id: i32) {
    let query = format!("SELECT * FROM users WHERE id = {}", id);
    db.execute_query(query);
}

#[cfg(test)]
mod tests {
    use mockall::predicate::eq;
    use crate::MockDatabase;

    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }

    #[test]
    fn get_user_test() {
        let mut db = Box::new(MockDatabase::new());
        db.expect_execute_query()
            .with(eq("SELECT * FROM users WHERE id = 1".to_string()))
            .times(1)
            .return_const(());
        super::get_user(db, 1);
    }
}

// mockall async
// https://github.com/mibes/mockall-async
fn main() {
    println!("Hello, World!");
}