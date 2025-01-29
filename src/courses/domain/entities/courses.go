package entities

type Course struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Duration int    `json:"duration"` 
}

func NewCourse(id int, name string, duration int) *Course {
    return &Course{ID: id, Name: name, Duration: duration}
}

func (c *Course) GetName() string {
    return c.Name
}

func (c *Course) SetName(name string) { 
    c.Name = name
}
