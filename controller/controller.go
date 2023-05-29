package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizalherniawan/backend-test-psn/dto/request"
	"github.com/rizalherniawan/backend-test-psn/dto/response"
	"github.com/rizalherniawan/backend-test-psn/exception"
	"github.com/rizalherniawan/backend-test-psn/service/score"
	student "github.com/rizalherniawan/backend-test-psn/service/student"
)

type ApiController struct {
	studentService student.StudentService
	scoreService   score.ScoreService
}

func StudentController(r *gin.RouterGroup, studentService student.StudentService, scoreService score.ScoreService) {
	controllerHandler := ApiController{
		studentService: studentService,
		scoreService:   scoreService,
	}
	r.POST("/create", controllerHandler.Create)
	r.POST("/create/:studentId", controllerHandler.PostSubjectAndScore)
	r.GET("/:subject", controllerHandler.GetStudentsBySubject)
	r.GET("/:subject/:studentId", controllerHandler.GetByStudentIdAndSubject)
	r.PUT("/:subject/:studentId", controllerHandler.UpdateScoreStudent)
	r.DELETE("/:studentId", controllerHandler.Delete)
}

func (controller *ApiController) PostSubjectAndScore(ctx *gin.Context) {
	studentId := ctx.Param("studentId")

	createStudentScoreRequest := request.ScoreRequest{}
	err := ctx.ShouldBindJSON(&createStudentScoreRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serviceErr := controller.scoreService.Create(studentId, createStudentScoreRequest)

	if serviceErr != nil {
		exception.NotFoundException(ctx, serviceErr.Error())
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ApiController) Create(ctx *gin.Context) {
	createStudentRequest := request.StudentRequest{}
	err := ctx.ShouldBindJSON(&createStudentRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.studentService.Create(createStudentRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ApiController) GetStudentsBySubject(ctx *gin.Context) {
	subject := ctx.Param("subject")
	result, err := controller.studentService.FindBySubject(subject)

	if err != nil {
		exception.NotFoundException(ctx, err.Error())
		return
	}

	var studentResponses []response.StudentResponse

	for _, element := range result {
		student := response.StudentResponse{}
		student.Id = element.Id
		student.Age = element.Age
		student.Name = element.Name
		for _, e := range element.StudentScores {
			score := response.ScoreResponse{}
			score.Subject = e.Subjects
			score.Score = e.Score
			student.Subject = append(student.Subject, score)
		}
		studentResponses = append(studentResponses, student)
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   studentResponses,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *ApiController) GetByStudentIdAndSubject(ctx *gin.Context) {
	subject := ctx.Param("subject")
	studentId := ctx.Param("studentId")
	result, err := controller.studentService.FindBySubjectAndStudentId(subject, studentId)

	if err != nil {
		exception.NotFoundException(ctx, err.Error())
		return
	}

	var studentResponse response.StudentResponse

	student := response.StudentResponse{}
	student.Id = result.Id
	student.Age = result.Age
	student.Name = result.Name
	for _, e := range result.StudentScores {
		score := response.ScoreResponse{}
		score.Subject = e.Subjects
		score.Score = e.Score
		student.Subject = append(student.Subject, score)
	}

	studentResponse = student

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   studentResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ApiController) UpdateScoreStudent(ctx *gin.Context) {

	subject := ctx.Param("subject")
	studentId := ctx.Param("studentId")

	createUpdateStudentScoreRequest := request.ScoreRequest{}
	err := ctx.ShouldBindJSON(&createUpdateStudentScoreRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errService := controller.scoreService.UpdateScoreBySubjectAndStucentId(subject, studentId, createUpdateStudentScoreRequest)

	if errService != nil {
		exception.NotFoundException(ctx, errService.Error())
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ApiController) Delete(ctx *gin.Context) {
	studentId := ctx.Param("studentId")

	err := controller.studentService.DeleteStudentById(studentId)

	if err != nil {
		exception.NotFoundException(ctx, err.Error())
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
