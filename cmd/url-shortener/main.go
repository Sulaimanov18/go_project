package main

import (
	"fmt"


	"github.com/yourusername/url-shortener/internal/config"

)

func main() {
	cfg := config.MustLoad()
	fmt.Printf(cfg)



 // TODO : Init config 
 // TODO : Init database connection
 // TODO : Init router
 // TODO : Start server
 // TODO : Handle graceful shutdown
 
}