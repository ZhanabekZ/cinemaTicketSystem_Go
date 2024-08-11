package main

import (
	"errors"
	"fmt"
)

type Movie struct {
	ID   int
	Name string
}

type User struct {
	ID   int
	Name string
}

type Ticket struct {
	ID      int
	UserID  int
	MovieID int
}

type CinemaTicketSystem struct {
	Movies          []Movie
	Users           []User
	Tickets         []Ticket
	movieIDCounter  int
	userIDCounter   int
	ticketIDCounter int
}

func NewCinemaTicketSystem() *CinemaTicketSystem {
	return &CinemaTicketSystem{}
}

func (cts *CinemaTicketSystem) AddMovie(name string) int {
	cts.movieIDCounter++
	movie := Movie{ID: cts.movieIDCounter, Name: name}
	cts.Movies = append(cts.Movies, movie)
	return movie.ID
}

func (cts *CinemaTicketSystem) ShowAllMovies() {
	for _, movie := range cts.Movies {
		fmt.Printf("%d. %s\n", movie.ID, movie.Name)
	}
}

func (cts *CinemaTicketSystem) AddUser(name string) int {
	cts.userIDCounter++
	user := User{ID: cts.userIDCounter, Name: name}
	cts.Users = append(cts.Users, user)
	return user.ID
}

func (cts *CinemaTicketSystem) BuyTicket(userID int, movieID int) (int, error) {
	userExists := false
	movieExists := false

	for _, user := range cts.Users {
		if user.ID == userID {
			userExists = true
			break
		}
	}

	for _, movie := range cts.Movies {
		if movie.ID == movieID {
			movieExists = true
			break
		}
	}

	if !userExists || !movieExists {
		return 0, errors.New("user or movie not found")
	}

	cts.ticketIDCounter++
	ticket := Ticket{ID: cts.ticketIDCounter, UserID: userID, MovieID: movieID}
	cts.Tickets = append(cts.Tickets, ticket)
	return ticket.ID, nil
}

func (cts *CinemaTicketSystem) CancelTicket(ticketID int) bool {
	for i, ticket := range cts.Tickets {
		if ticket.ID == ticketID {
			cts.Tickets = append(cts.Tickets[:i], cts.Tickets[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	cts := NewCinemaTicketSystem()

	for {
		fmt.Println("\nДобро пожаловать в систему управления билетами в кинотеатр!")
		fmt.Println("Выберите действие:")
		fmt.Println("1. Добавить новый фильм")
		fmt.Println("2. Показать все доступные фильмы")
		fmt.Println("3. Добавить нового пользователя")
		fmt.Println("4. Купить билет")
		fmt.Println("5. Отменить покупку билета")
		fmt.Println("6. Выйти")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var movieName string
			fmt.Println("Введите название фильма:")
			fmt.Scanln(&movieName)
			movieID := cts.AddMovie(movieName)
			fmt.Printf("Фильм добавлен с ID %d\n", movieID)

		case 2:
			fmt.Println("Доступные фильмы:")
			cts.ShowAllMovies()

		case 3:
			var userName string
			fmt.Println("Введите имя пользователя:")
			fmt.Scanln(&userName)
			userID := cts.AddUser(userName)
			fmt.Printf("Пользователь добавлен с ID %d\n", userID)

		case 4:
			var userID, movieID int
			fmt.Println("Введите ID пользователя:")
			fmt.Scanln(&userID)
			fmt.Println("Введите ID фильма:")
			fmt.Scanln(&movieID)
			ticketID, err := cts.BuyTicket(userID, movieID)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Билет куплен с ID %d\n", ticketID)
			}

		case 5:
			var ticketID int
			fmt.Println("Введите ID билета для отмены:")
			fmt.Scanln(&ticketID)
			if cts.CancelTicket(ticketID) {
				fmt.Println("Билет успешно отменен")
			} else {
				fmt.Println("Билет с таким ID не найден")
			}

		case 6:
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Неверный выбор")
		}
	}
}
