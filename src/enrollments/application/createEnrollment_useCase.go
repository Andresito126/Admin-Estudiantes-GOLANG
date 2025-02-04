package application

import (
	"github.com/Andresito126/go-estudiantes/src/enrollments/domain/entities"
	"github.com/Andresito126/go-estudiantes/src/enrollments/domain/ports"
)

type CreateEnrollmentUseCase struct {
	enrollmentRepository ports.IEnrollmentRepository
}

func NewCreateEnrollmentUseCase(repo ports.IEnrollmentRepository) *CreateEnrollmentUseCase{

	return &CreateEnrollmentUseCase{enrollmentRepository: repo}
}

func (useCase *CreateEnrollmentUseCase) Run(enrollment *entities.Enrollment) error{

	return useCase.enrollmentRepository.CreateEnrollment(enrollment)

}

