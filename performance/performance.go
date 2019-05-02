package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"math"
	"time"
)

func main() {
	rate := 200
	duration := 4 * time.Second
	login := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "https:https://team-projectv1.herokuapp.com/api/v1/login?login=oks_zh&password=$2a$14$3JFqIzSAXhHk8Opq0/BSxuSWkeiZCiLo2gXmeKt0pQ11MU4YY8O/K",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	//Rate is  vegeta struct to define attack frequency and time
	rt := vegeta.Rate{Freq: rate, Per: duration}
	for res := range attacker.Attack(login, rt, duration, "") {
		metrics.Add(res)
	}
	metrics.Close()
	fmt.Println("Max latency:", metrics.Latencies.Max)
	fmt.Println("Requests:", metrics.Requests)
	fmt.Println("Rate:", math.Round(metrics.Rate))
	fmt.Println("BytesIn:", metrics.BytesIn)
	fmt.Println("BytesOut:", metrics.BytesOut)
	fmt.Println("Errors:", metrics.Errors)
	fmt.Println("StatusCodes: ", metrics.StatusCodes)
	fmt.Println("Success: ", metrics.Success)
}
