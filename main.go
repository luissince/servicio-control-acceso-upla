package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar la ubicación horaria "America/Lima"
	time.LoadLocation("America/Lima")

	// Cargar las variables de entorno desde el archivo .env
	godotenv.Load()

	// Obtener el valor de la variable de entorno GO_PORT
	var go_port string = os.Getenv("GO_PORT")

	// Obtener el valor de la variable de entorno RUTA_LOG
	var ruta_log string = os.Getenv("RUTA_LOG")

	// Crear el archivo de registro
	f, err := os.Create(ruta_log)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// Establecer la salida del registro al archivo creado
	log.SetOutput(f)

	// Crear una instancia del enrutador Gin
	router := gin.Default()

	// Utilizar el middleware Logger de Gin
	router.Use(gin.Logger())

	// Middleware para CORS
	router.Use(cors.Default())

	// Middleware para manejar el error 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Ruta no encontrada."})
	})

	// Rutas
	router.GET("/", func(c *gin.Context) {
		log.Println("Endpoint ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "API GO LANG APP MOVIL UPLA",
		})
	})

	// Rutas

	// Define una ruta GET para obtener los datos de un estudiante dado su código
	// Ejemplo de uso: /obtenerdatos/A87560E
	// router.GET("/obtenerdatos/:codigo", controller.ObtenerDatosEstudiante)

	// Define una ruta GET para obtener los datos de su grado académico
	// Ejemplo de uso: /progresoacademico/A87560E
	// router.GET("/progresoacademico/:codigo", controller.ObtenerProgresoAcademico)

	// Define una ruta GET para obtener los datos de su cuenta de Wifi
	// Ejemplo de uso: /obtenerwifi/A87560E
	// router.GET("/obtenerwifi/:codigo", controller.ObtenerWifi)

	// Iniciar el servidor
	router.Run(go_port)
}
