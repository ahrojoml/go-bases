/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:


Name: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]


Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.

Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

package main

import (
	"fmt"
	"time"
)

type Student struct {
	name     string
	lastName string
	dni      string
	joinDate time.Time
}

func (s *Student) PrintStudent() {
	fmt.Printf("%s %s, %s, %s", s.name, s.lastName, s.dni, s.joinDate.Format("01-02-2006"))
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) SetLastName(lastName string) {
	s.lastName = lastName
}

func (s *Student) SetDni(dni string) {
	s.dni = dni
}

func (s *Student) SetJoinDate(day, month, year int) {
	timeMonth := time.Month(month)
	date := time.Date(year, timeMonth, day, 0, 0, 0, 0, time.Local)
	s.joinDate = date
}

func main() {
	var student Student = Student{}

	var input string
	var day, month, year int

	fmt.Printf("Type Name: ")
	_, _ = fmt.Scan(&input)
	student.SetName(input)

	fmt.Printf("Type Last Name: ")
	_, _ = fmt.Scan(&input)
	student.SetLastName(input)

	fmt.Printf("Type DNI: ")
	_, _ = fmt.Scan(&input)
	student.SetDni(input)

	fmt.Printf("Join Date day/month/year: ")
	_, _ = fmt.Scanf("%d/%d/%d\n", &day, &month, &year)
	student.SetJoinDate(day, month, year)

	student.PrintStudent()
}
