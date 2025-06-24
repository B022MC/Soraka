package score

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ScoreService struct {
	logger *log.Logger
	uc     *scoreBiz.ScoreUseCase
}

func NewScoreService(logger *log.Logger, uc *scoreBiz.ScoreUseCase) *ScoreService {
	return &ScoreService{
		logger: logger,
		uc:     uc,
	}
}

func (s *ScoreService) RegisterRouter(group gin.IRoutes) {
	group.POST("/score/getUserScore", s.GetUserScore)
}

func (s *ScoreService) GetUserScore(c *gin.Context) {
	var req reqScore.GetUserScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	result, err := s.uc.GetUserScore(c.Request.Context(), req)
	if err != nil {
		response.Fail(c, ecode.Failed, err)
		return
	}
	response.Success(c, result)
}
