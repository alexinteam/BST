package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var tree = &Tree{}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"val": []string{"required", "numeric"},
	}
	messages := govalidator.MapData{
		"val": []string{"required:Val required"},
	}
	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		Messages:        messages,
		RequiredDefault: true,
	}
	v := govalidator.New(opts)
	e := v.Validate()
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(err)
	}
	val := r.FormValue("val")
	d, found := tree.Find(val)
	if !found {
		w.Header().Set("Content-type", "application/json")
		err := map[string]interface{}{"node not found": e}
		json.NewEncoder(w).Encode(err)
	} else {
		w.Header().Set("Content-type", "application/json")
		result := map[string]interface{}{"node": d}
		json.NewEncoder(w).Encode(result)
	}
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		rules := govalidator.MapData{
			"val": []string{"required", "numeric"},
		}
		messages := govalidator.MapData{
			"val": []string{"required:Val required"},
		}
		opts := govalidator.Options{
			Request:         r,
			Rules:           rules,
			Messages:        messages,
			RequiredDefault: true,
		}
		v := govalidator.New(opts)
		e := v.Validate()
		if len(e) > 0 {
			err := map[string]interface{}{"validationError": e}
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(err)
		}
		val := r.FormValue("val")
		err := tree.Insert(val, val)
		if nil != err {
			w.Header().Set("Content-type", "application/json")
			err := map[string]interface{}{"insert error": err}
			json.NewEncoder(w).Encode(err)
		} else {
			w.Header().Set("Content-type", "application/json")
			result := map[string]interface{}{"result": "insert success"}
			json.NewEncoder(w).Encode(result)
		}
	} else {
		w.Header().Set("Content-type", "application/json")
		result := map[string]interface{}{"result": "only POST allowed"}
		json.NewEncoder(w).Encode(result)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		rules := govalidator.MapData{
			"val": []string{"required", "numeric"},
		}
		messages := govalidator.MapData{
			"val": []string{"required:Val required"},
		}
		opts := govalidator.Options{
			Request:         r,
			Rules:           rules,
			Messages:        messages,
			RequiredDefault: true,
		}
		v := govalidator.New(opts)
		e := v.Validate()
		if len(e) > 0 {
			err := map[string]interface{}{"validationError": e}
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(err)
		}
		val := r.FormValue("val")
		err := tree.Delete(val)
		if nil != err {
			w.Header().Set("Content-type", "application/json")
			err := map[string]interface{}{"delete error": err}
			json.NewEncoder(w).Encode(err)
		} else {
			w.Header().Set("Content-type", "application/json")
			result := map[string]interface{}{"result": "delete success"}
			json.NewEncoder(w).Encode(result)
		}
	} else {
		w.Header().Set("Content-type", "application/json")
		result := map[string]interface{}{"result": "only DELETE allowed"}
		json.NewEncoder(w).Encode(result)
	}
}

func main() {
	var config string
	flag.StringVar(&config, "config", "init.json", "config filename")
	file, err := ioutil.ReadFile(config)
	if nil != err {
		log.Fatal("Config not found", err)
	}
	var parse [30]int64
	_ = json.Unmarshal([]byte(file), &parse)

	// Create a tree and fill it from the json.
	for i := 0; i < len(parse); i++ {
		err := tree.Insert(strconv.FormatInt(parse[i], 10), strconv.FormatInt(parse[i], 10))
		if err != nil {
			log.Fatal("Error inserting value '", parse[i], "': ", err)
		}
	}

	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/delete", deleteHandler)
	fmt.Println("Listening on port: 9000")
	http.ListenAndServe(":9000", nil)
}
