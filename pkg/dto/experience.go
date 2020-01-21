package dto

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wendellliu/good-search/pkg/common/dbAdapter"
	"github.com/wendellliu/good-search/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	SalaryTypeYear  = "year"
	SalaryTypeMonth = "month"
	SalaryTypeDay   = "day"
	SalaryTypeHour  = "hour"

	InterviewExperienceType = "interview"
	WorkExperienceType      = "work"
)

type ExpCompany struct {
	Name string `bson:"name" json:"name"`
}

type ExpSection struct {
	ID       int64   `bson:"id" json:"id"`
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle"`
	Content  string  `bson:"content" json:"content"`
}

type InterviewTime struct {
	Year  int64 `bson:"year" json:"year"`
	Month int64 `bson:"month" json:"month"`
}

type InterviewQA struct {
	Question string `bson:"question" json:"question"`
	Answer   string `bson:"answer" json:"answer"`
}

type Salary struct {
	Type   string `bson:"type" json:"type"`
	Amount int64  `bson:"amount" json:"amount"`
}

type Archive struct {
	IsArchived bool   `bson:"is_archived" json:"is_archived"`
	Reason     string `bson:"reason" json:"reason"`
}

type YearMonth struct {
	Year  int64 `bson:"year" json:"year"`
	Month int64 `bson:"month" json:"month"`
}

type Experience struct {
	// common
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Type        string             `bson:"type,omitempty" json:"type"`
	AuthorID    primitive.ObjectID `bson:"author_id,omitempty" json:"author_id"`
	Compnay     ExpCompany         `bson:"company,omitempty" json:"company"`
	LikeCount   int64              `bson:"like_count,omitempty" json:"like_count"`
	ReplyCount  int64              `bson:"reply_count,omitempty" json:"reply_count"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at"`
	Region      *string            `bson:"region,omitempty" json:"region"`
	JobTitle    *string            `bson:"job_title,omitempty" json:"job_title"`
	Title       *string            `bson:"title,omitempty" json:"title"`
	Education   *string            `bson:"education,omitempty" json:"education"`
	Status      *string            `bson:"status,omitempty" json:"status"`
	ReportCount int64              `bson:"report_count,omitempty" json:"report_count"`
	Sections    []ExpSection       `bson:"sections,omitempty" json:"sections"`
	Salary      *Salary            `bson:"salary,omitempty" json:"salary"`
	Archive     *Archive           `bson:"archive,omitempty" json:"archive"`
	UID         *string            `bson:"id,omitempty" json:"id"`

	// interview
	OverallRating               *int64         `bson:"overall_rating,omitempty" json:"overall_rating"`
	InterviewResult             *string        `bson:"interview_result,omitempty" json:"interview_result"`
	InterviewTime               *InterviewTime `bson:"interview_time,omitempty" json:"interview"`
	InterviewQAs                []InterviewQA  `bson:"interview_qas,omitempty" json:"interview_qas"`
	InterviewSensitiveQuestions []string       `bson:"interview_sensitive_questions,omitempty" json:"interview_sensitive_questions"`

	// work
	WeekWorkTime      float64    `bson:"week_work_time,omitempty" json:"week_work_time"`
	RecommendToOthers *string    `bson:"recommend_to_others,omitempty" json:"recommend_to_others"`
	DataTime          *YearMonth `bson:"data_time,omitempty" json:"data_time"`
}

type ExperiencesParams struct {
	Type *string `bson:"type,omitempty" json:"type,omitempty"`
}

type ExperienceDTO interface {
	GetExperience(ctx context.Context, id string) (Experience, error)
	GetExperiences(ctx context.Context, params *ExperiencesParams, opts dbAdapter.Options) ([]Experience, error)
	GetExperiencesCount(ctx context.Context) (int64, error)
}

func (repo Repository) GetExperience(ctx context.Context, id string) (Experience, error) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "dto-experience-GetExperience"},
	)
	collectionName := "experiences"
	result := Experience{}
	var err error

	collection := repo.DB.UseTable(collectionName)

	err = collection.QueryOne(
		context.Background(),
		id,
		&result,
	)

	if err != nil {
		localLogger.Error(err)
	}

	return result, err

}

func (repo Repository) GetExperiences(
	ctx context.Context,
	params *ExperiencesParams,
	opts dbAdapter.Options,
) ([]Experience, error) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "dto-experience-GetExperiences"},
	)
	collectionName := "experiences"
	results := []Experience{}

	query := make(map[string]interface{})
	if params.Type != nil {
		query["type"] = params.Type
	}

	collection := repo.DB.UseTable(collectionName)
	err := collection.QueryPagination(
		ctx,
		query,
		opts,
		&results,
	)

	if err != nil {
		localLogger.Error(err)
	}

	return results, err
}

func (repo Repository) GetExperiencesCount(ctx context.Context) (count int64, err error) {
	localLogger := logger.Logger.WithFields(
		logrus.Fields{"endpoint": "dto-experience-GetExperiencesCount"},
	)
	collectionName := "experiences"
	collection := repo.DB.UseTable(collectionName)

	count, err = collection.AllCount(ctx)

	if err != nil {
		localLogger.Error(err)
	}
	return count, err
}
