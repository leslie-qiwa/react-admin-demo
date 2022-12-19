package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/leslie-qiwa/react-admin-demo/models"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	ctype       = "application/json"
	categoryURL = "http://127.0.0.1:8008/v1/category"
	productURL  = "http://127.0.0.1:8008/v1/product"
	commandURL  = "http://127.0.0.1:8008/v1/command"
	customerURL = "http://127.0.0.1:8008/v1/customer"
	reviewURL   = "http://127.0.0.1:8008/v1/review"
)

var (
	categories = []string{
		"animals",
		"beard",
		"business",
		"cars",
		"city",
		"flowers",
		"food",
		"nature",
		"people",
		"sports",
		"tech",
		"travel",
		"water",
	}
	productNames = [][]string{
		{
			"Cat Nose",
			"Dog Grass",
			"Brown Cow",
			"Leopard Road",
			"Sad Dog",
			"Pelican Pier",
			"Green Birds",
			"Concrete Seaguls",
			"Hiding Seagul",
			"Sand Caravan",
		},
		{
			"Black Auburn",
			"Basket Beard",
			"Handlebar Moustache",
			"White Beard",
			"Sailor Man",
			"Natural Beard",
			"Yeard Phone",
			"Braid Beard",
			"Terminal Black",
			"Short Boxed",
		},
		{
			"Corporate Prop",
			"Office Chairs",
			"White Clock",
			"Work Suit",
			"Suit & Tie",
			"Shake Hands",
			"Building Sky",
			"Yellow Pad",
			"Work Devices",
			"Hands Clap",
			"Work Meeting",
		},
		{
			"Old Combi",
			"Asian Plates",
			"Pedestrian Crossing",
			"Farmer Boy",
			"Make Over",
			"Sports Sunset",
			"Desert Jeep",
			"Highway Bridge",
			"Race Stickers",
			"White Deluxe",
		},
		{
			"Bridge Lights",
			"Color Dots",
			"Cloud Suspension",
			"Paved Street",
			"Blue Bay",
			"Wooden Door",
			"Concrete Angles",
			"London Lights",
			"Fort Point",
			"Rainy Glass",
		},
		{
			"Apricot Tree",
			"Orange Rose",
			"Purple Petunia",
			"Water Lily",
			"White Peony",
			"Poppy Field",
			"Blue Flax",
			"Love Roses",
			"California Poppy",
			"Dalhia Colors",
		},
		{
			"Fuzzy Forks",
			"Stamp Mug",
			"Two Expressos",
			"Red Latte",
			"Black Grapes",
			"Forgotten Strawberries",
			"Close Steam",
			"Brewing Tea",
			"Red Onions",
			"Dark Honey",
		},
		{
			"Distant Mountains",
			"Fog Pond",
			"Sand Rocks",
			"Pebble Shore",
			"Eroded Fractals",
			"Water Fall",
			"Drif Wood",
			"Dirt Track",
			"Green Grass",
			"Yellow Lichen",
		},
		{
			"Crossing Alone",
			"Budding Grove",
			"Light Hair",
			"Black & White",
			"Rock Concert",
			"Meeting Bench",
			"Son & Lumière",
			"Running Boy",
			"Dining Hall",
			"Tunnel People",
		},
		{
			"Feather Ball",
			"Wall Skate",
			"Kick Flip",
			"Down Hill",
			"Baseball Night",
			"Touch Line",
			"Alone Jogger",
			"Green Basket",
			"Mud Hug",
			"Metal Cycle",
		},
		{
			"Black Screen",
			"Phone Call",
			"Tablet & Phone",
			"No Battery",
			"Phone Book",
			"Camera Parts",
			"Fuzzy Phone",
			"Music & Light",
			"Eye Rest",
			"Aligned Parts",
		},
		{
			"Distant Jet",
			"Foggy Beach",
			"White Lime",
			"Mysterious Cloud",
			"Mountain Top",
			"Light House",
			"Gray Day",
			"Desert Walkway",
			"Train Track",
			"Plane Trees",
		},
		{
			"Fresh Stream",
			"Reed Line",
			"Mud Tracks",
			"Beach Gazebo",
			"Calm Sea",
			"Early Bath",
			"Aerial Coast",
			"Canal Street",
			"Artificial Beach",
			"Rainy Day",
		},
	}
)

func weightedBoolean(likelyhood int) bool {
	return rand.Intn(99) < likelyhood
}

func weightedArrayElement(values, weights []int) int {
	histogram := []int{}
	for idx, v := range weights {
		for i := 0; i < v; i++ {
			histogram = append(histogram, values[idx])
		}
	}
	return histogram[rand.Intn(len(histogram))]
}

