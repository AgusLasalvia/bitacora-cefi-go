package controllers

import (
	"github.com/gin-gonic/gin"
)

type RecordController struct{}

func NewRecordController() *RecordController {
	return &RecordController{}
}

func (rc *RecordController) AddRecord(c *gin.Context) {}

func (rc *RecordController) GetRecordByID(c *gin.Context) {}

func (rc *RecordController) GetRecordByMachine(c *gin.Context) {}
