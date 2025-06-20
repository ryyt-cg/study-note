// Write me a sample trait and struct that uses the trait as a field.
trait Repoer {
    async fn get(&self) -> String;
}

pub struct Repo {
    data: String,
}

// Inject constructor - inject the dependency (DB) into the struct
impl Repo {
    pub fn new() -> Self {
        Repo {
            data: "Hello, Data!".to_string(),
        }
    }
}

impl Repoer for Repo {
    async fn get(&self) -> String {
        self.data.clone()
    }
}

trait Gatewayer {
    async fn get(&self) -> String;
}

struct Gateway {}

impl Gateway {
    pub fn new() -> Self {
        Gateway {}
    }
}

impl Gatewayer for Gateway {
    async fn get(&self) -> String {
        "Hello, Gateway".to_string()
    }
}

trait Servicer {
    async fn get(&self) -> String;
}

pub struct Service<R, G> {
    repo: R,
    gateway: G
}

// Constructor injection - inject the dependency (repo) into the struct
impl<R, G> Service<R, G> {
    fn new(repo: R, gateway: G) -> Self{
        Service {
            repo,
            gateway,
        }
    }

}
impl<R, G> Servicer for Service<R, G> where R: Repoer, G: Gatewayer {
    async fn get(&self) -> String {
        self.repo.get().await + "  " + self.gateway.get().await.as_str()
    }
}

#[tokio::main]
async fn main() {
    let repo = Repo::new();
    let gateway = Gateway::new();
    let service = Service::new(repo, gateway);

    println!("{}", service.get().await);
}