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

func (s *Student) GetEmail() string {
    return s.Email
}

func (s *Student) SetEmail(email string) {
    s.Email = email
}

func (s *Student) GetCareer() string {
    return s.Career
}

func (s *Student) SetCareer(career string) {
    s.Career = career
}

func (s *Student) GetMatricula() string {
    return s.Matricula
}

func (s *Student) SetMatricula(matricula string) {
    s.Matricula = matricula
}

func (s *Student) GetID() int {
    return s.ID
}   

func (s *Student) SetID(id int) {
    s.ID = id
}   

