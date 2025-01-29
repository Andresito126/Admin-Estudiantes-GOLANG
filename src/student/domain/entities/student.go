package entities


type Student struct {
    ID 			int `json:"id"`
    Name      	string `json:"name"`
    Email     	string `json:"email"`
    Career    	string `json:"career"`
    Matricula 	string `json:"matricula"`
}

func NewStudent(id int, name string, email string, career string, matricula string) *Student {
    return &Student{ID: id, Name: name, Email: email, Career: career, Matricula: matricula}
}

func (s *Student) GetName() string {
    return s.Name
}

func (s *Student) SetName(name string) {
    s.Name = name
}