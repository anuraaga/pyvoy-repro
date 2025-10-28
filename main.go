package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

func main() {
	var grp sync.WaitGroup
	var cnt atomic.Uint64
	defer func() {
		fmt.Printf("executed %d requests\n", cnt.Load())
	}()
	for range 40 {
		grp.Go(func() {
			for {
				cnt.Add(1)
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
					fmt.Println(err)
					return
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
