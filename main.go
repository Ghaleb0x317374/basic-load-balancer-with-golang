package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type servers struct {
	URL *url.URL
}

type config struct {
	stratgey        string
	servers         []servers
	LoadBalanerPort int
}

func newServer(server string) servers {

	URL, _ := url.Parse(server)

	return servers{URL}
}

func (conf *config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server := conf.servers[rand.IntN(len(conf.servers))]
	fmt.Println("we have a request to", server.URL.String())
	proxy := httputil.NewSingleHostReverseProxy(server.URL)

	proxy.ServeHTTP(w, r)
}

func main() {
	var serverList []servers

	serverList = append(serverList, newServer("http://localhost:8080"))
	serverList = append(serverList, newServer("http://localhost:8081"))
	serverList = append(serverList, newServer("http://localhost:8082"))

	conf := config{
		stratgey:        "round_robin",
		servers:         serverList,
		LoadBalanerPort: 511,
	}

	if conf.stratgey == "round_robin" {
		fmt.Println("Round Robin Strategy")
		http.Handle("/", &conf)

		if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.LoadBalanerPort), nil); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}

}
