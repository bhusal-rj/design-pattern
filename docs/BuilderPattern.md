# Builder Design Pattern
- Creational Design Pattern.
- Helps in creating complex objects step by step.
- In builder pattern we extractaway the object construction code out of its own class and move it to separate objects called builders.
- Like for creating a house we can extract away creation of walls, door, floors into the builders.
- Here builders are the factory function which creates the corresponding objects.
- Now method chaining is used to create the entire class from these builds.


## Use Case Scenario
- Imagine having a complex object the requires step by step initialization of many fields and nested objects. To initialize these objects we need a monstrous constructor with lots of parameters.
- What if in future we need to add some features to the class. Dealing with the monstrous constructor will be extremely hard and unmaintainable so what we can do is just create the subclasses so adding the features within the class will just require growing the hierarchy.
- For example, while creating a house there are some properites which are not required for many houses like chimney, swimming pool etc. Initializing the normal houses will make the parameters unused making the constructor call ugly.

