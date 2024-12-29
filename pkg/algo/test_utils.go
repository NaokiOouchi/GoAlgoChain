package algo

// Item is a common type used for testing.
type Item struct {
	ID     int
	Name   string
	Active bool
}

// User is another common type used for testing.
type User struct {
	ID     int
	Name   string
	Active bool
}

// Order is another common type used for testing.
type Order struct {
	OrderID int
	UserID  int
	Item    string
}
