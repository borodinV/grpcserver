package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"grpc/client/proto"
	"log"
	"os"
)

var option string

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := proto.NewLibraryClient(conn)

	for {
		fmt.Println()
		fmt.Println()
		fmt.Println("Available methods:")
		fmt.Println("Add book to library - Add")
		fmt.Println("Get book from library - Get")
		fmt.Println("Get all book from library - All")
		fmt.Println("Update info about book - Update")
		fmt.Println("Delete book from library - Delete")
		fmt.Println("Search book by name from library - Search")
		fmt.Println()
		fmt.Println("Exit from program - Exit")
		fmt.Println()

		_, err := fmt.Scan(&option)
		if err != nil {
			log.Printf("Error on scan: %v\n", err)
			continue
		}

		switch option {

		case "Add":

			var input proto.Book

			for {
				fmt.Println("Book name: ")
				_, err := fmt.Scan(&input.Name)
				if err != nil {
					log.Println("Error on scan!", err)
					continue
				}
				fmt.Println("Book author: ")
				_, err = fmt.Scan(&input.Author)
				if err != nil {
					log.Println("Error on scan!", err)
					continue
				}
				fmt.Println("Release year of book: ")
				_, err = fmt.Scan(&input.Year)
				if err != nil {
					log.Println("Error on scan!\n", err)
					continue
				}

				fmt.Printf("Your book: \nName: %s | Author: %s | Year: %s\n",
					input.Name, input.Author, input.Year)

				id, err := client.AddBook(ctx, &input)
				if err != nil {
					log.Printf("Error on executing method: %v\n", err)
				}
				fmt.Printf("Book succesfully added!\nYour book ID: %v", id.Id)
				fmt.Println()

				break
			}

		case "Get":

			var bookId proto.BookID

			for {
				fmt.Println("Book ID: ")
				_, err := fmt.Scan(&bookId.Id)
				if err != nil {
					log.Println("Error on scan!", err)
					continue
				}

				book, err := client.GetBook(ctx, &bookId)
				if err != nil {
					log.Printf("Error on executing method: %v\n", err)
				}
				fmt.Println("Your book: ")
				fmt.Println(book)
				break

			}
		case "All":

			books, err := client.GetAll(ctx, &wrappers.StringValue{Value: ""})
			if err != nil {
				log.Printf("Error on executing method: %v\n", err)
			}

			fmt.Println("Searching result: ")

			for _, book := range books.Books {
				fmt.Println(book)
			}

			break

		case "Update":

			var input proto.Book

			for {
				fmt.Println("Book id: ")
				_, err := fmt.Scan(&input.Id)
				if err != nil {
					log.Println("Error on scan!", err)
					continue
				}
				fmt.Println("New book name: ")
				_, err = fmt.Scan(&input.Name)
				if err != nil {
					log.Println("Error on scan!", err)
					continue
				}
				fmt.Println("New book author: ")
				_, err = fmt.Scan(&input.Author)
				if err != nil {
					log.Println("Error on scan!", err)
					continue
				}
				fmt.Println("New release year of book: ")
				_, err = fmt.Scan(&input.Year)
				if err != nil {
					log.Println("Error on scan!\n", err)
					continue
				}
				fmt.Printf("Your new book: \nName: %s | Author: %s | Year: %s\n",
					input.Name, input.Author, input.Year)

				response, err := client.UpdateBook(ctx, &input)
				if err != nil {
					log.Printf("Error on executing method: %v\n", err)
				}
				fmt.Println(response)
				break

			}
		case "Delete":

			var bookId proto.BookID

			for {
				fmt.Println("Book ID: ")
				_, err := fmt.Scan(&bookId.Id)
				if err != nil {
					log.Println("Error on scan!", err)
					continue
				}

				response, err := client.DeleteBook(ctx, &bookId)
				if err != nil {
					log.Printf("Error on executing method: %v\n", err)
				}
				fmt.Println(response)
				break

			}
		case "Search":

			var searchName proto.BookName

			for {
				fmt.Println("Input book name for searching: ")
				_, err := fmt.Scan(&searchName.Name)
				if err != nil {
					log.Println("Error on scan!", err)
					continue
				}

				books, err := client.SearchBookByName(ctx, &searchName)
				if err != nil {
					log.Printf("Error on executing method: %v\n", err)
				}

				fmt.Println("Searching result: ")

				for _, book := range books.Books {
					fmt.Println(book)
				}

				break
			}
		case "Exit":
			os.Exit(1)
		default:
			log.Println("Incorrect input data!")
			continue
		}
	}
}
