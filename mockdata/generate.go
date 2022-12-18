package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	lorem "github.com/drhodes/golorem"
	"github.com/leslie-qiwa/react-admin-demo/models"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	categories := []string{
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

	client := &http.Client{}
	categoryURL := "http://127.0.0.1:8008/v1/category/"
	ct := "application/json"
	for i, v := range categories {
		cat := models.Category{ID: i + 1, Name: v}
		content, err := json.Marshal(&cat)
		if err != nil {
			log.Fatal(err)
		}
		_, err = client.Post(categoryURL, ct, bytes.NewBuffer(content))
		if err != nil {
			log.Fatal(err)
		}
	}

	productURL := "http://127.0.0.1:8008/v1/product/"
	products := [][]string{
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
				Description: lorem.Paragraph(10, 30),
			}
			content, err := json.Marshal(&product)
			if err != nil {
				log.Fatal(err)
			}
			_, err = client.Post(productURL, ct, bytes.NewBuffer(content))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
