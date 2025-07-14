package person

type Person struct {
	Id    *int   `json:"id,omitempty" gorm:"primaryKey"`
	Name  string `json:"name" validate:"required,min=3,max=25"`
	Email string `json:"email" validate:"required,email"`
}
