package formatter

import "github.com/ekokurniadi/pos-golang/entity"

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func FormatUser(user entity.User) UserFormatter {
	userFormatter := UserFormatter{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Username = user.Username
	userFormatter.Password = user.Password
	userFormatter.Role = user.Role
	return userFormatter
}
func FormatUsers(users []entity.User) []UserFormatter {
	usersFormatter := []UserFormatter{}
	for _, user := range users {
		userFormatter := FormatUser(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}
	return usersFormatter
}

//Generated by Micagen at 08 Desember 2021
