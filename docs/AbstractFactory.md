# Abstract Factory Method
- It is the creational method whereby object which belong to a family of similar objects is created.
- Also known as **Factory of factory pattern**.
- Implementation using a common interface.
- It is the one more additional layer that decides which factory pattern to call on.
- It is mainly to create the family of related objects.
- We mentioned in [Factory Pattern](/docs/FactoryPattern.md) the challenges on implementing different return objects from the concrete classes which can be solved using the abstract factory pattern.

## Design of Abstract Factory Pattern

```
Abstract Factory Interface (Abstract Interface) <------ Concrete Factory Interface(Logical layer) <------ Concrete Object Class(Real Implementation)

```
- `A <----- B  Represents the Inheritance from B to A`



## Analogy 
- There might be a client who wishes to order **Mobile Phone** from Apple. We place the order in Apple but all the components required to prepare that mobile phone isnot prepared by Apple instead it orders other **factories** for the different components and these different factories prepares the product for the company. As it is upon the **Apple** which factory to call to get the required object, apple here can be called as an **Abstract Factory**
- Here, Apple is a abstract factory and **Screen factory,Processor Factory**etc are the concrete factories.
