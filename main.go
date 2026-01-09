package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"encoding/csv"
	"os"
  
)

type Response struct {
    Type     string    `json:"type"`
    Features []Feature `json:"features"`
}

type Feature struct {
    Type     string   `json:"type"`
    Geometry Geometry `json:"geometry"`
}

type Geometry struct {
    Type        string    `json:"type"`
    Coordinates []float64 `json:"coordinates"`
}

func main() {
resp, err := http.Get("https://data.geopf.fr/geocodage/search?q=paris")
if err != nil {
	
log.Fatalln(err)
}
defer resp.Body.Close()

body, err := ioutil.ReadAll(resp.Body)
if err != nil {
log.Fatalln(err)
}

fmt.Println(string(body))


    var result Response
    err = json.Unmarshal(body, &result)
    if err != nil {
        log.Fatal(err)
    }

    if len(result.Features) > 0 {
        lon := result.Features[0].Geometry.Coordinates[0]
        lat := result.Features[0].Geometry.Coordinates[1]

        fmt.Println("Latitude :", lat)
        fmt.Println("Longitude:", lon)
    }

	// Open the CSV file
   file, err := os.Open("2018_01_Sites_mobiles_2G_3G_4G_France_metropolitaine_L93_ver2(3).csv")
   if err != nil {
       panic(err)
   }
   defer file.Close()

    // Read the CSV data
   reader := csv.NewReader(file)
   reader.FieldsPerRecord = -1 // Allow variable number of fields
   data, err := reader.ReadAll()
   if err != nil {
       panic(err)
   }
   
   // Print the CSV data
   for _, row := range data {
       for _, col := range row {
           fmt.Printf("%s,", col)
       }
       fmt.Println()
   }

   var result_csv Response
    err = json.Unmarshal(body, &result)
    if err != nil {
        log.Fatal(err)
    }

    if len(result_csv.Features) > 0 {
        lon := result_csv.Features[0].Geometry.Coordinates[0]
        lat := result_csv.Features[0].Geometry.Coordinates[1]

        fmt.Println("Latitude_csv :", lat)
        fmt.Println("Longitude_csv:", lon)
    }


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bienvenue sur mon API Go !")
	})

	fmt.Println("Le port 8080 est utilisé pour lancer l'API Go !")
	http.ListenAndServe(":8080", nil)
	
}
