Project Structure
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   ├── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod

    Main File (main.go)
This is the entry point of the application.
    Description:

    The main package imports the controllers and services packages.
    It initializes the library using sevices.Newlibrary().
    It calls the Controller() function from the controllers package to start the user interaction.


    Controllers (controllers/library_controller.go)

This file contains the controller logic for handling user inputs and managing the flow of the application.

Description:

    The Controller() function provides a menu-driven interface for interacting with the library management system.
    It allows users to add, remove, borrow, return books, and list available and borrowed books.
    It uses the sevices package for the actual operations on the library.


Models (models/book.go)
This file contains the Book struct.
Description:
The Book struct defines the attributes of a book in the library: Id, Title, Author, and Status.
Models (models/member.go)
This file contains the Member struct.
Description:
The Member struct defines the attributes of a member in the library: Id, Name, and a list of BorrowedBooks.


    Services (services/library_service.go)

This file contains the service functions for library operations.
Description:

Newlibrary: Initializes a new library with empty bookStore and members.
AddBook: Adds a book to the library.
RemoveBook: Removes a book from the library.
BorrowBook: Allows a member to borrow a book from the library.
ReturnBook: Allows a member to return a borrowed book to the library.
ListAvailableBooks: Lists all available books in the library.
Checker: Checks if a member exists in the library.
ListBorrowedBooks: Lists all borrowed books of a member.


Usage Instructions
Run the Application:

Run the main.go file to start the library management system.
Main Menu:

The application will display a menu with options to add, remove, borrow, return books, and list available and borrowed books.
Follow the on-screen prompts to interact with the system.
Select option 1 to add a book.
Select option 2 to remove a book.
Select option 3 to borrow a book.
Select option 4 to return a book..
Select option 5 to list all available books in the library.
Select option 6 to list all borrowed books of a member.
Enter the member ID to view the borrowed books
Select option 0 to exit the program.


Additional Information
The system uses maps to store books and members, allowing for efficient lookups.
Regular expressions are used to validate author names.
Functions are provided to handle various library operations, ensuring modularity and separation of concerns.
