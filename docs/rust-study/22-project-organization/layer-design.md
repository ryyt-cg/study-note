The first approach would look something like :

trait UserService { ... }
struct UserServiceImpl {
repository: Arc<dyn UserRepository>
}
impl UserService for UserServiceImpl { .. }

trait UserRepository { ... }
struct UserRepositoryImpl { ... }
impl UserRepository for UserRepositoryImpl
The second would be :

trait UserService { ... }
struct UserServiceImpl<R : UserRepository> {
repository: Arc<R>
}
impl<R: UserRepository> UserService for UserServiceImpl<R> { .. }