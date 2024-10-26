package location

import "errors"

var users = make([]User, 0)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Room     string `json:"room"`
}

func GetUser(id string) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}

	return User{}, errors.New("user not found")
}

func GetRoomUsers(room string) []User {
	total := make([]User, 0)

	for _, u := range users {
		if u.Room == room {
			total = append(total, u)
		}
	}

	return total
}

func RemoveUser(id string) {
	for i, u := range users {
		if u.Id == id {
			users = remove(users, i)
			break
		}
	}
}

func remove(s []User, i int) []User {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]

	return s[:len(s)-1]
}
