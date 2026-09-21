package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrBookNotAvailable  = errors.New("book is currently not available for borrowing")
	ErrLoanAlreadyClosed = errors.New("this loan has already been returned")
	ErrBookMismatch      = errors.New("book does not match the loan record")
)

type Book struct {
	ID          int
	Title       string
	Author      string
	IsAvailable bool
}

type Member struct {
	ID   int
	Name string
}

type Loan struct {
	ID         int
	Book       *Book
	Member     *Member
	BorrowDate time.Time
	DueDate    time.Time
	ReturnDate time.Time
	IsActive   bool
}

func NewBook(id int, title, author string) *Book {
	return &Book{
		ID:          id,
		Title:       title,
		Author:      author,
		IsAvailable: true,
	}
}

func NewMember(id int, name string) *Member {
	return &Member{
		ID:   id,
		Name: name,
	}
}

func BorrowBook(loanID int, book *Book, member *Member, durationDays int) (*Loan, error) {
	if !book.IsAvailable {
		return nil, ErrBookNotAvailable
	}

	book.IsAvailable = false
	now := time.Now()

	loan := &Loan{
		ID:         loanID,
		Book:       book,
		Member:     member,
		BorrowDate: now,
		DueDate:    now.AddDate(0, 0, durationDays),
		IsActive:   true,
	}

	return loan, nil
}

func (l *Loan) ReturnBook() error {
	if !l.IsActive {
		return ErrLoanAlreadyClosed
	}

	l.Book.IsAvailable = true
	l.ReturnDate = time.Now()
	l.IsActive = false

	return nil
}

func main() {
	book1 := NewBook(101, "The Go Programming Language", "Alan A. A. Donovan")
	book2 := NewBook(102, "Clean Code", "Robert C. Martin")

	member1 := NewMember(1, "Ahmed")
	member2 := NewMember(2, "Sara")

	fmt.Println("--- Attempting to borrow Book 1 ---")
	loan1, err := BorrowBook(1, book1, member1, 14)
	if err != nil {
		fmt.Println("Error borrowing:", err)
	} else {
		fmt.Printf("Success! '%s' borrowed by %s. Due Date: %s\n",
			loan1.Book.Title, loan1.Member.Name, loan1.DueDate.Format("2006-01-02"))
	}

	fmt.Println("\n--- Attempting to borrow Book 1 again (Unavailable) ---")
	_, err = BorrowBook(2, book1, member2, 7)
	if err != nil {
		fmt.Println("Error borrowing:", err)
	}

	fmt.Println("\n--- Borrowing Book 2 ---")
	loan2, err := BorrowBook(3, book2, member2, 10)
	if err == nil {
		fmt.Printf("Success! '%s' borrowed by %s. Due Date: %s\n",
			loan2.Book.Title, loan2.Member.Name, loan2.DueDate.Format("2006-01-02"))
	}

	fmt.Println("\n--- Returning Book 1 ---")
	err = loan1.ReturnBook()
	if err != nil {
		fmt.Println("Error returning:", err)
	} else {
		fmt.Printf("Book '%s' successfully returned on %s!\n",
			loan1.Book.Title, loan1.ReturnDate.Format("2006-01-02 15:04:05"))
	}

	fmt.Println("\n--- Attempting to borrow Book 1 after return ---")
	loan3, err := BorrowBook(4, book1, member2, 7)
	if err != nil {
		fmt.Println("Error borrowing:", err)
	} else {
		fmt.Printf("Success! '%s' borrowed by %s. Due Date: %s\n",
			loan3.Book.Title, loan3.Member.Name, loan3.DueDate.Format("2006-01-02"))
	}
}
