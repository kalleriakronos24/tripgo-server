package v1

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/config"
	"github.com/kalleriakronos24/khaimal-group/constants"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/pkg/mail-service"
	"github.com/kalleriakronos24/khaimal-group/services"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"github.com/stripe/stripe-go/v81/webhook"
)

// AuthLogin godoc
// @Summary      Method to create payment intent by Stripe
// @Description  A POST Request to create a payment intent by Stripe
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param data body dto.CreatePaymentIntentValidator true "insert payment payload"
// @Router       /payment/create/intent [post]
func POSTCreatePaymentIntent(c *gin.Context) {
	var err error

	userLoggedInId := c.GetString("user_id")
	userId, _ := uuid.Parse(userLoggedInId)

	pValidator := &dto.CreatePaymentIntentValidator{}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	var cred master.Credentials
	if cred, err = services.Handler.RetrieveEntityCredentialsByUserID(userId); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "customer"))
		return
	}

	stripe.Key = config.AppConfig.STRIPE_SECRET_KEY

	customerParams := &stripe.CustomerParams{
		Name:  stripe.String(cred.CredentialCustomer.Name),
		Email: stripe.String(cred.Email),
	}

	var result *stripe.Customer
	if result, err = customer.New(customerParams); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "customer"))
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(pValidator.Amount * 100),
		Currency: stripe.String(pValidator.Currency),
		Customer: stripe.String(result.ID),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		SetupFutureUsage: stripe.String("on_session"),
		PaymentMethodOptions: &stripe.PaymentIntentPaymentMethodOptionsParams{
			PayNow: &stripe.PaymentIntentPaymentMethodOptionsPayNowParams{
				SetupFutureUsage: stripe.String("none"),
			},
			Card: &stripe.PaymentIntentPaymentMethodOptionsCardParams{
				SetupFutureUsage: stripe.String("on_session"),
			},
			Grabpay: &stripe.PaymentIntentPaymentMethodOptionsGrabpayParams{
				SetupFutureUsage: stripe.String("none"),
			},
		},
	}

	pi, err := paymentintent.New(params)
	log.Printf("INTENT ID >>> %v", pi.ID)
	if err != nil {
		// Try to safely cast a generic error to a stripe.Error so that we can get at
		// some additional Stripe-specific information about what went wrong.
		if stripeErr, ok := err.(*stripe.Error); ok {
			log.Printf("Other Stripe error occurred: %v\n", stripeErr.Error())
			// writeJSONErrorMessage(w, stripeErr.Error(), 400)
		} else {
			log.Printf("Other error occurred: %v\n", err.Error())
			// writeJSONErrorMessage(w, "Unknown server error", 500)
		}
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: pi.ClientSecret, Message: "success"})
}

