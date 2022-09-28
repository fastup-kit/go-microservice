package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mehanizm/airtable"
)

type UserData struct {
	Email string `json:"email"`
}

const TABLE_NAME = "subscribers"
const baseID = ""

var AIRTABLE_API_KEY string = ""

func subscribe(c *gin.Context) {
	var userData UserData

	if AIRTABLE_API_KEY == "" {
		AIRTABLE_API_KEY = os.Getenv("AIRTABLE_API_KEY")
	}

	client := airtable.NewClient(AIRTABLE_API_KEY)
	table := client.GetTable(baseID, TABLE_NAME)

	c.BindJSON(&userData)

	recordsToSend := &airtable.Records{
		Records: []*airtable.Record{
			{
				Fields: map[string]interface{}{
					"createdAt": time.Now().Format(time.RFC3339),
					"email":     userData.Email,
				},
			},
		},
	}

	res, err := table.AddRecords(recordsToSend)

	if err != nil {
		println(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"result": res,
		})
	}
}

func main() {
	godotenv.Load()

	router := gin.Default()
	router.POST("/", subscribe)

	router.Run("0.0.0.0:8080")
}
