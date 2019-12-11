package handlers

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
)

type mockRepo struct {
	dto.Repository
	mux        *sync.Mutex
	experience dto.Experience
}

func (m *mockRepo) setExperience(e dto.Experience) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.experience = e
	fmt.Printf("set experience %+v \n", m.experience)
}

func (m mockRepo) GetExperience(ctx context.Context, id string) (dto.Experience, error) {
	fmt.Printf("mock get experience %+v\n", m.experience)
	return m.experience, nil
}

func TestUpdateExperience(t *testing.T) {
	logger.Load()
	repo := &mockRepo{mux: &sync.Mutex{}, experience: dto.Experience{Type: "foo"}}
	handlers := Server{Repository: repo}

	tests := []struct {
		description string
		paramID     string
		wantResp    *pb.UpdateExperienceResp
		wantErr     bool
		setup       func()
	}{
		{
			description: "run success",
			paramID:     "123",
			wantResp: &pb.UpdateExperienceResp{
				Status: pb.Status_SUCCESS,
				Experience: &pb.ExperiencePayload{
					Type: "work",
				},
			},
			wantErr: false,
			setup: func() {
				repo.setExperience(dto.Experience{
					Type: "work",
				})
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			req := &pb.UpdateExperienceReq{
				Id: tt.paramID,
			}
			resp, gotErr := handlers.UpdateExperience(context.Background(), req)

			assert.EqualValues(t, tt.wantResp, resp)
			switch tt.wantErr {
			case true:
				assert.NotNil(t, gotErr)
			case false:
				assert.Nil(t, gotErr)
			}
		})
	}
}
