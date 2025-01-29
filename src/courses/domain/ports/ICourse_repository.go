package ports

import (

)

type CourseRepository interface {
	Save()
	FindAll()
	Delete()
	Edit()
}