package models

import (
	"errors"
	"fmt"
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
	"math"
	"time"
)

type Credit struct {
	Id                   int
	Person               valueObjects.PersonInterface `json:"person"`
	Amount               data.Money                   `json:"amount"`
	AgreementAt          data.Date                    `json:"agreementAt"`
	Currency             data.Currency                `json:"currency"`
	Duration             int32                        `json:"duration"`
	Percent              int32                        `json:"percent"`
	Rounding             int32
	RemainAmount         data.Money
	regularPayment       data.Money
	Payments             []valueObjects.PaymentInterface
	isFinished           bool
	remainPaymentsAmount int32
}

func GenCredit(person valueObjects.PersonInterface, amount data.Money, agreementAt data.Date,
	currency data.Currency, duration int32, percent int32) (CreditInterface, error) {
	c := new(Credit)
	c.Person = person
	c.Amount = amount
	c.AgreementAt = agreementAt
	c.Currency = currency
	c.Duration = duration
	c.remainPaymentsAmount = duration
	c.Percent = percent
	annuityPayment := c.getAnnuityPayment(c.Amount, c.Duration)
	rounding, err := c.getRounding(annuityPayment)
	if err != nil {
		return nil, err
	}
	c.Rounding = rounding
	c.regularPayment = c.getRegularPayment(annuityPayment)
	c.Id = -1
	c.isFinished = false
	return c, nil
}

func (c *Credit) GetId() int {
	return c.Id
}

func (c *Credit) SetId(id int) {
	if c.Id == -1 {
		c.Id = id
	} else {
		panic("Your set one more ID!")
	}
}

func (c *Credit) GetPerson() valueObjects.PersonInterface {
	return c.Person
}

func (c *Credit) GetAmount() data.Money {
	return c.Amount
}

func (c *Credit) GetAgreementAt() data.Date {
	return c.AgreementAt
}

func (c *Credit) GetCurrency() data.Currency {
	return c.Currency
}

func (c *Credit) GetDuration() int32 {
	return c.Duration
}

func (c *Credit) GetPercent() int32 {
	return c.Percent
}

func (c *Credit) GetRounding() int32 {
	return c.Rounding
}

func (c *Credit) GetRemainAmount() data.Money {
	return c.RemainAmount
}

func (c *Credit) IsFinished() bool {
	return c.isFinished
}

func (c *Credit) GetPayments(type_ data.Type, state data.State) []valueObjects.PaymentInterface {
	var payments []valueObjects.PaymentInterface
	var tmp_payments []valueObjects.PaymentInterface
	if state == data.State("") || state == data.Paid {
		switch type_ {
		case data.Regular, data.Early:
			{
				for _, payment := range c.Payments {
					if payment.GetType() == type_ {
						payments = append(payments, payment)
					}
				}
			}
		case data.None:
			{
				payments = c.Payments
			}
		}
	}
	if (state == data.Upcoming || state == data.State("")) && !c.IsFinished() {
		switch type_ {
		case data.Regular, data.None:
			{
				payment := c.fetchNextPayment(c.getLastPayment(), data.Next, data.NullDate())
				tmp_payments = append(tmp_payments, payment)
				for payment.GetPayment() != payment.GetFullEarlyRepayment() {
					payment = c.fetchNextPayment(payment, data.Regular, data.NullDate())
					tmp_payments = append(tmp_payments, payment)
				}
				payments = append(payments, tmp_payments...)
			}
		case data.Next:
			{
				payments = append(payments, c.fetchNextPayment(c.getLastPayment(), type_, data.NullDate()))
			}
		}
	}
	return payments
}

