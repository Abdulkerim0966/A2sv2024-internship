package sevices

import (
	"bufio"
	"fmt"
	"library_management/models"
	"os"
	"regexp"
	"strings"
)

func getInput(prompt string ,r *bufio.Reader)(string,error){
    fmt.Print(prompt)
    input,err :=r.ReadString('\n')
    return strings.TrimSpace(input),err
}

func helper(name string) string {
	re := regexp.MustCompile(`^[a-zA-Z\s]+(,[a-zA-Z\s]+)*$`)
	for !re.MatchString(name) {
		fmt.Print("Enter valid  author, if they are multiple separatethem by comma   ")
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
type LIBRARY struct{

	bookStore map[int]models.Book
	members map[int]models.Member
}
var newlib LIBRARY

func Newlibrary (){
	newlib.bookStore =make(map[int]models.Book)
	newlib.members =make(map[int]models.Member)

}


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

	var title string
	reader := bufio.NewReader(os.Stdin)
    title,_ = getInput("Enter the book title:  ",reader)	
	book.Title = title

	var author string
	author,_ = getInput("Enter  author of the book:  ",reader)
	book.Author = helper(author)
	book.Status = "Available"
	newlib.bookStore[book.Id] = book

	return "The book is added succesfully :"

}

func RemoveBook(bookID int) string {
	if _, ok := newlib.bookStore[bookID]; ok {
		delete(newlib.bookStore, bookID)
		return "The book succesfully removed:  "

	} else {
		return ("The book doesnot exist:  ")

	}

}
func BorrowBook(bookID int, memberid int) string {
	newlib.members = make(map[int]models.Member)
	if _, ok := newlib.members[memberid]; !ok {

		var name string
		reader := bufio.NewReader(os.Stdin)
    	name,_ = getInput("Enter your name to register:  ",reader)	
		newMember := models.Member{}
		newMember.Name = helper(name)
		newMember.Id = memberid
		newlib.members[memberid] = newMember

	}
	if _, ok := newlib.bookStore[bookID]; !ok {
		return ("The book doesnot exist:  ")

	}

	if newlib.bookStore[bookID].Status == "Borrowed" {
		return "The book is not available now:  "
	}

	mem := newlib.members[memberid]
	mem.BorrowedBooks = append(mem.BorrowedBooks, newlib.bookStore[bookID])
	boo := newlib.bookStore[bookID]
	boo.Status = "Borrowed"
	newlib.bookStore[bookID] = boo
	return `succesfully borrowed`

}
func ReturnBook(bookid int, memberid int) string {
	if _, ok := newlib.members[memberid]; !ok {
		return ("You are not a member :  ")

	}
	if _, ok := newlib.bookStore[bookid]; !ok {
		return ("The book doesnot exist:  ")

	}

	if newlib.bookStore[bookid].Status == "Borrowed" {
		mem := newlib.members[memberid]
		borrowed := mem.BorrowedBooks
		boo := newlib.bookStore[bookid]
		idx := findIndex(mem.BorrowedBooks, boo)
		if idx == -1 {
			return "Sorry you didn't borrow it here!:  "
		}

		mem.BorrowedBooks = append(borrowed[:idx], borrowed[idx+1:]...)
		boo.Status = "Available"
		newlib.bookStore[bookid] = boo
		return "succesfully returned :  "

	}

	return "The book is not borrowed"

}
func ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range newlib.bookStore {
		if book.Status == "Available" {	
			availableBooks = append(availableBooks, book)
		}
	}
	
	return availableBooks

}
func Checker(memberID int) bool {
	if _, ok := newlib.members[memberID]; !ok {
		return false
	}
	return true
}
func ListBorrowedBooks(memberID int) []models.Book {

	borrowed := newlib.members[memberID]
	return borrowed.BorrowedBooks

}
