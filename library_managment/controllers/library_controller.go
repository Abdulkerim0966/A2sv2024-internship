package controllers

import (
	"fmt"
	"library_management/models"
	sevices "library_management/services"
)

func Controller() {
	fmt.Println("=======================================")
	fmt.Println("|                                     |")
	fmt.Println("|   WELL COME TO LIBRARY MANEGMENT    |")
	fmt.Println(`|                                     |`)
	fmt.Println("=======================================")

	fmt.Println("========================")
	fmt.Println("Enter 1 to Add Book:")
	fmt.Println("Enter 2 to Remove Book:")
	fmt.Println("Enter 3 to Borrow Book:")
	fmt.Println("Enter 4 to Return Book: ")
	fmt.Println("Enter 5 to get List of Available Books:")
	fmt.Println("Enter 6 to  get List Borrowed Books:")
	fmt.Println("Enter 0 to exit the Program:")
	fmt.Println("========================")

    lable1:
	fmt.Print("Please enter your choice: ")
	var input int
	_, err := fmt.Scanln(&input)
	
	if err != nil || input > 6 || input < 0 {
		fmt.Println("Invalid !!")
		Controller()
	}
	switch input {
	case 0:
		fmt.Println("The program succesfully closed:")
	
	case 1:
		fmt.Println("You chose to Add a Book.")
		newBook := models.Book{}
		fmt.Print("\n\n",sevices.AddBook(newBook),"\n\n")
		Controller()

	case 2:
		fmt.Println("You chose to Remove a Book.")
		fmt.Println("Enter the Id of the book")
		var id int
	lable:
		_, err := fmt.Scanln(&id)
		if err != nil {
			fmt.Print("Invalid id please enter integer only:  ")
			goto lable
		}
		fmt.Print("\n\n",sevices.RemoveBook(id),"\n\n")
	    Controller()

	case 3:
		fmt.Println("You chose to Borrow a Book.")
		fmt.Println("Enter book ID:  ")
		var id int
		
	lable4:
		_, err := fmt.Scanln(&id)
		if err != nil {
			fmt.Print("Invalid id please enter integer only:  ")
			goto lable4
		}

		fmt.Println("Enter book ID:  ")
		var memberid int
	lable5:
		_, er := fmt.Scanln(&memberid)
		if er != nil {
			fmt.Print("Invalid id please enter integer only:  ")
			goto lable5
		}
		fmt.Print("\n\n",sevices.BorrowBook(id, memberid),"\n\n")
		Controller()
	case 4:
		fmt.Println("You choose to Return a Book.")

		fmt.Println("Enter book ID:  ")
		var id int
	lable6:
		_, err := fmt.Scanln(&id)
		if err != nil {
			fmt.Print("Invalid id please enter integer only:  ")
			goto lable6
		}

		fmt.Println("Enter member ID:  ")
		var memberid int
	lable7:
		_, er := fmt.Scanln(&memberid)
		if er != nil {
			fmt.Print("Invalid id please enter integer only:  ")
			goto lable7
		}
		fmt.Print("\n\n",sevices.ReturnBook(id, memberid),"\n\n")
		Controller()
	case 5:
		fmt.Println("You chose to List Available Books.")
		fmt.Print("\n\n",sevices.ListAvailableBooks(),"\n\n")
        Controller()
	case 6:
		fmt.Println("You chose to List Borrowed Books.")
		fmt.Println("Enter member ID:  ")
		var memberid int
	lable8:
		_, er := fmt.Scanln(&memberid)
		if er != nil {
			fmt.Print("Invalid id please enter integer only:  ")
			goto lable8
		}
		if !sevices.Checker(memberid) {
			fmt.Print("\n\n",sevices.ListBorrowedBooks(memberid),"\n\n")

		} else {

			fmt.Println("You are not a member :  ")

		}
        Controller()

	default:
		fmt.Println("Invalid choice. Please enter a number between 0 and 6.")
		goto lable1
	}


}
