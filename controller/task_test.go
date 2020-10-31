package controller

/*
import (
	"fmt"
	"net/http"
	"testing"

	mocks "github.com/hjcalderon10/bunny-backend/test/mocks/service"
	testUtils "github.com/hjcalderon10/bunny-backend/test/util"
	"github.com/stretchr/testify/assert"
)

func TestReadTask(t *testing.T) {
	taskID := 1
	endPoint := fmt.Sprintf("/api/tasks/%d", taskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("ReadTask").Return()

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.ReadTask(c)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}
func TestUpdateTask(t *testing.T) {
	taskID := 1
	endPoint := fmt.Sprintf("/api/tasks/%d", taskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("UpdateTask").Return()

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.UpdateTask(c)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}
func TestDeleteTask(t *testing.T) {
	taskID := 1
	endPoint := fmt.Sprintf("/api/tasks/%d", taskID)

	_, _, rec, c := testUtils.SetupServerTest(
		http.MethodGet,
		endPoint,
		nil,
	)

	taskSrvMock := mocks.TaskMock{}
	taskSrvMock.On("DeleteTask").Return()

	ctrl := StartTask()
	ctrl.service = &taskSrvMock
	ctrl.DeleteTask(c)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}
*/
