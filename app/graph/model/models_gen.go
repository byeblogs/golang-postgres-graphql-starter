// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTodoInput struct {
	Title   string `json:"title"`
	Comment string `json:"comment"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type Todo struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Comment   string  `json:"comment"`
	User      *User   `json:"user"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt *string `json:"deletedAt"`
}

type UpdateTodoInput struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
}

type User struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	ImageURL  *string `json:"imageUrl"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt *string `json:"deletedAt"`
}

type UpdatePasswordInput struct {
	OldPassword        string `json:"oldPassword"`
	NewPassword        string `json:"newPassword"`
	NewPasswordConfirm string `json:"newPasswordConfirm"`
}
