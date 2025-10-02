package controllers

import (
	"bitacora/core"
	"bitacora/models"
	"fmt"
	"net/http"
	"time"

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
		fmt.Print(err.Error())
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (rc *RecordController) GetRecordByID(c *gin.Context) {

	id := c.Query("id")
	fmt.Print(id)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se ha ingresado ningun ID para buscar"})
		return
	}

	record, err := models.GetRecordByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NO se econtraron registros para ese equipo"})
		return
	}

	start := time.UnixMilli(record.StartDateTime).Format("02/01/2006 15:04")

	var end string
	if record.EndDateTime != 0 {
		end = time.UnixMilli(record.EndDateTime).Format("02/01/2006 15:04")
	} else {
		end = "" // no mostrar nada
	}

	c.HTML(http.StatusOK, "item.html", gin.H{
		"Record": record,
		"Start":  start,
		"End":    end,
	})
}

func (rc *RecordController) GetRecordByMachine(c *gin.Context) {
	machineType := c.Query("type")

	records, err := models.GetRecordByMachine(machineType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NO se econtraron registros para ese equipo"})
		return
	}
	c.JSON(http.StatusOK, records)
}
