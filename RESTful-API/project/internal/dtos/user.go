package dtos

// UserDTO represents the structure of a user in DTO form
type UserDTO struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