func (c *Credit) fetchNextPayment(previousPayment valueObjects.PaymentInterface,
	type_ data.Type, nextPaymentDate data.Date) valueObjects.PaymentInterface {
	var body data.Money
	var date data.Date
	if nextPaymentDate == data.NullDate() {
		date = c.fetchNextPaymentDate(previousPayment)
	} else {
		date = nextPaymentDate
	}

	remainCreditBody := data.Money(previousPayment.GetRemainCreditBody() - previousPayment.GetBody())
	currentPayment := c.regularPayment

	percent := c.fetchPercent(previousPayment.GetDate(), date, remainCreditBody, false)

	if currentPayment-percent < remainCreditBody {
		body = data.Money(currentPayment - percent)
	} else {
		// different order and formulas for payment, body and percent calculation
		// when it is last payment
		currentPayment = data.Money(int64(math.Floor(float64(percent+remainCreditBody)/10.0)) * 10)
		body = data.Money(remainCreditBody)
		percent = data.Money(currentPayment - body)
	}
	return valueObjects.GenPayment(
		currentPayment,
		type_,
		c.GetCurrency(),
		date,
		data.Upcoming,
		percent,
		body,
		remainCreditBody,
		data.Money(math.Floor(float64(remainCreditBody+percent)/10.0)*10))
}

func (c *Credit) fetchNextPaymentDate(previousPayment valueObjects.PaymentInterface) data.Date {
	paymentDayOfMonth := c.AgreementAt.GetDay()

	if previousPayment.GetType() == data.Early {
		paymentDateCandidate := previousPayment.GetDate()
		daysInMonth := time.Date(paymentDateCandidate.GetYear(),
			paymentDateCandidate.GetMonth()+1, 0, 0, 0, 0, 0, time.UTC).Day()
		if paymentDayOfMonth > paymentDateCandidate.GetDay() && daysInMonth > paymentDateCandidate.GetDay() {
			return paymentDateCandidate.SetDay(min(paymentDayOfMonth, daysInMonth))
		}
	}

	paymentDateCandidate := previousPayment.GetDate().AddDate(0, 1, 0)
	daysInMonth := time.Date(paymentDateCandidate.GetYear(),
		paymentDateCandidate.GetMonth()+1, 0, 0, 0, 0, 0, time.UTC).Day()

	if daysInMonth > paymentDayOfMonth {
		return paymentDateCandidate.SetDay(paymentDayOfMonth)
	} else {
		return paymentDateCandidate.SetDay(daysInMonth)
	}
}

func (c *Credit) fetchPercent(from data.Date, to data.Date, creditBody data.Money, inclusiveTo bool) data.Money {
	if from.GetYear() != to.GetYear() && to.GetDay() != 1 {
		firstPercent := c.fetchPercent(from, data.GenDate(from.GetYear(), time.Month(12), 31), creditBody, true)
		secondPercent := c.fetchPercent(data.GenDate(to.GetYear(), time.Month(1), 1), to, creditBody, false)
		return firstPercent + secondPercent
	}
	var percentDays float64
	if inclusiveTo {
		percentDays = math.Round(time.Time(from).Sub(time.Time(to)).Hours()/24 + 1)
	} else {
		percentDays = math.Round(time.Time(from).Sub(time.Time(to)).Hours() / 24)
	}
	yearDays := from.GetNumberOfDays()
	return data.Money(math.Round(float64(creditBody) * float64(c.GetPercent()) * percentDays / 100.0 / yearDays))

}

func (c *Credit) getLastPayment() valueObjects.PaymentInterface {
	if len(c.Payments) == 0 {
		return valueObjects.GenPayment(
			0,
			data.Regular,
			c.Currency,
			c.AgreementAt,
			data.Paid,
			data.Money(0),
			data.Money(0),
			c.GetAmount(),
			c.GetAmount())
	}
	return c.Payments[len(c.Payments)-1]
}

