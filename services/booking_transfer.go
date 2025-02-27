package services

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/config"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/onesignal"
	"github.com/kalleriakronos24/khaimal-group/pkg/mail-service"
	"github.com/kalleriakronos24/khaimal-group/templates/email"
	"github.com/kalleriakronos24/khaimal-group/utils"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/refund"
)

type CheckExistingBookingTransferStruct struct {
	*models.BookingTransfer
}

func (module *module) RetrieveLastOrderByCustomerID(userId uuid.UUID) (m models.BookingTransfer, err error) {
	if m, err = module.db.bookingTransfer.GetLastOrderByCustomerId(userId); err != nil {
		return m, errors.New("failed to get booking transfers")
	}
	return
}

func (module *module) RetrieveBookingTransferByUserID(userId uuid.UUID) (m models.BookingTransfer, err error) {
	if m, err = module.db.bookingTransfer.GetOneByID(userId); err != nil {
		return m, errors.New("failed to get booking transfers")
	}
	return
}

func (module *module) RetrieveAllKhaimalBookingTransfer() (m []*models.BookingTransfer, err error) {
	if m, err = module.db.bookingTransfer.GetAllKhaimalBookingOrders(); err != nil {
		return m, errors.New("failed to get booking transfers")
	}
	return
}

func (module *module) RetrieveAllPartnerBookingTransfer(refferalCode string) (m []*models.BookingTransfer, err error) {
	if m, err = module.db.bookingTransfer.GetAllPartnerBookingOrders(refferalCode); err != nil {
		return m, errors.New("failed to get booking transfers")
	}
	return
}

func (module *module) RetrieveAllBookingTransferByCustomer(id uuid.UUID) (m []*models.BookingTransfer, err error) {
	if m, err = module.db.bookingTransfer.GetAllByCustomerID(id); err != nil {
		return m, errors.New("failed to get booking transfers")
	}
	return
}

