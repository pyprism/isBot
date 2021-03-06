package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-Finder/conf"
	"log"
)

func UpdateIndex(c *gin.Context)  {
	id := c.PostForm("id")
	sterm := c.PostForm("sterm")

	data := Document{
		ID: id,
		SearchTerms: sterm,
	}

	bleveIndex := conf.Get().BleveIndex
	err := bleveIndex.Delete(data.ID)
	if err != nil {
		log.Fatal(err)
	}

	err = bleveIndex.Index(data.ID, data.SearchTerms)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"status":  "saved",
	})
}