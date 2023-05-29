package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/rizalherniawan/backend-test-psn/controller"
	scoreRepo "github.com/rizalherniawan/backend-test-psn/repository/score"
	studentRepo "github.com/rizalherniawan/backend-test-psn/repository/student"
	scoreService "github.com/rizalherniawan/backend-test-psn/service/score"
	studentService "github.com/rizalherniawan/backend-test-psn/service/student"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB) {
	r := gin.New()

	var err = godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	studentRepo := studentRepo.NewStudentRepositoryImpl(db)

	studentService := studentService.NewStudentServiceImpl(studentRepo)

	scoreRepo := scoreRepo.NewScoreRepositoryImpl(db)

	scoreService := scoreService.NewScoreServiceImpl(scoreRepo, studentRepo)

	api := r.Group("/api")

	controller.StudentController(api.Group("/student"), studentService, scoreService)

	r.Run(fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")))
}