func (module *module) InsertBookingTransfer(p *dto.InsertBookingTransfer) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	// TODO CHANGE TO NEAREST LOCATION FROM COMPANY ADDRESS
	var carManagement master.CarManagement
	if _, err = module.db.carManagementModel.CheckKhaimalDriverAvailable(p.CarModelID); err == nil {
		if carManagement, err = module.db.carManagementModel.GetKhaimalManagerId(p.CarModelID); err == nil {
			log.Println("THIS1")
		} else {
			var driverBalance *models.BalanceDriver
			if driverBalance, err = module.db.balanceDriver.GetAllDriverHasEnoughBalance(float64(p.Price), p.CarModelID); err != nil {
				tx.Rollback()
				return errors.New("drivers seems busy. please try again later")
			}

			var driverBalanceDetail *models.BalanceDriver
			if driverBalanceDetail, err = module.db.balanceDriver.GetOneDetailByID(driverBalance.ID); err != nil {
				tx.Rollback()
				return errors.New("drivers seems busy. please try again later")
			}
			if carManagement, err = module.db.carManagementModel.GetOneByCarModelIDAndAvailable(p.CarModelID, driverBalanceDetail.Driver.ID, driverBalanceDetail.Driver.Company.ID); err != nil {
				tx.Rollback()
				return errors.New("drivers seems busy. please try again later")
			}
		}
	} else {
		var driverBalance *models.BalanceDriver
		if driverBalance, err = module.db.balanceDriver.GetAllDriverHasEnoughBalance(float64(p.Price), p.CarModelID); err != nil {
			tx.Rollback()
			return errors.New("drivers seems busy. please try again later")
		}
		var driverBalanceDetail *models.BalanceDriver
		if driverBalanceDetail, err = module.db.balanceDriver.GetOneDetailByID(driverBalance.ID); err != nil {
			tx.Rollback()
			return errors.New("drivers seems busy. please try again later")
		}
		if carManagement, err = module.db.carManagementModel.GetOneByCarModelIDAndAvailable(p.CarModelID, driverBalanceDetail.Driver.ID, driverBalanceDetail.Driver.Company.ID); err != nil {
			tx.Rollback()
			return errors.New("drivers seems busy. please try again later")
		}
	}

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	month := int(currentMonth)
	randomUid, _ := utils.GenerateNumber(10)

	var bookingTransfer models.BookingTransfer
	if bookingTransfer, err = module.db.bookingTransfer.InsertBookingTransfer(models.BookingTransfer{
		AdultSeater:       p.AdultSeater,
		ChildSeater:       p.ChildSeater,
		FromLatCoordinate: p.FromLatCoordinate,
		FromLngCoordinate: p.FromLngCoordinate,
		FromLocation:      p.FromLocation,
		ToLocation:        p.ToLocation,
		ToLatCoordinate:   p.ToLatCoordinate,
		ToLngCoordinate:   p.ToLngCoordinate,
		Distance:          p.TotalDistance,
		PassengerNotes:    p.PassengerNotes,
		PickUpDate:        p.PickUpDate,
		Price:             p.Price,
		CarModelID:        p.CarModelID,
		CustomerID:        p.CustomerID,
		AddPickupPoint:    p.AddPickupPoint,
		AddDropPoint:      p.AddDropoffPoint,
		GrandTotal:        p.GrandTotal,
		RefferalCode:      p.RefferalCode,
		PaymentOption:     p.PaymentOption,
		Uid:               fmt.Sprintf("TRF/%v%v/%v", utils.IntegerToRoman(currentYear), utils.IntegerToRoman(month), randomUid),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err = module.db.paymentModel.InsertPayment(models.Payment{
		Amount:            float64(p.GrandTotal),
		Status:            "pending",
		PI:                p.PI,
		PaymentMethod:     strings.ToUpper(p.PaymentOption),
		BookingTransferID: bookingTransfer.ID,
		CustomerID:        p.CustomerID,
		DriverID:          carManagement.DriverID,
		Currency:          "MYR",
		CardLastNumber:    "XXXX",
		ReceiptURL:        "#",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	var customer master.Customer
	if customer, err = module.db.userCustomerModel.GetOneByCustomerID(p.CustomerID); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	// immediately assign to the driver
	if err = module.db.bookingTransferAssigned.InsertBookingTransferAssigned(models.BookingTransferAssigned{
		CarManagementID:   carManagement.ID,
		BookingTransferID: bookingTransfer.ID,
		DriverID:          carManagement.DriverID,
		IsAccepted:        utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	// formattedNotificationMessage := fmt.Sprintf("Booking Transfer Request <br/> %v <br/> %v Person Pax x %v Luggagge <br/> Pickup Date %v <br/> Notes: %v <br/> Price: %v", carManagement.Name, carManagement.CarModel.PersonCount, carManagement.CarModel.LuggageCount, utils.ConvertEnToIDDateTime(bookingTransfer.PickUpDate), bookingTransfer.PassengerNotes, bookingTransfer.GrandTotal)
	// send notification to the selected driver
	// if err = onesignal.PushNotificationSingleExternalId(carManagement.Driver.Credentials.Email, formattedNotificationMessage); err != nil {
	// 	return errors.New("server-error. please try again later")
	// }

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  carManagement.Driver.Credentials.Email,
		Subject: fmt.Sprintf("Booking Received (%v)", bookingTransfer.Uid),
		Body: email.ETOrderSuccess(
			carManagement.Name,
			"#",
			carManagement.Driver.Name,
			"#",
			fmt.Sprintf("%v \n %v", customer.Name, customer.Phone),
			carManagement.PlateNumber,
			"#",
			fmt.Sprintf("%v", carManagement.CarModel.PersonCount),
			fmt.Sprintf("%v", carManagement.CarModel.LuggageCount),
			p.PaymentOption,
			fmt.Sprintf("RM %v", p.GrandTotal),
			"#",
			p.FromLocation,
			p.ToLocation,
			bookingTransfer.Uid,
			utils.ConvertEnToIDDateTime(bookingTransfer.CreatedAt),
			utils.ConvertEnToIDDateTime(p.PickUpDate),
			fmt.Sprintf("You've Received order. Please check the booking information through WadahGo Driver App for booking statuses and details. Booking UID: %v", bookingTransfer.Uid),
			"If issue happened do not hesitate to contact us and ask for help <br/> +60 13 686 8745",
			"#",
			"wadahgo-driver://",
			"CUSTOMER NAME",
		),
	}); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	// send backup email to the driver
	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  customer.Credentials.Email,
		Subject: fmt.Sprintf("Booking Confirmation (%v)", bookingTransfer.Uid),
		Body: email.ETOrderSuccess(
			carManagement.Name,
			"#",
			customer.Name,
			"#",
			carManagement.Driver.Name,
			carManagement.PlateNumber,
			"#",
			fmt.Sprintf("%v", carManagement.CarModel.PersonCount),
			fmt.Sprintf("%v", carManagement.CarModel.LuggageCount),
			p.PaymentOption,
			fmt.Sprintf("RM %v", p.GrandTotal),
			"#",
			p.FromLocation,
			p.ToLocation,
			bookingTransfer.Uid,
			utils.ConvertEnToIDDateTime(bookingTransfer.CreatedAt),
			utils.ConvertEnToIDDateTime(p.PickUpDate),
			fmt.Sprintf("Your was sent to %v and being processed to the next step. You will receive an update from us regarding the status of your Booking ID: %v", carManagement.Driver.Name, bookingTransfer.Uid),
			"If there is no updates from the driver for quite long time, please cancel the order or contact us for assistance <br/> +60 13 686 8745",
			"Track the Booking statuses under wadahgo.com > menu > bookings > booking detail",
			"https://wadahgo.com",
			"DRIVER NAME",
		),
	}); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	tx.Commit()
	return err
}

func (module *module) UpdateBookingTransfer(id uuid.UUID, p *dto.UpdateBookingTransfer) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	if err = module.db.bookingTransfer.UpdateBookingTransfer(id, models.BookingTransfer{
		AdultSeater:       p.AdultSeater,
		ChildSeater:       p.ChildSeater,
		FromLatCoordinate: p.FromLatCoordinate,
		FromLngCoordinate: p.FromLngCoordinate,
		FromLocation:      p.FromLocation,
		ToLocation:        p.ToLocation,
		ToLatCoordinate:   p.ToLatCoordinate,
		ToLngCoordinate:   p.ToLngCoordinate,
		PassengerNotes:    p.PassengerNotes,
		PickUpDate:        p.PickUpDate,
		Price:             p.Price,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	// update driver status to busy and put the car busy also, and waiting to approval by the driver
	tx.Commit()
	return
}

func (module *module) RetrieveAllBookingTransferPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	return
}

func (module *module) RetrieveBookingTransfer(id uuid.UUID) (m models.BookingTransfer, err error) {
	return
}

func (module *module) CustomerCancelBooking(id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	var bookingTransferAssigned models.BookingTransferAssigned
	if bookingTransferAssigned, err = module.db.bookingTransferAssigned.GetOneByID(id); err != nil {
		return errors.New(err.Error())
	}

	if bookingTransferAssigned.IsAccepted == utils.NewTrue() {
		var companyManager master.Driver
		if companyManager, err = module.db.driverModel.GetOneMainAgentByCompanyId(bookingTransferAssigned.Driver.CompanyID); err != nil {
			return errors.New(err.Error())
		}

		var balanceDriver models.BalanceDriver
		if balanceDriver, err = module.db.balanceDriver.GetOneByID(companyManager.ID); err != nil {
			return errors.New(err.Error())
		}

		if BalanceDriverErr := tx.Model(&models.BalanceDriver{}).Where("driver_id", companyManager.ID).Updates(&models.BalanceDriver{
			Amount: utils.ToFixed((balanceDriver.Amount + (float64(bookingTransferAssigned.BookingTransfer.Price) * 0.14)), 1),
		}); BalanceDriverErr.Error != nil {
			tx.Rollback()
			return errors.New("failed to refund driver balance")
		}

		now := time.Now()
		currentYear, currentMonth, _ := now.Date()
		month := int(currentMonth)
		randomUid, _ := utils.GenerateNumber(10)

		// config.AppConfig.APPUrl
		DriverTopup := models.DriverTopup{
			Uid:      fmt.Sprintf("DRV/RF/%v%v/%v", utils.IntegerToRoman(currentYear), utils.IntegerToRoman(month), randomUid),
			Amount:   utils.ToFixed((float64(bookingTransferAssigned.BookingTransfer.Price) * 0.14), 1),
			DriverID: companyManager.ID,
		}

		if DriverTopupErr := tx.Create(&DriverTopup); DriverTopupErr.Error != nil {
			tx.Rollback()
			return errors.New("failed to refund driver balance")
		}

		DriverTransactionHistory := models.DriverTransactionHistory{
			Remark:        "cancel / refund",
			Status:        "refund",
			DriverID:      companyManager.ID,
			DriverTopupID: DriverTopup.ID,
		}

		if DriverTransactionHistoryErr := tx.Create(&DriverTransactionHistory); DriverTransactionHistoryErr.Error != nil {
			tx.Rollback()
			return errors.New("failed to refund driver balance")
		}
	}

	if err := module.db.bookingTransfer.UpdateBookingTransfer(bookingTransferAssigned.BookingTransferID, models.BookingTransfer{
		Status: "passenger cancelled the booking",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err := module.db.bookingTransferAssigned.UpdateBookingTransferAssigned(id, models.BookingTransferAssigned{
		IsAccepted:  utils.NewFalse(),
		IsCancelled: utils.NewTrue(),
		IsOnGoing:   utils.NewFalse(),
		IsPickedUp:  utils.NewFalse(),
		IsCompleted: utils.NewFalse(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}
	if err := module.db.bookingTransfer.UpdateBookingTransfer(id, models.BookingTransfer{
		IsCompleted: utils.NewTrue(),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to set complete booking")
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  bookingTransferAssigned.CarManagement.Driver.Credentials.Email,
		Subject: fmt.Sprintf("Booking Cancelled (%v)", bookingTransferAssigned.BookingTransfer.Uid),
		Body: email.ETOrderSuccess(
			bookingTransferAssigned.CarManagement.Name,
			"#",
			bookingTransferAssigned.CarManagement.Driver.Name,
			"#",
			fmt.Sprintf("%v \n %v", bookingTransferAssigned.BookingTransfer.Customer.Name, bookingTransferAssigned.BookingTransfer.Customer.Phone),
			bookingTransferAssigned.CarManagement.PlateNumber,
			"#",
			fmt.Sprintf("%v", bookingTransferAssigned.CarManagement.CarModel.PersonCount),
			fmt.Sprintf("%v", bookingTransferAssigned.CarManagement.CarModel.LuggageCount),
			bookingTransferAssigned.BookingTransfer.PaymentOption,
			fmt.Sprintf("RM %v", bookingTransferAssigned.BookingTransfer.GrandTotal),
			"#",
			bookingTransferAssigned.BookingTransfer.FromLocation,
			bookingTransferAssigned.BookingTransfer.ToLocation,
			bookingTransferAssigned.BookingTransfer.Uid,
			utils.ConvertEnToIDDateTime(bookingTransferAssigned.UpdatedAt),
			utils.ConvertEnToIDDateTime(bookingTransferAssigned.BookingTransfer.PickUpDate),
			fmt.Sprintf("Oops..., Unfortunately. Customer has Cancelled the Booking UID: %v", bookingTransferAssigned.BookingTransfer.Uid),
			"If issue happened do not hesitate to contact us and ask for help <br/> +60 13 686 8745",
			"#",
			"wadahgo-driver://",
			"CUSTOMER NAME",
		),
	}); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	// if err = mail.SendMailV3(&mail.TSendMail{
	// 	From:    "WadahGo <notification@wadahgo.com>",
	// 	MailTo:  bookingTransferAssigned.CarManagement.Driver.Credentials.Email,
	// 	Subject: "Booking Transfer Cancelled",
	// 	Body: `<html><body>
	// 	<p>Oops.. Customer cancelled your booking due to some reason :(, please keep an eye for another booking request</p>
	// 	</body></html>`,
	// }); err != nil {
	// 	return errors.New(err.Error())
	// }
	formattedNotificationMessage := fmt.Sprintf("Booking Transfer Cancelled \n We are sorry that your Booking was cancelled by the Passenger %v", bookingTransferAssigned.BookingTransfer.Customer.Name)
	// send notification to the selected driver
	onesignal.PushNotificationSingleExternalId(bookingTransferAssigned.CarManagement.Driver.Credentials.Email, formattedNotificationMessage)

	if bookingTransferAssigned.BookingTransfer.PaymentOption != "cash" {
		// update the payment to refund
		stripe.Key = config.AppConfig.STRIPE_SECRET_KEY

		var paymentQuery models.Payment
		if paymentQuery, err = module.db.paymentModel.GetOneLastCreatedByCustomerID(bookingTransferAssigned.BookingTransfer.CustomerID); err != nil {
			return errors.New(err.Error())
		}

		// create the refund request
		params := &stripe.RefundParams{PaymentIntent: stripe.String(paymentQuery.PI), Reason: stripe.String("requested_by_customer")}
		_, err = refund.New(params)

		if err != nil {
			if stripeErr, ok := err.(*stripe.Error); ok {
				log.Printf("Refund Stripe Error: %v\n", stripeErr.Error())
				tx.Rollback()
				return errors.New("failed to refund by stripe. please try again")
			} else {
				log.Printf("Refund Error: %v\n", err.Error())
				tx.Rollback()
				return errors.New("issue when trying to refund customer payment. please try again")
			}
		}
	}

	tx.Commit()
	return
}

func (module *module) GetOutdatedActiveBookingAndSetCancelled() (err error) {
	tx := database.GetDatabaseConnection().Begin()

	// update driver status to busy and put the car busy also, and waiting to approval by the driver
	tx.Commit()
	return
}
