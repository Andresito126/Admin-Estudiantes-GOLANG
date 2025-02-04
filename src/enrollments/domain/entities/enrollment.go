package entities

type Enrollment struct{

	ID 				int `json:"id"`
	StudentID 		int `json:"student_id"`
	CourseID 		int `json:"course_id"`
	
}	

func NewEnrollment(id int, studentID int, courseID int, ) *Enrollment {
	return &Enrollment{
		ID:             id,
		StudentID:      studentID,
		CourseID:       courseID,

	}
}