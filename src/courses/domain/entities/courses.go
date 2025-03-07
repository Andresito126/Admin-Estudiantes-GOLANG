package entities

type Course struct {
    ID             int    `json:"id"`
    Name           string `json:"name"`
    Duration       int    `json:"duration"`
    AvailableSlots int    `json:"available_slots"`
}

func NewCourse(id int, name string, duration int, availableSlots int) *Course {
    return &Course{ID: id, Name: name, Duration: duration, AvailableSlots: availableSlots}
}

func (c *Course) GetName() string {
    return c.Name
}

func (c *Course) SetName(name string) { 
    c.Name = name
}

func (c *Course) GetAvailableSlots() int {
    return c.AvailableSlots
}

func (c *Course) SetAvailableSlots(slots int) {
    c.AvailableSlots = slots
}
