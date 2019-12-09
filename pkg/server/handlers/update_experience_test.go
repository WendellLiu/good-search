package handlers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wendellliu/good-search/pkg/dto"
	"github.com/wendellliu/good-search/pkg/logger"
	pb "github.com/wendellliu/good-search/pkg/pb"
)

type mockRepo struct {
	dto.Repository
}

func (m mockRepo) GetExperience(ctx context.Context, id string) (dto.Experience, error) {
	return dto.Experience{}, nil
}

func TestUpdateExperience(t *testing.T) {
	tests := []struct {
		description string
		paramID     string
		wantResp    *pb.UpdateExperienceResp
		wantErr     bool
	}{
		{
			description: "run success",
			paramID:     "123",
			wantResp:    &pb.UpdateExperienceResp{Status: pb.Status_SUCCESS},
			wantErr:     false,
		},
	}

	logger.Load()
	repo := mockRepo{}
	handlers := Server{Repository: repo}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
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
