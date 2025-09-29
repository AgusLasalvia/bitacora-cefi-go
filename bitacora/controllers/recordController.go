package controllers

import (
	"bitacora/core"
	"bitacora/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordController struct{}

func NewRecordController() *RecordController {
	return &RecordController{}
}

func (rc *RecordController) AddRecord(c *gin.Context) {
	var incomingRecord core.Record

	if err := c.ShouldBindJSON(&incomingRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos Invalidos"})
		return
	}

	err := models.AddRecord(incomingRecord)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear registro"})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (rc *RecordController) GetRecordByID(c *gin.Context) {
	record, err := models.GetRecordByID(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NO se econtraron registros para ese equipo"})
		return
	}
	c.JSON(http.StatusOK, record)
}

func (rc *RecordController) GetRecordByMachine(c *gin.Context) {
	records, err := models.GetRecordByMachine(c.Query("machine"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NO se econtraron registros para ese equipo"})
		return
	}
	c.JSON(http.StatusOK, records)

}