func (c *Credit) WriteOf(paymentRequest valueObjects.PaymentInterface) error {
	paymentType := paymentRequest.GetType()
	lastPayment := c.getLastPayment()
	nextPayment := c.fetchNextPayment(lastPayment, data.Regular, data.NullDate())
	fmt.Printf("%+v\n", paymentRequest)
	fmt.Printf("%+v\n", nextPayment)
	fmt.Printf("%+v\n", lastPayment)
	paymentRequestDate := c.getPaymentRequestDate(paymentRequest, nextPayment)
	isPaymentLikeRegular := false
	fmt.Println(time.Time(nextPayment.GetDate()))
	fmt.Println(time.Time(lastPayment.GetDate()))
	fmt.Println(time.Time(paymentRequestDate))
	fmt.Println(time.Time(paymentRequestDate).Sub(time.Time(nextPayment.GetDate())).Hours() / 24)
	if time.Time(nextPayment.GetDate()).Sub(time.Time(paymentRequestDate)).Hours()/24 >= 1 {
		return errors.New("Payment date is more than next payment date")
	}
	fmt.Println(time.Time(lastPayment.GetDate()).Sub(time.Time(paymentRequestDate)).Hours() / 24)
	if time.Time(lastPayment.GetDate()).Sub(time.Time(paymentRequestDate)).Hours()/24 >= 1 {
		return errors.New("Payment date is outdate")
	}
	if paymentRequestDate != nextPayment.GetDate() && paymentType == data.Regular {
		return errors.New("Payment has got invalid Type")
	}
	if paymentRequestDate == nextPayment.GetDate() && paymentType == data.Early &&
		paymentRequest.GetPayment() == nextPayment.GetPayment() {
		return errors.New("Payment has got invalid Type")
	}
	if nextPayment.GetDate() != paymentRequestDate {
		nextPayment = c.fetchNextPayment(lastPayment, data.Early, paymentRequestDate)
	} else {
		isPaymentLikeRegular = true
	}
	if paymentRequest.GetPayment() < 100 && paymentRequest.GetPayment() != nextPayment.GetPayment() {
		return errors.New("Payment is less than minimal")
	}
	if paymentRequest.GetPayment() > nextPayment.GetFullEarlyRepayment() {
		return errors.New("Payment more than FullEarlyRepaiment")
	}
	if paymentRequest.GetPayment() < c.regularPayment &&
		c.regularPayment < nextPayment.GetFullEarlyRepayment() {
		return errors.New("Payment less than regular")
	}
	if paymentType == data.None {
		if isPaymentLikeRegular {
			paymentType = data.Regular
		} else {
			paymentType = data.Early
		}
	}
	payment := valueObjects.GenPayment(
		paymentRequest.GetPayment(),
		paymentType,
		c.Currency,
		paymentRequestDate,
		data.Paid,
		nextPayment.GetPercent(),
		nextPayment.GetBody(),
		nextPayment.GetRemainCreditBody(),
		nextPayment.GetFullEarlyRepayment())
	fmt.Printf("%+v\n", payment)
	if isPaymentLikeRegular {
		c.remainPaymentsAmount--
	}
	if payment.GetFullEarlyRepayment() == payment.GetPayment() {
		c.isFinished = true
	}
	c.Payments = append(c.Payments, payment)

	if paymentType == data.Early {
		annuityPayment := c.getAnnuityPayment(
			payment.GetRemainCreditBody()-payment.GetBody(),
			c.remainPaymentsAmount)
		c.regularPayment = c.getRegularPayment(annuityPayment)
		if c.regularPayment < 100 {
			c.regularPayment = 100
		}
	}
	return nil
}

func (c *Credit) getPaymentRequestDate(paymentRequest valueObjects.PaymentInterface,
	nextPayment valueObjects.PaymentInterface) data.Date {
	if paymentRequest.GetDate() == data.NullDate() {
		return nextPayment.GetDate()
	} else {
		return paymentRequest.GetDate()
	}
}

func (c *Credit) getAnnuityPayment(amount data.Money, duration int32) float64 {
	monthPercent := float64(c.Percent) / 12.0 / 100.0
	power := math.Pow(1.0+monthPercent, float64(duration))
	return amount.Mon2Float64() * monthPercent * (power / (power - 1.0))
}

func (c *Credit) getRounding(annuityPayment float64) (int32, error) {
	if annuityPayment < 100 {
		return -1, errors.New("Credit payment is less than 100.")
	}
	for _, round := range []int{100, 10, 1} {
		if (round-(int(annuityPayment)%round))*int(c.Duration) < int(annuityPayment) {

			return int32(round), nil
		}
	}
	return -1, errors.New("Credit amount too small for rounding.")
}

func (c *Credit) getRegularPayment(annuityPayment float64) data.Money {
	return data.Money(math.Ceil(annuityPayment/float64(c.Rounding)) * float64(c.Rounding))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
