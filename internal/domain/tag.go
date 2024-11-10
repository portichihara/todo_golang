package domain

type Tag struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name" gorm:"unique"`
    Todos []Todo `json:"-" gorm:"many2many:todo_tags;"`
}
