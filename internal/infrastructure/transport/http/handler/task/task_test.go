package task

import (
	"bytes"
	"dddapib/internal/domain/model/aperr"
	"dddapib/internal/domain/model/entity"
	"dddapib/internal/domain/service"
	mocktask "dddapib/internal/domain/service/task/mock"
	"dddapib/internal/infrastructure/transport/http/dto"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type endpointTest struct {
	name             string
	method           string
	path             string
	input            any
	expectedResponse any
	expectedStatus   int
	errorCode        string
	setup            func(tc *endpointTest)
}

// test the spec of task API endpoint
func TestUpdateTask(t *testing.T) {

	ctl := gomock.NewController(t)
	mockTaskSvc := mocktask.NewMockService(ctl)

	router := gin.Default()

	Init(router, &service.Service{
		TaskService: mockTaskSvc,
	})

	testCases := []endpointTest{
		{
			name:   "list-task",
			method: "GET",
			path:   "/tasks",
			expectedResponse: &dto.ApiResponse{
				Code: "",
				Data: []*entity.Task{{"1", "1", 0}},
			},
			expectedStatus: 200,
			setup: func(_ *endpointTest) {
				mockTaskSvc.EXPECT().
					ListTasks().
					Return([]*entity.Task{{"1", "1", 0}}, nil)
			},
		},
		{
			name:   "close-task-ok",
			method: "PUT",
			path:   "/tasks/task-001",
			expectedResponse: &dto.ApiResponse{
				Code: "",
				Data: &entity.Task{"task-001", "unit-test", 1},
			},
			expectedStatus: 200,
			setup: func(_ *endpointTest) {
				mockTaskSvc.EXPECT().
					CloseTask("task-001").
					Return(&entity.Task{"task-001", "unit-test", 1}, nil)
			},
		},
		{
			name:           "close-task-not-found",
			method:         "PUT",
			path:           "/tasks/task-001",
			expectedStatus: 400,
			setup: func(_ *endpointTest) {
				mockTaskSvc.EXPECT().
					CloseTask("task-001").
					Return(nil, aperr.TaskNotFound)
			},
			errorCode: aperr.TaskNotFound.Code,
		},
		{
			name:             "delete-task",
			method:           "DELETE",
			path:             "/tasks/task-001",
			expectedResponse: &dto.ApiResponse{},
			expectedStatus:   200,
			setup: func(_ *endpointTest) {
				mockTaskSvc.EXPECT().
					DeleteTask("task-001").
					Return(nil)
			},
		},
		{
			name:   "create-task-ok",
			method: "POST",
			path:   "/tasks",
			input:  &dto.CreateTaskRequest{Name: "unit-test"},
			expectedResponse: &dto.ApiResponse{
				Data: &entity.Task{"task-001", "unit-test", 0},
			},
			expectedStatus: 200,
			setup: func(_ *endpointTest) {
				mockTaskSvc.EXPECT().
					CreateTask("unit-test").
					Return(&entity.Task{"task-001", "unit-test", 0}, nil)
			},
		},
		{
			name:           "create-task-no-name",
			method:         "POST",
			path:           "/tasks",
			input:          &dto.CreateTaskRequest{},
			expectedStatus: 400,
			errorCode:      aperr.InvalidRequest("").Code,
		},
		{
			name:           "create-task-name-too-long",
			method:         "POST",
			path:           "/tasks",
			expectedStatus: 400,
			errorCode:      aperr.InvalidRequest("").Code,
			setup: func(tc *endpointTest) {
				longName := strings.Repeat("a", 101)
				tc.input = &dto.CreateTaskRequest{Name: longName}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// basic setup
			recorder := httptest.NewRecorder()
			if tc.setup != nil {
				tc.setup(&tc)
			}

			// prepare request
			payloadBuf := &bytes.Buffer{}
			if tc.input != nil {
				payload, err := json.Marshal(tc.input)
				assert.NoError(t, err)
				payloadBuf.Write(payload)
			}
			req, _ := http.NewRequest(tc.method, tc.path, payloadBuf)

			// run test
			router.ServeHTTP(recorder, req)

			// evaluate
			assert.Equal(t, tc.expectedStatus, recorder.Code)
			if tc.expectedResponse != nil {
				assert.Equal(t, toJson(tc.expectedResponse), recorder.Body.String())
			}
			if tc.errorCode != "" {
				apiRsp := unmarshal[dto.ApiResponse](recorder.Body.Bytes())
				assert.Equal(t, tc.errorCode, apiRsp.Code)
			}
		})

	}
}

func toJson(input any) string {
	jsonBytes, _ := json.Marshal(input)
	return string(jsonBytes)
}

func unmarshal[T any](input []byte) *T {
	var output *T
	json.Unmarshal(input, &output)
	return output
}
