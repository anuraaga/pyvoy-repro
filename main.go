package main

import (
	"net/http"
	"sync"
)

func main() {
	var grp sync.WaitGroup
	for range 40 {
		grp.Go(func() {
			for {
				var prots http.Protocols
				prots.SetUnencryptedHTTP2(true)
				cl := &http.Client{
					Transport: &http.Transport{
						ForceAttemptHTTP2: true,
						Protocols:         &prots,
					},
				}
				req, _ := http.NewRequest("GET", "http://localhost:8000/controlled", nil)
				res, err := cl.Do(req)
				if err != nil {
					panic(err)
				}
				res.Body.Close()
				if res.StatusCode != 200 {
					panic("Invalid response")
				}
			}
		})
	}
	grp.Wait()
}
