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
	categoryURL = "http://127.0.0.1:8008/v1/category/"
	productURL  = "http://127.0.0.1:8008/v1/product/"
	commandURL  = "http://127.0.0.1:8008/v1/command/"
	customerURL = "http://127.0.0.1:8008/v1/customer/"
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
	products = [][]string{
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

func main() {
	client := &http.Client{}
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

	id := 0
	for cid, cats := range products {
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

	totalCustomers := 900
	maxOrderedCustomers := 223
	numberOfCustomers := 0
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
}
