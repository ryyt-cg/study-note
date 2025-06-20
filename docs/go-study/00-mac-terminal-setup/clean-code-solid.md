# Clean Code

## SOLID Principles

SOLID principles are 5 concepts that will help us make our programs better.

* Single Responsibility Principle
* Open / Closed Principles
* Liskov Substitution Principle
* Interface Segregation Principle
* Dependency Inversion Principle

### Single Responsibility Principle

A class should have one, and only one, reason to change. This means that a class should have only one job.

### Open / Closed Principles

A class should be open for extension but closed for modification. This means that we should be able to add new
functionality to a class without changing the existing code.

### Liskov Substitution Principle

Objects of a superclass should be replaceable with objects of its subclasses without affecting the correctness of the
program. This means that we should be able to use a subclass in place of a superclass without any problems.

### Interface Segregation Principle

A class should not be forced to implement an interface that it does not use. This means that we should have small,
specific interfaces instead of large, general ones.

### Dependency Inversion Principle

High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not
depend on details. Details should depend on abstractions.
