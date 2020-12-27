package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type GoodWatchesMovie struct {
	Name        string `json:"Title"`
	MovieID     string `json:"movieID"`
	ImageLink   string `json:"ImageLink"`
	Description string `json:"Description"`
}

type MovieResponse struct {
	Movies []GoodWatchesMovie
}

func loadDatabase() *sql.DB {
	// godotenv.Load(".env")
	// username := os.Getenv("db_username")
	// password := os.Getenv("db_password")
	// address := os.Getenv("db_hostname")
	// dbName := os.Getenv("db_name")
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
	// 	username,
	// 	password,
	// 	address,
	// 	dbName)
	// fmt.Println(username, password, address, dataSourceName)
	// fmt.Println("attempting to load database...")
	// database, err := sql.Open("mysql", dataSourceName)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// fmt.Println("database successfully loaded!")
	return nil
}

func searchMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`{"message": "oops, we haven't implemented this yet"}`))
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	queryParams := request.URL.Query()
	movieID := queryParams.Get("movieID")
	if queryParams == nil || movieID == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"message": "invalid parameters for getting movie; must present a valid ID"}`))
		return
	}
	sample := GoodWatchesMovie{
		Name:        "Hello World",
		MovieID:     movieID,
		ImageLink:   "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.clovisroundup.com%2Fclovis-west-heads-to-statewide-science-competition%2F&psig=AOvVaw3bNhVI_H8spMIi5_-e5RIA&ust=1609182760741000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCNjl1Ivv7u0CFQAAAAAdAAAAABAD",
		Description: "visit this link: https://www.google.com/imgres?imgurl=https%3A%2F%2Fi.ytimg.com%2Fvi%2F0CeVCbCrH3g%2Fmaxresdefault.jpg&imgrefurl=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3D0CeVCbCrH3g&tbnid=Y1SE7y7D_gw2dM&vet=12ahUKEwiSlPeH7-7tAhWNgp4KHeG1BhMQMygJegQIARBE..i&docid=6zaLYi7j8Ze-IM&w=1280&h=720&q=naasir%20farooqi&hl=en&safe=strict&ved=2ahUKEwiSlPeH7-7tAhWNgp4KHeG1BhMQMygJegQIARBE",
	}
	json.NewEncoder(writer).Encode(sample)
}

func getAllMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	sample1 := GoodWatchesMovie{
		Name:        "Hello World",
		MovieID:     "yawyeet1",
		ImageLink:   "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.clovisroundup.com%2Fclovis-west-heads-to-statewide-science-competition%2F&psig=AOvVaw3bNhVI_H8spMIi5_-e5RIA&ust=1609182760741000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCNjl1Ivv7u0CFQAAAAAdAAAAABAD",
		Description: "visit this link: https://www.google.com/imgres?imgurl=https%3A%2F%2Fi.ytimg.com%2Fvi%2F0CeVCbCrH3g%2Fmaxresdefault.jpg&imgrefurl=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3D0CeVCbCrH3g&tbnid=Y1SE7y7D_gw2dM&vet=12ahUKEwiSlPeH7-7tAhWNgp4KHeG1BhMQMygJegQIARBE..i&docid=6zaLYi7j8Ze-IM&w=1280&h=720&q=naasir%20farooqi&hl=en&safe=strict&ved=2ahUKEwiSlPeH7-7tAhWNgp4KHeG1BhMQMygJegQIARBE",
	}
	sample2 := GoodWatchesMovie{
		Name:        "Hello Again",
		MovieID:     "yawyeet2",
		ImageLink:   "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.clovisroundup.com%2Fclovis-west-heads-to-statewide-science-competition%2F&psig=AOvVaw3bNhVI_H8spMIi5_-e5RIA&ust=1609182760741000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCNjl1Ivv7u0CFQAAAAAdAAAAABAD",
		Description: "visit this link: https://www.google.com/imgres?imgurl=https%3A%2F%2Fi.ytimg.com%2Fvi%2F0CeVCbCrH3g%2Fmaxresdefault.jpg&imgrefurl=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3D0CeVCbCrH3g&tbnid=Y1SE7y7D_gw2dM&vet=12ahUKEwiSlPeH7-7tAhWNgp4KHeG1BhMQMygJegQIARBE..i&docid=6zaLYi7j8Ze-IM&w=1280&h=720&q=naasir%20farooqi&hl=en&safe=strict&ved=2ahUKEwiSlPeH7-7tAhWNgp4KHeG1BhMQMygJegQIARBE",
	}
	sample3 := GoodWatchesMovie{
		Name:        "Hello Part 3",
		MovieID:     "yawyeet3",
		ImageLink:   "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.clovisroundup.com%2Fclovis-west-heads-to-statewide-science-competition%2F&psig=AOvVaw3bNhVI_H8spMIi5_-e5RIA&ust=1609182760741000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCNjl1Ivv7u0CFQAAAAAdAAAAABAD",
		Description: "visit this link: https://www.google.com/imgres?imgurl=https%3A%2F%2Fi.ytimg.com%2Fvi%2F0CeVCbCrH3g%2Fmaxresdefault.jpg&imgrefurl=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3D0CeVCbCrH3g&tbnid=Y1SE7y7D_gw2dM&vet=12ahUKEwiSlPeH7-7tAhWNgp4KHeG1BhMQMygJegQIARBE..i&docid=6zaLYi7j8Ze-IM&w=1280&h=720&q=naasir%20farooqi&hl=en&safe=strict&ved=2ahUKEwiSlPeH7-7tAhWNgp4KHeG1BhMQMygJegQIARBE",
	}
	response := MovieResponse{Movies: []GoodWatchesMovie{sample1, sample2, sample3}}
	json.NewEncoder(writer).Encode(response)
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`{"message": "movie was fake deleted, yay!"}`))
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	queryParams := request.URL.Query()
	movieID := queryParams.Get("movieID")
	movieName := queryParams.Get("movieName")
	image := queryParams.Get("image")
	desc := queryParams.Get("description")
	if queryParams == nil || movieID == "" || movieName == "" || image == "" || desc == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"message": "invalid parameters for creating movie; valid params are movieID, movieName, image, description"}`))
		return
	}
	writer.WriteHeader(http.StatusOK)
	res := GoodWatchesMovie{
		Name:        movieName,
		MovieID:     movieID,
		ImageLink:   image,
		Description: desc,
	}
	json.NewEncoder(writer).Encode(res)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	queryParams := request.URL.Query()
	movieID := queryParams.Get("movieID")
	movieName := queryParams.Get("movieName")
	image := queryParams.Get("image")
	desc := queryParams.Get("description")
	if queryParams == nil || movieID == "" || movieName == "" || image == "" || desc == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"message": "invalid parameters for updating movie; valid params are movieID, movieName, image, description"}`))
		return
	}
	writer.WriteHeader(http.StatusOK)
	res := GoodWatchesMovie{
		Name:        movieName,
		MovieID:     movieID,
		ImageLink:   image,
		Description: desc,
	}
	json.NewEncoder(writer).Encode(res)
}

func main() {
	// database := loadDatabase()
	// http.HandleFunc("/movies", handler)
	// fmt.Println("Listening on port 3000...")
	// err := http.ListenAndServe(":3000", nil)
	// log.Fatal(err)

	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	// Movies endpoints
	api.HandleFunc("/movies/delete", deleteMovie).Methods(http.MethodDelete)
	api.HandleFunc("/movies/create", createMovie).Methods(http.MethodPut)
	api.HandleFunc("/movies/update", createMovie).Methods(http.MethodPut)
	api.HandleFunc("/movies/getByID", getMovie).Methods(http.MethodGet)
	api.HandleFunc("/movies/getAllMovies", getAllMovies).Methods(http.MethodGet)
	api.HandleFunc("/movies/searchMovie", searchMovie).Methods(http.MethodGet)

	dateTime := time.Now()
	fmt.Println("Listening on port 3000 at time: ", dateTime.Format("01-02-2006 15:04:05"))
	log.Fatal(http.ListenAndServe(":3000", router))
}
