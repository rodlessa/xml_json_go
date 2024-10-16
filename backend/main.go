package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	xj "github.com/basgys/goxml2json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/convert", func(c *gin.Context) {
		xmlData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível ler o corpo da requisição"})
			return
		}

		jsonBuffer, err := xj.Convert(strings.NewReader(string(xmlData)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter XML para JSON"})
			return
		}

		jsonString := jsonBuffer.String()

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonString), &result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar JSON"})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	r.Run(":8080")
}