// AuthLogin godoc
// @Summary      Method to create payment intent by Stripe
// @Description  A POST Request to create a payment intent by Stripe
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param data body dto.CreatePaymentIntentValidator true "insert payment payload"
// @Router       /payment/get/webhook [post]
func POSTListenStripeWebhook(c *gin.Context) {
	var err error

	if c.Request.Method != "POST" {
		// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("ioutil.ReadAll: %v", err)
		return
	}

	stripe.Key = config.AppConfig.STRIPE_SECRET_KEY

	event, err := webhook.ConstructEvent(b, c.Request.Header.Get("Stripe-Signature"), config.AppConfig.STRIPE_WEBHOOK_SECRET)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("webhook.ConstructEvent: %v", err)
		return
	}

	// if event.Type == "payment_intent.created" {
	// 	log.Println("Checkout Session completed!")
	// }

	// if event.Type == "payment_intent.payment_failed" {
	// 	log.Println("Checkout Session completed!")
	// }

	// if event.Type == "charge.failed" {
	// 	log.Println("Checkout Session completed!")
	// }

	if event.Type == "payment_intent.succeeded" {
		log.Println("Payment Intent Success")
	}

	if event.Type == "charge.updated" {
		log.Println("Charge Success")
	}

	if event.Type == "charge.succeeded" {
		customerParams := &stripe.CustomerParams{}

		var customerResult *stripe.Customer
		if customerResult, err = customer.Get(event.GetObjectValue("customer"), customerParams); err != nil {
			log.Println("Failed to get Customer")
		}

		var cred master.Credentials
		if cred, err = services.Handler.RetrieveEntityCredentialsByEmail(customerResult.Email); err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "customer"))
			return
		}

		// get last booking transfer id
		var bookingTransfer models.BookingTransfer
		if bookingTransfer, err = services.Handler.RetrieveLastOrderByCustomerID(cred.CredentialCustomer.ID); err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
			return
		}

		paymentMethodParams := &stripe.CustomerRetrievePaymentMethodParams{
			Customer: &customerResult.ID,
		}

		var paymentMethodUsed *stripe.PaymentMethod
		if paymentMethodUsed, err = customer.RetrievePaymentMethod(event.GetObjectValue("payment_method"), paymentMethodParams); err != nil {
			log.Printf("PAYMENT METHOD ERROR >> %v", err)
			return
		}

		if err = mail.SendMailV3(&mail.TSendMail{
			From:    "WadahGo <notification@wadahgo.com>",
			MailTo:  customerResult.Email,
			Subject: "Payment Success",
			Body: fmt.Sprintf(`<html><body>
			<p>Dear Customer, below is a receipt from Stripe that because you have successfully made a payment through WadahGo \n\n <a href="%v">View Receipt</a></p>
			</body></html>`, event.GetObjectValue("receipt_url")),
		}); err != nil {
			return
		}

		parsedAmountPayment, _ := strconv.ParseFloat(event.GetObjectValue("amount"), 64)
		if err := services.Handler.InsertPayment(&models.Payment{
			Amount:            parsedAmountPayment / 100,
			Status:            event.GetObjectValue("status"),
			PI:                event.GetObjectValue("payment_intent"),
			PaymentMethod:     paymentMethodUsed.Card.DisplayBrand,
			BookingTransferID: bookingTransfer.ID,
			CustomerID:        bookingTransfer.CustomerID,
			DriverID:          bookingTransfer.BookingTransferAssigned.DriverID,
			CardLastNumber:    paymentMethodUsed.Card.Last4,
			Currency:          event.GetObjectValue("currency"),
			ReceiptURL:        event.GetObjectValue("receipt_url"),
		}); err != nil {
			c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "payment"))
			return
		}
		// insert the payment table
		// sent the receipt payment email to the customer
		log.Println("charge.updated")
	}

	// refund.created
	// refund.updated
	// charge.refunded
	if event.Type == "charge.refunded" {
		customerParams := &stripe.CustomerParams{}
		var customerResult *stripe.Customer
		log.Printf("CUSTOMER ID >>> %v", event.GetObjectValue("customer"))
		if customerResult, err = customer.Get(event.GetObjectValue("customer"), customerParams); err != nil {
			log.Println("Failed to get Customer")
		}
		// var cred master.Credentials
		// if cred, err = services.Handler.RetrieveEntityCredentialsByEmail(customerResult.Email); err != nil {
		// 	c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "customer"))
		// 	return
		// }
		// // get last booking transfer id
		// var bookingTransfer models.BookingTransfer
		// if bookingTransfer, err = services.Handler.RetrieveLastOrderByCustomerID(cred.CredentialCustomer.ID); err != nil {
		// 	c.JSON(http.StatusBadRequest, constants.GetErrorResponse("data-not-found", err, "booking transfer"))
		// 	return
		// }

		if err = mail.SendMailV3(&mail.TSendMail{
			From:    "WadahGo <notification@wadahgo.com>",
			MailTo:  customerResult.Email,
			Subject: "You've cancelled the order",
			Body: fmt.Sprintf(`<html><body>
			<p>Dear Customer, because you've cancelled below is your Refund receipt from Stripe \n\n <a href="%v">View Receipt</a></p>
			</body></html>`, event.GetObjectValue("receipt_url")),
		}); err != nil {
			return
		}

		tx := database.GetDatabaseConnection().Begin()
		if PaymentErr := tx.Model(&models.Payment{}).Where("pi = ?", event.GetObjectValue("payment_intent")).Updates(&models.Payment{
			ReceiptURL: event.GetObjectValue("receipt_url"),
			Status:     "refunded",
		}); PaymentErr.Error != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}

	c.JSON(http.StatusOK, dto.Response{Data: nil, Message: "success"})
}
