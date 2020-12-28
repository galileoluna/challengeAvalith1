package database

import (
	"fmt"

	"github.com/galileoluna/challengeAvalith1/models"
)

func InsertAlumno(nuevoAlumno models.Alumno) (int, error) {
	//Create
	var alumnoID int
	err := db.QueryRow(`INSERT INTO ALUMNO(NOMBRE) VALUES($1) RETURNING IDALUMNO`, nuevoAlumno.nombre).Scan(&alumnoID)

	if err != nil {
		return 0, err
	}
	fmt.Printf("Last inserted ID: %v\n", alumnoID)
	return alumnoID, err
}

func GetAlumnos() ([]models.Alumno, error) {
	alumnos := []models.Alumno{}
	rows, err := db.Query(`SELECT IDALUMNO,NOMBRE FROM ALUMNO`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var IDALUMNO int
		var NOMBRE string
		var alumnoActual Alumno

		err = rows.Scan(&IDALUMNO, &NOMBRE)
		if err != nil {
			return alumnos, err
		}
		alumnoActual = models.Alumno{nombre: NOMBRE}
		alumnos = append(alumnos, alumnoActual)
	}
	return alumnos, err
}
