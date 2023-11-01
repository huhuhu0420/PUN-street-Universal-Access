package delivery

import (
	"strconv"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DiscountHandler struct {
	DiscountUsecase domain.DiscountUsecase
}

func NewDiscountHandler(e *gin.Engine, discountUsecase domain.DiscountUsecase) {
	handler := &DiscountHandler{
		DiscountUsecase: discountUsecase,
	}
	e.GET("/api/v1/store/:storeID/discounts", handler.GetAllDiscountByStoreId)
}

func (s *DiscountHandler) GetAllDiscountByStoreId(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	discounts, err := s.DiscountUsecase.GetByStoreID(c, storeID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}
	c.JSON(200, discounts)

}
