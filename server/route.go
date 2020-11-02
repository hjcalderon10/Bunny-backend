package server

import (
	"net/http"
	"time"

	controllers "github.com/hjcalderon10/bunny-backend/controller"
	models "github.com/hjcalderon10/bunny-backend/model"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	"github.com/labstack/echo"
)

const BASE_PATH = "api/v1"

func SetupRoutes(server *echo.Echo) {
	baseRout := server.Group(BASE_PATH)

	baseRout.GET("/health-check", Get)

	userCtrl := controllers.StartUser()
	userBase := baseRout.Group(userCtrl.BasePath)
	userBase.GET("", userCtrl.GetAllUsers)
	userBase.POST("", userCtrl.CreateUser)
	userBase.GET(userCtrl.UserPath, userCtrl.ReadUser)
	userBase.PUT(userCtrl.UserPath, userCtrl.UpdateUser)
	userBase.DELETE(userCtrl.UserPath, userCtrl.DeleteUser)
	userBase.POST(userCtrl.UserTaskPath, userCtrl.CreateUserTask)

	taskCtrl := controllers.StartTask()
	taskBase := baseRout.Group(taskCtrl.BasePath)
	taskBase.GET("", taskCtrl.GetAllTasks)
	taskBase.GET(taskCtrl.TaskPath, taskCtrl.ReadTask)
	taskBase.PUT(taskCtrl.TaskPath, taskCtrl.UpdateTask)
	taskBase.DELETE(taskCtrl.TaskPath, taskCtrl.DeleteTask)
	taskBase.GET(taskCtrl.StatesPath, taskCtrl.GetAllTaskStates)

}

func Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &models.HealthCheck{
		Name:    settings.Commons.ProjectName,
		Version: settings.Commons.ProjectVersion,
		Date:    time.Now().UTC(),
	})
}
