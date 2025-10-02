package main

import (
    "user-service/config"
    "user-service/routes"
)

func main() {
    config.ConnectDatabase()
    r := routes.SetupRouter()
    r.Run(":8081") 
}

