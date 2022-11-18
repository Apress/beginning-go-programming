package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Bid struct {
	Price float64
	URL   string
}

var (
	defaultBid = Bid{
		Price: 0.02,
		URL:   "https://j.mp/3cbDsIY",
	}
	bidTimeout = 10 * time.Millisecond
)

func bidOn(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1)
	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		log.Printf("bid for %q times out, returning default", url)
		return defaultBid
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), bidTimeout)
	defer cancel()
	bid := bidOn(ctx, "https://353solution.com")
	fmt.Println(bid)

	ctx, cancel = context.WithTimeout(context.Background(), bidTimeout)
	defer cancel()
	bid = bidOn(ctx, "https://google.com")
	fmt.Println(bid)
}

/*
OUTPUT
{0.35 https://j.mp/3f3Dpkb}

2022/02/22 15:46:00 bid for "https://example.com" timed out, returning default
{0.02 https://j.mp/3f3Dpkb}

*/
