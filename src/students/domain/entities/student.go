package entities


type Student struct {
    ID 			int `json:"id"`
    Name      	string `json:"name"`
    Email     	string `json:"email"`
    Career    	string `json:"career"`
    Matricula 	string `json:"matricula"`
}