package logging

import (
	"chip/libs/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
)

type response struct {
	Status  int    `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type fake struct {
	UserName           string  `faker:"username"`
	PhoneNumber        string  `faker:"phone_number"`
	IPV4               string  `faker:"ipv4"`
	IPV6               string  `faker:"ipv6"`
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
	PaymentMethod      string  `faker:"oneof: cc, paypal, check, money order"`
}

// Logs godoc
// @Summary Sent log events to application stdout
// @Tags Logging
// @Produce json
// @Param events query string false "Number of log events; default 1000"
// @Success 200 {object} response
// @Router /logging [get]
func Get(c *gin.Context) {
	log := logger.Instance()

	events := c.DefaultQuery("events", "1000")

	intEvents, err := strconv.Atoi(events)

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < intEvents; i++ {

		a := fake{}
		err := faker.FakeData(&a)
		if err != nil {
			fmt.Println(err)
		}

		log.Warn().
			Str("username", a.UserName).
			Str("phone", a.PhoneNumber).
			Str("mac_address", a.MacAddress).
			Str("ipv4", a.IPV4).
			Str("ipv6", a.IPV6).
			Str("url", a.URL).
			Str("day_of_week", a.DayOfWeek).
			Str("day_of_month", a.DayOfMonth).
			Str("timestamp", a.Timestamp).
			Str("century", a.Century).
			Str("timezone", a.TimeZone).
			Str("period", a.TimePeriod).
			Str("word", a.Word).
			Str("message", a.Sentence).
			Str("paragraph", a.Paragraph).
			Str("currency", a.Currency).
			Str("id", a.UUID).
			Str("transaction", a.UUIDHypenated).
			Str("amount", a.AmountWithCurrency).
			Str("payment_method", a.PaymentMethod).
			Msg("Mock log")
	}

	var res response
	res.Message = fmt.Sprintf("%v logging events sent to stdout", events)
	res.Status = http.StatusOK
	c.JSON(http.StatusOK, res)
}
