package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Printf("%v", os.Args)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/static"))

	mux.Handle("/", fs)
	mux.HandleFunc("/predict", predictHandler)

	err := http.ListenAndServeTLS(":443", "", "", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func getWordCodeSum(word string) uint32 {
	wordRune := []rune(word)
	wordSum := uint32(0)
	for _, c := range wordRune {
		wordSum += uint32(c)
	}

	return wordSum
}

func getGameResult(home string, guest string) string {
	homeSum := getWordCodeSum(home)
	guestSum := getWordCodeSum(guest)

	if homeSum > guestSum {
		return home
	} else if homeSum < guestSum {
		return guest
	} else {
		return "Ничья"
	}
}

func getGameTotal(home string, guest string) string {
	var total string

	homeSum := getWordCodeSum(home)
	guestSum := getWordCodeSum(guest)

	homeCount := homeSum / uint32(len(home))
	guestCount := guestSum / uint32(len(guest))

	homeGoals := 0
	guestGoals := 0

	if homeSum > guestSum {
		homeGoals = 1 + time.Unix(int64(homeCount), 0).Day()%5
	} else if homeSum < guestSum {
		guestGoals = 1 + time.Unix(int64(guestCount), 0).Day()%5
	} else {
		guestGoals = 1 + time.Unix(int64(guestCount), 0).Day()%5
		homeGoals = guestGoals
	}

	if guestCount%2 == 0 {
		total = fmt.Sprintf("Тм%.1f", float32(homeGoals)+float32(guestGoals)+0.5)
	} else {
		total = fmt.Sprintf("Тм%.1f", float32(homeGoals)+float32(guestGoals)-0.5)
	}

	return total
}

func predictHandler(w http.ResponseWriter, req *http.Request) {
	var predictOptions PredictOptions
	body, _ := io.ReadAll(req.Body)
	json.Unmarshal(body, &predictOptions)
	log.Printf("%s", body)

	gameResult := getGameResult(predictOptions.Home, predictOptions.Guest)
	gameTotal := getGameTotal(predictOptions.Home, predictOptions.Guest)

	predict := Predict{gameResult, gameTotal}
	rb, _ := json.Marshal(predict)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, string(rb))
}
