package main

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

var firstNames = []string{
	"Emma", "Olivia", "Ava", "Isabella", "Sophia",
	"Mia", "Amelia", "Harper", "Evelyn", "Abigail",
	"Ella", "Grace", "Chloe", "Lily", "Scarlett",
}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones",
	"Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
	"Hernandez", "Wilson", "Anderson", "Taylor", "Thomas",
}

func genName() (string, string) {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	return firstName, lastName
}

func genAmount() float64 {

	minAmount := 50.0
	maxAmount := 5000.0

	amount := minAmount + rand.Float64()*(maxAmount-minAmount)

	amount = float64(int(amount*100)) / 100.0

	return amount
}

type paymentMode string

const (
	CreditCard   paymentMode = "Credit Card"
	PayPal       paymentMode = "PayPal"
	BankTransfer paymentMode = "Bank Transfer"
	UPI          paymentMode = "UPI transaction"
)

func genPaymentMode() paymentMode {
	var paymentModes = []paymentMode{
		CreditCard,
		PayPal,
		BankTransfer,
		UPI,
	}
	getMode := paymentModes[rand.Intn(len(paymentModes))]
	return getMode
}

type payStatus string

const (
	Pending   payStatus = "Pending"
	Completed payStatus = "Completed"
	Failed    payStatus = "Failed"
)

func genePaymentStatus() payStatus {
	var payStatus = []payStatus{
		Pending,
		Completed,
		Failed,
	}
	getstatus := payStatus[rand.Intn(len(payStatus))]
	return getstatus
}

func genPhoneNmber() string {
	areaCode := rand.Intn(8) + 2 
	centralOfficeCode := rand.Intn(8) + 2
	stationCode := rand.Intn(10000)

	return fmt.Sprintf("(%03d)%03d-%04d", areaCode*100+rand.Intn(100), centralOfficeCode*100+rand.Intn(100), stationCode)
}



var locations = []string{
	"New York/Manhattan",
	"Los Angeles/Downtown",
	"Chicago/Lakeview",
	"Houston/Westchase",
	"Phoenix/Desert Ridge",
	"Philadelphia/Centre City",
	"San Antonio/Alamo Heights",
	"San Diego/La Jolla",
	"Dallas/Uptown",
	"San Jose/Willow Glen",
	"Austin/SoCo",
	"Jacksonville/San Marco",
	"San Francisco/Financial District",
	"Indianapolis/Fountain Square",
	"Columbus/Short North",
	"Fort Worth/Stockyards",
	"Charlotte/NoDa",
	"Seattle/Belltown",
	"Denver/Central Business District",
	"El Paso/Westside",
	"Nashville/12 South",
	"Detroit/Midtown",
	"Boston/Back Bay",
	"Memphis/Overton Park",
	"Baltimore/Fells Point",
	"Oklahoma City/Bricktown",
	"Las Vegas/Strip",
	"Louisville/NuLu",
	"Milwaukee/Third Ward",
	"Albuquerque/Old Town",
	"Tucson/Downtown",
	"Fresno/Clovis",
	"Montreal/Plateau",
	"Toronto/Downtown",
	"Vancouver/Yaletown",
	"London/Covent Garden",
	"Paris/Le Marais",
	"Berlin/Mitte",
	"Rome/Trastevere",
	"Madrid/Chueca",
	"Barcelona/Gothic Quarter",
	"Amsterdam/De Pijp",
	"Vienna/Innere Stadt",
	"Zurich/Old Town",
	"Geneva/Carouge",
	"Stockholm/Sodermalm",
	"Copenhagen/Nyhavn",
	"Helsinki/Kamppi",
	"Oslo/Grunerlokka",
}

func genLocation() string {
	Location := locations[rand.Intn(len(locations))]
	return Location
}

func genTime() string {
	now := time.Now()
	formatted := now.Format("03:04PM Mon 02, Jan")
	return formatted
}

type transaction struct {
	UID               string
	FirstName         string
	LastName          string
	PayAmount         float64
	PaymentMode       paymentMode
	PayStatus         payStatus
	DateOfTransaction any
	Location          string
	Contact           string
	Description       string
}

func response() transaction {

	firstName, lastName := genName()

	return transaction{
		UID:               uuid.NewString(),
		FirstName:         firstName,
		LastName:          lastName,
		PayAmount:         genAmount(),
		PaymentMode:       genPaymentMode(),
		PayStatus:         genePaymentStatus(),
		DateOfTransaction: genTime(),
		Location:          genLocation(),
		Contact:           genPhoneNmber()  ,
		Description:       "Payment for services",
	}
}

func sendResponse(limit int) {

	for i := 0; i < limit; i++ {
		response := response()
		fmt.Println(response)
	}
}

func main() {
	sendResponse(17)
}
