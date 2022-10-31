package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/rusli4k/fevo/app/entities"
)

const (
	statusOptAcc   = "accepted"
	statusOptDec   = "declined"
	paymentOptCash = "cash"
	paymentOptCard = "card"
	bitSize        = 64
	timeLayout     = "2006-01-02 15:04:05"
	dateLayout     = "2006-01-02"
	timeOnlyLayout = " 00:00:00"
	minTime        = "1970-01-01 00:00:00"
	maxTime        = "2100-01-01 00:00:00"
	lenOfDateFrom  = 14
	lenOfDateTo    = 12
)

func CSVToTransactions(file io.Reader) ([]entities.Transaction, error) {
	var ta []entities.Transaction

	var rowWithError int

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("error while reading .csv file: %w", err)
	}

	for i := 1; i < len(lines); i++ {
		var t entities.Transaction

		s := lines[i]

		// ID type of int.
		t.ID, err = strconv.Atoi(s[0])
		if err != nil {
			rowWithError = i

			break
		}

		// RequestID type of int.
		t.RequestID, err = strconv.Atoi(s[1])
		if err != nil {
			rowWithError = i

			break
		}

		// TerminalID type of int.
		t.TerminalID, err = strconv.Atoi(s[2])
		if err != nil {
			rowWithError = i

			break
		}

		// PartnerObjectID type of int.
		t.PartnerObjectID, err = strconv.Atoi(s[3])
		if err != nil {
			rowWithError = i

			break
		}

		// AmountTotal type of float64.
		t.AmountTotal, err = strconv.ParseFloat(s[4], bitSize)
		if err != nil {
			rowWithError = i

			break
		}

		// AmountOriginal type of float64.
		t.AmountOriginal, err = strconv.ParseFloat(s[5], bitSize)
		if err != nil {
			rowWithError = i

			break
		}

		// CommissionPS type of float64.
		t.CommissionPS, err = strconv.ParseFloat(s[6], bitSize)
		if err != nil {
			rowWithError = i

			break
		}

		// CommissionClient type of float.
		t.CommissionClient, err = strconv.ParseFloat(s[7], bitSize)
		if err != nil {
			rowWithError = i

			break
		}

		// CommissionProvider type of float64.
		t.CommissionProvider, err = strconv.ParseFloat(s[8], bitSize)
		if err != nil {
			rowWithError = i

			break
		}

		// DateInput type of time.Time.
		t.DateInput, err = time.Parse(timeLayout, s[9])
		if err != nil {
			rowWithError = i

			break
		}

		// DatePost type of time.Time.
		t.DatePost, err = time.Parse(timeLayout, s[10])
		if err != nil {
			rowWithError = i

			break
		}

		// Status type of string.
		if s[11] != statusOptAcc && s[11] != statusOptDec {
			rowWithError = i

			break
		}

		t.Status = s[11]

		// PaymentType type of string.
		if s[12] != paymentOptCard && s[12] != paymentOptCash {
			rowWithError = i

			break
		}

		t.PaymentType = s[12]

		// PaymentNumber type of string.
		t.PaymentNumber = s[13]

		// ServiceID type of int.
		t.ServiceID, err = strconv.Atoi(s[14])
		if err != nil {
			rowWithError = i

			break
		}

		// Service type of string.
		t.Service = s[15]

		// PayeeID type of int.
		t.PayeeID, err = strconv.Atoi(s[16])
		if err != nil {
			rowWithError = i

			break
		}

		// PayeeName type of string.
		t.PayeeName = s[17]

		// PayeeBankMFO type of int.
		t.PayeeBankMFO, err = strconv.Atoi(s[18])
		if err != nil {
			rowWithError = i

			break
		}

		// PayeeBankAccount type of string.
		t.PayeeBankAccount = s[19]

		// PaymentNarrative type of string.
		t.PaymentNarrative = s[20]

		ta = append(ta, t)
	}

	if rowWithError != 0 {
		return ta, fmt.Errorf("wrong line in .csv file is: %v", rowWithError)
	}

	return ta, nil
}

// ParseDateFromString gets string, return map with time.Time "from" and time.Time "to".
// Return minTime or maxTime if one of filter is missing.
func ParseDateFromString(s string) (map[string]time.Time, error) {
	date := make(map[string]time.Time)

	var err error

	lenOfStringDate := len(s)
	switch lenOfStringDate {
	case lenOfDateFrom:
		{
			dateFromString := s[4:]
			dateFromString += timeOnlyLayout
			date["from"], err = time.Parse(timeLayout, dateFromString)
			if err != nil {
				return nil, fmt.Errorf("error while parsing \"from\": %w", err)
			}
			date["to"], _ = time.Parse(timeLayout, maxTime)
		}
	case lenOfDateTo:
		{
			dateToString := s[2:]
			dateToString += timeOnlyLayout
			date["to"], err = time.Parse(timeLayout, dateToString)
			if err != nil {
				return nil, fmt.Errorf("error while parsing \"to\": %w", err)
			}
			date["from"], _ = time.Parse(timeLayout, minTime)
		}
	case lenOfDateFrom + lenOfDateTo:
		{
			dateFromString := s[4:14]
			dateFromString += timeOnlyLayout
			date["from"], err = time.Parse(timeLayout, dateFromString)
			if err != nil {
				return nil, fmt.Errorf("error while parsing \"from\": %w", err)
			}
			dateToString := s[16:]
			dateToString += timeOnlyLayout
			date["to"], err = time.Parse(timeLayout, dateToString)
			if err != nil {
				return nil, fmt.Errorf("error while parsing \"to\": %w", err)
			}
		}
	default:
		{
			return nil, fmt.Errorf("error while parsing string: %w", err)
		}
	}

	return date, nil
}
