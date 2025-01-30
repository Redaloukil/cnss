package main

// TODO: recive interface return struct
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
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
	// weight probability 8:2:1
	weightedPayStatus := []payStatus{
		Completed, Completed, Completed, Completed, Completed, Completed, Completed, Completed,
		Pending, Pending,
		Failed,
	}
	getstatus := weightedPayStatus[rand.Intn(len(weightedPayStatus))]
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
		Contact:           genPhoneNmber(),
	}
}

func index(w http.ResponseWriter, resp *http.Request) {

	message := `🌸🌸 Welcome to the dummy API! 🌸🌸

*--------------------------------*
| Available Endpoints:           |
| - Check Headers : /headers     |
| - Get Data      : /api/v1/data |
*--------------------------------*

🔑 When you request data from /api/v1/data, 
you'll receive a JSON response in following format:

{
  "UID": "<unique identifier>",
  "FirstName": "<user first name>",
  "LastName": "<user last name>",
  "PayAmount": "<user amount>",
  "PaymentMode": "<user payment method>",
  "PayStatus": "<user payment status>",
  "DateOfTransaction": "<user date and time>",
  "Location": "<user location>",
  "Contact": "<user contact number>"
}

🐛 The API mimics data from a women's hospital, using feminine names.
`

	fmt.Fprintf(w, "%v", message)

}

func headers(w http.ResponseWriter, resp *http.Request) {

	for name, headers := range resp.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}

	}
}

func data(w http.ResponseWriter, resp *http.Request) {

	response := response()
	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/v1/data", data)
	http.HandleFunc("/headers", headers)

	log.Println("\n\nServer is running on http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
