package sevices

import (
	"fmt"
	"library_management/models"
	"regexp"
)

func helper(name string) string {
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	for !re.MatchString(name) {
		fmt.Print("Enter valid  author:  ")
		_, _ = fmt.Scanln(&name)

	}
	return name
}

func findIndex(slice []models.Book, target models.Book) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	// Return -1 if the target is not found in the slice
	return -1
}

var bookStore map[int]models.Book
var members map[int]models.Member

func AddBook(book models.Book) string {
	fmt.Println("Enter book ID:  ")
	var id int
lable:
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Print("Invalid id please enter integer only:  ")
		goto lable
	}
	book.Id = id

	fmt.Print("Enter the book title:  ")
	var title string
	_, _ = fmt.Scanln(&title)
	book.Title = title

	fmt.Print("Enter  author of the book:  ")
	var author string
	_, _ = fmt.Scanln(&author)
	book.Author = helper(title)

	book.Status = "Available"
	bookStore[book.Id] = book

	return "The book is added succesfully :"

}

func RemoveBook(bookID int) string {
	if _, ok := bookStore[bookID]; ok {
		delete(bookStore, bookID)
		return "The book succesfully removed:  "

	} else {
		return ("The book doesnot exist:  ")

	}

}
func BorrowBook(bookID int, memberid int) string {
	if _, ok := members[memberid]; !ok {
		fmt.Print("Enter your name to register:  ")
		var name string
		_, _ = fmt.Scanln(&name)
		newMember := models.Member{}
		newMember.Name = helper(name)
		newMember.Id = memberid
		members[memberid] = newMember

	}
	if _, ok := bookStore[bookID]; ok {
		if bookStore[bookID].Status == "Borrowed" {
			return "The book is not available now:  "
		}

		mem := members[memberid]
		mem.BorrowedBooks = append(mem.BorrowedBooks, bookStore[bookID])
		boo := bookStore[bookID]
		boo.Status = "Borrowed"
		bookStore[bookID] = boo
		return "succesfully borowed"

	} else {
		return ("The book doesnot exist:  ")

	}

}
func ReturnBook(bookid int, memberid int) string {
	if _, ok := members[memberid]; !ok {
		return ("You are not a member :  ")

	}
	if _, ok := bookStore[bookid]; !ok {
		return ("The book doesnot exist:  ")

	}

	if bookStore[bookid].Status == "Borrowed" {
		mem := members[memberid]
		borrowed := mem.BorrowedBooks
		boo := bookStore[bookid]
		idx := findIndex(mem.BorrowedBooks, boo)
		if idx == -1 {
			return "Sorry you didn't borrow it here!:  "
		}

		mem.BorrowedBooks = append(borrowed[:idx], borrowed[idx+1:]...)
		boo.Status = "Available"
		bookStore[bookid] = boo
		return "succesfully returned :  "

	}

	return "The book is not borrowed"

}
func ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range bookStore {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks

}
func Checker(memberID int) bool {
	if _, ok := members[memberID]; !ok {
		return false
	}
	return true
}
func ListBorrowedBooks(memberID int) []models.Book {

	borrowed := members[memberID]
	return borrowed.BorrowedBooks

}
