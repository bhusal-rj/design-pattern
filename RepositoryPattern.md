# Repository Design Pattern
- It is a software design pattern which acts as an intermediary layer between an application business logic and data storage.
- It's main purpose is to provide the structured and standarized way to access, manage and manipulate data while abstracting the underlying details of data storage technologies.
- Let's say today we are using the **Mysql** database tomorrow we wish to switch to **Mongodb**. In normal program flow we need to change the business logic if the repository pattern is not implemented.


## Analogy
- In a library to find a book we donot need to goto storage room to search for book ourself, we can ask the librarian to help find a book. The librarian knows where the books are kept and can give the book without having to worry about where it is placed. In application program, instead of program looking directly for data, it asks for repository to find save and delete the data it needs.
