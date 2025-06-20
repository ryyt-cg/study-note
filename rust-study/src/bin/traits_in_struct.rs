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
            data: "Hello, Data!".to_string(),
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

#[tokio::main]
async fn main() {
    let repo = Box::new(Repo::new());
    let service = Service::new(repo);

    println!("{}", service.get());
}