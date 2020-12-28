package models

type Alumno struct {
	nombre   string
	apellido string
	dni      int32
}

func NewAlumno(nombre_input string, apellido_input string, dni_input int32) Alumno {
	return Alumno{
		nombre:   nombre_input,
		apellido: apellido_input,
		dni:      dni_input,
	}
}
func GetNombre() string {
	return Alumno.nombre
}
