# Factory Pattern
- It is the one of the most simplest design pattern.
- It is often regarded as the creational design pattern which is used to create the objects of similar type.
- It hides the complexity of object creation.
- Client code has no idea about the object creation. It will just invoke the **new()** or similar function for object creation.
- Client will not know the complexities in creating the objects as all the complexities are handled by the classes.


## Design of Factory Pattern

                                          Inherits
**Factory Class**(Abstract Interface) <-------------- **Concrete Class**(Real Implementation)

## Example
- In a logistic management application, the initial version of the application can only handle transactions by Trucks. In future we wish to add the **transportation** by trucks as well. At present, most of our code is coupled to the class Truck. Adding another transportation will require making changes to the entire codebase resulting in unmaintainable code.
- 


## Pros and Cons
- Gurantee abstraction.
- Code is flexible and adaptable.
- Very useful for frameworks and libraries.
- Complex code
- Takes time to set the base.
- Not a pattern that canbe refactored into.
