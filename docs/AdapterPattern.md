# Adapter Pattern
- It is the structural design pattern that allows the objects with incompatible interfaces to collaborate.
- Allows us to have different programs to communicate with one another.
- Allows existing source of code work with other part without modifying the actual code.
- Let's imagine we are building a Stock Market application that displays the stock market analytics with the XML data. Later, we try to expand the analytics with JSON data as well. One approach to handle this will be to convert the analytics app to support JSON instead of XML which incase in future if we needed to go back to XML will create a hardtime and another approach will be write the wrapper that converts the JSON data to XML so that analytics app can understand
