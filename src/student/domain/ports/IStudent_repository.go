package ports

import (

)

type StudentRepository interface {
	Save()
	FindAll()
	Delete()
	Edit()
}