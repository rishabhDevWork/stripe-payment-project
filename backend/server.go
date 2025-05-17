package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

func main() {
	stripe.Key = "sk_test_51RP6Xc00MHVKlRaC0wY3a625hA1NdAWJoKmyVYs9hlra9ad5S3LUV4RYr2rzGnaKtdPo8CRPOEjG2RGWy0apPthp0054MR5xwp"
	r := mux.NewRouter()

	r.HandleFunc("/create-payment-intent", handleCreatePaymentIntent).Methods("POST")
	r.HandleFunc("/health", handleHealth).Methods("GET")

	log.Println("Listening on localhost:4242....")
	var err error = http.ListenAndServe("localhost:4242", r)
	if err != nil {
		log.Fatal(err)
	}

}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(paymentIntent.ClientSecret)
	var resp struct {
		ClientSecret string `json:"clientSecret"`
	}
	resp.ClientSecret = paymentIntent.ClientSecret

	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(w, &buf)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func calculateOrderAmount(productId string) int64 {
	switch productId {
	case "Forever Pants":
		return 26000
	case "Forever Shirt":
		return 1550
	case "Forever Shorts":
		return 3000
	}
	return 0
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := "Server is running!!"

	responseByte := []byte(response)
	_, err := w.Write(responseByte)
	if err != nil {
		log.Println(err)
		return
	}
}
