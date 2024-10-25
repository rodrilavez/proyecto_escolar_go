package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"sistema_escolar/database"
	"sistema_escolar/handlers"
)

var db *sql.DB

func main() {
	var err error
	db, err = database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.Default())

	// Rutas para Estudiantes
	router.POST("/api/students", handlers.CreateStudent)               // Crear un nuevo estudiante
	router.DELETE("/api/students/:student_id", handlers.DeleteStudent) // Eliminar un estudiante por ID
	router.PUT("/api/students/:student_id", handlers.UpdateStudent)    // Actualizar información de un estudiante por ID
	router.GET("/api/students", handlers.GetAllStudents)               // Obtener la lista de todos los estudiantes
	router.GET("/api/students/:student_id", handlers.GetStudent)       // Obtener la información de un estudiante por ID

	// Rutas para Materias
	router.POST("/api/subjects", handlers.CreateSubject)               // Crear una nueva materia
	router.PUT("/api/subjects/:subject_id", handlers.UpdateSubject)    // Actualizar una materia por ID
	router.GET("/api/subjects/:subject_id", handlers.GetSubject)       // Obtener información de una materia por ID
	router.GET("/api/subjects", handlers.GetAllSubjects)               // Obtener la lista de todas las materias
	router.DELETE("/api/subjects/:subject_id", handlers.DeleteSubject) // Eliminar una materia por ID

	// Rutas para Calificaciones
	router.POST("/api/grades", handlers.CreateGrade)                              // Crear una nueva calificación
	router.PUT("/api/grades/:grade_id", handlers.UpdateGrade)                     // Actualizar una calificación por ID
	router.DELETE("/api/grades/:grade_id", handlers.DeleteGrade)                  // Eliminar una calificación por ID
	router.GET("/api/grades/:grade_id", handlers.GetGrade)                        // Obtener información de una calificación por ID
	router.GET("/api/grades/student/:student_id", handlers.GetAllGradesByStudent) // Obtener todas las calificaciones de un estudiante por ID

	router.Run(":8080")
}
