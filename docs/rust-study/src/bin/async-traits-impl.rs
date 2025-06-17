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

#[tokio::main]
async fn main() {
    let repo = Repo::new();
    let service = Service::new(repo);

    println!("{}", service.get().await);
}