func main() {
	client := &http.Client{}

	// insert category
	for i, v := range categories {
		cat := models.Category{ID: i + 1, Name: v}
		content, err := json.Marshal(&cat)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Post(categoryURL, ctype, bytes.NewBuffer(content))
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
	}

	// insert product
	id := 0
	products := []models.Product{}
	for cid, cats := range productNames {
		for pid, name := range cats {
			id += 1
			width := 10 + rand.Intn(30)
			height := 10 + rand.Intn(30)
			product := models.Product{
				ID:          id,
				CategoryID:  cid + 1,
				Reference:   name,
				Width:       float32(width),
				Height:      float32(height),
				Price:       float32(width*height/20 + rand.Intn(width*height/15-width*height/20)),
				Thumbnail:   fmt.Sprintf("https://marmelab.com/posters/%s-%d.jpeg", categories[0], pid+1),
				Image:       fmt.Sprintf("https://marmelab.com/posters/%s-%d.jpeg", categories[0], pid+1),
				Stock:       rand.Intn(150),
				Description: gofakeit.Paragraph(1, 3, 30, ","),
			}
			products = append(products, product)
			content, err := json.Marshal(&product)
			if err != nil {
				log.Fatal(err)
			}
			resp, err := client.Post(productURL, ctype, bytes.NewBuffer(content))
			if err != nil {
				log.Fatal(err)
			}
			resp.Body.Close()
		}
	}

	// insert customer
	totalCustomers := 900
	maxOrderedCustomers := 223
	numberOfCustomers := 0
	realCustomers := []models.Customer{}
	reviewCustomers := map[int]models.Customer{}
	for i := 1; i <= totalCustomers; i++ {
		customer := models.Customer{
			ID:            i,
			FirstName:     gofakeit.FirstName(),
			LastName:      gofakeit.LastName(),
			HasNewsletter: true,
			FirstSeen:     gofakeit.DateRange(time.Now().Add(-5*365*24*time.Hour), time.Now()),
			LastSeen:      gofakeit.DateRange(time.Now().Add(-5*365*24*time.Hour), time.Now()),
			Birthday: gofakeit.DateRange(
				time.Now().Add(-60*365*24*time.Hour),
				time.Now().Add(-20*365*24*time.Hour)),
		}
		customer.Email = customer.FirstName + "." + customer.LastName + "@fake.io"
		hasOrdered := weightedBoolean(25) && numberOfCustomers < maxOrderedCustomers
		if hasOrdered {
			numberOfCustomers++
			customer.Address = gofakeit.Street()
			customer.Zipcode = gofakeit.Zip()
			customer.City = gofakeit.City()
			customer.StateAbbr = gofakeit.StateAbr()
			customer.Avatar = "https://marmelab.com/posters/avatar-" + strconv.Itoa(numberOfCustomers) +
				".jpeg"
			customer.HasOrdered = true
			customer.HasNewsletter = weightedBoolean(30)
			realCustomers = append(realCustomers, customer)

			// only 60% of buyers write reviews
			if weightedBoolean(60) {
				reviewCustomers[customer.ID] = customer
			}
		}
		content, err := json.Marshal(&customer)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Post(customerURL, ctype, bytes.NewBuffer(content))
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
	}

	// insert order
	reviewID := 1
	totalCommands := 600
	taxRate := []float32{0.12, 0.17, 0.2}
	commands := make([]models.Command, totalCommands)
	for i := 1; i <= totalCommands; i++ {
		customer := realCustomers[rand.Intn(len(realCustomers))]
		command := models.Command{
			ID:         i,
			CustomerID: customer.ID,
			Reference:  gofakeit.LetterN(6),
			Date:       gofakeit.DateRange(customer.FirstSeen, customer.LastSeen),
		}
		nbProducts := weightedArrayElement(
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{30, 20, 5, 2, 1, 1, 1, 1, 1, 1},
		)
		quantity := weightedArrayElement(
			[]int{1, 2, 3, 4, 5},
			[]int{10, 5, 3, 2, 1},
		)
		command.Baskets = make([]models.Basket, nbProducts)
		for i := 0; i < nbProducts; i++ {
			basket := models.Basket{
				ProductID: products[rand.Intn(len(products))].ID,
				Quantity:  quantity,
			}
			command.Baskets[i] = basket
			command.TotalExTaxes += float32(basket.Quantity) * products[basket.ProductID-1].Price
		}
		command.DeliveryFees = float32(rand.Intn(500))/100.0 + 3.0
		command.TaxRate = taxRate[rand.Intn(3)]
		command.Taxes = (command.TotalExTaxes + command.DeliveryFees) * command.TaxRate
		command.Total = command.TotalExTaxes + command.DeliveryFees + command.Taxes

		if command.Date.After(time.Now().Add(-90*24*time.Hour)) && rand.Intn(2) == 1 {
			command.Status = models.StatusOrdered
		} else if rand.Intn(10) == 1 {
			command.Status = models.StatusCanceled
		} else {
			command.Status = models.StatusDelivered
			if weightedBoolean(10) {
				command.Returned = true
			}
		}
		commands[i-1] = command

		content, err := json.Marshal(&command)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Post(commandURL, ctype, bytes.NewBuffer(content))
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()

		_, hasReview := reviewCustomers[customer.ID]
		if !hasReview {
			continue
		}

		// insert review
		for _, basket := range command.Baskets {
			// reviewers review 40% of their products
			if !weightedBoolean(40) {
				continue
			}
			review := models.Review{
				ID:         reviewID,
				CommandID:  command.ID,
				ProductID:  basket.ProductID,
				CustomerID: customer.ID,
				Rating:     1 + rand.Intn(4),
				Comment:    gofakeit.SentenceSimple(),
				Date:       gofakeit.DateRange(command.Date, time.Now()),
			}
			if review.Date.After(review.Date.Add(-30 * 24 * time.Hour)) {
				if rand.Intn(4) == 1 {
					review.Status = models.StatusRejected
				} else {
					review.Status = models.StatusAccepted
				}
			} else {
				switch weightedArrayElement([]int{1, 2, 3}, []int{5, 3, 1}) {
				case 1:
					review.Status = models.StatusPending
				case 2:
					review.Status = models.StatusAccepted
				case 3:
					review.Status = models.StatusRejected
				}
			}
			reviewID++

			content, err := json.Marshal(&review)
			if err != nil {
				log.Fatal(err)
			}
			resp, err := client.Post(reviewURL, ctype, bytes.NewBuffer(content))
			if err != nil {
				log.Fatal(err)
			}
			resp.Body.Close()
		}
	}

}
