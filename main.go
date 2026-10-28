package main

import (
	"log"
	"main-health-check/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := handler.ContainerList(); err != nil {
		log.Printf("Docker container listing failed: %v", err)
	}
}

type EngineVersions struct {
	CPU string `json:"cpuEngineVersion"`
	GPU string `json:"gpuEngineVersion"`
}

type DocumentQuery struct {
	DocumentName string `json:"documentName"`
	DocumentID   string `json:"documentId"`
	DocumentType string `json:"documentType"`
}

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.POST("/api/:engineVersion", func(c *gin.Context) {
		var cpuEngineVersion, gpuEngineVersion string
		if c.Param("engineVersion") == "v1" {
			cpuEngineVersion = "v1"
			log.Fatalf("Use %s engine version is CPU hardware AI model", cpuEngineVersion)
		} else {
			gpuEngineVersion = "v2"
			log.Fatalf("Use %s engine version is GPU hardware AI model", gpuEngineVersion)
		}

		if cpuEngineVersion == "" && gpuEngineVersion == "" {
			c.JSON(400, gin.H{
				"error": "invalid engine version",
			})
			return
		}

		c.JSON(200, gin.H{
			"engineVersion":    cpuEngineVersion,
			"gpuEngineVersion": gpuEngineVersion,
		})
	})

	router.GET("/api/:engineVersion/:document/:document_id", func(c *gin.Context) {
		engineVersion := c.Param("engineVersion")
		document := c.Param("document")
		documentID := c.Param("document_id")
		c.JSON(200, gin.H{
			"engineVersion": engineVersion,
			"document":      document,
			"document_id":   documentID,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
