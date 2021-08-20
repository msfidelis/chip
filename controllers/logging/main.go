package logging

import (
	"github.com/gin-gonic/gin"
	"github.com/bxcodec/faker/v3"
	"chip/libs/logger"
	"fmt"
	"net/http"
)

type response struct {
	Status string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type fake struct {
	UserName string  `faker:"username"`
	PhoneNumber        string  `faker:"phone_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`	
	DayOfWeek          string  `faker:"day_of_week"`
	DayOfMonth         string  `faker:"day_of_month"`
	Timestamp          string  `faker:"timestamp"`
	Century            string  `faker:"century"`
	TimeZone           string  `faker:"timezone"`
	TimePeriod         string  `faker:"time_period"`
	Word               string  `faker:"word"`
	Sentence           string  `faker:"sentence"`
	Paragraph          string  `faker:"paragraph"`
	Currency           string  `faker:"currency"`
	Amount             float64 `faker:"amount"`
	AmountWithCurrency string  `faker:"amount_with_currency"`
	UUIDHypenated      string  `faker:"uuid_hyphenated"`
	UUID               string  `faker:"uuid_digit"`	
}

func Get(c *gin.Context) {
	log := logger.Instance()
	for i := 0; i < 1000; i++ {

		a := fake{}
		err := faker.FakeData(&a)
		if err != nil {
			fmt.Println(err)
		}

		log.Warn().
			Str("username", a.UserName).
			Str("phone", a.PhoneNumber).
			Str("mac_address", a.MacAddress).
			Str("url", a.URL).
			Str("day_of_week", a.DayOfWeek).
			Str("day_of_month", a.DayOfMonth).
			Str("timestamp", a.Timestamp).
			Str("century", a.Century).
			Str("timezone", a.TimeZone).
			Str("period", a.TimePeriod).
			Str("word", a.Word).
			Str("sentence", a.Sentence).
			Str("paragraph", a.Paragraph).
			Str("currency", a.Currency).
			Str("id", a.UUID).
			Str("transaction", a.UUIDHypenated).
			Str("amount", a.AmountWithCurrency).
			Msg("Mock log")
	}

	var res response
	res.Status = "1000 events logged in stdout"
	c.JSON(http.StatusOK, res)
}