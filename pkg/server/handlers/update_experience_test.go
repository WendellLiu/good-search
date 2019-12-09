package handlers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wendellliu/good-search/pkg/dto"
	pb "github.com/wendellliu/good-search/pkg/pb"
)

type mockRepo struct {
	dto.Repository
}

func (m *mockRepo) mockGetExperience(ctx context.Context, id string) (dto.Experience, error) {
	return dto.Experience{}, nil
}

func TestUpdateExperience(t *testing.T) {
	tests := []struct {
		description string
		wantRep     *pb.UpdateExperienceResp
		wantErr     bool
	}{
		{
			description: "run success",
			wantRep:     &pb.UpdateExperienceResp{Status: pb.Status_SUCCESS},
			wantErr:     false,
		},
	}

	repo := mockRepo{}
	server := Server{Repository: repo}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {

			switch tt.wantErr {
			case true:
				assert.NotNil(t, gotErr)
			case false:
				assert.Nil(t, gotErr)
			}
		})
	}
}
