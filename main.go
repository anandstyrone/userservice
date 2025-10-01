package main

import (
    "user-service/config"
    "user-service/routes"
)

func main() {
    config.ConnectDB()
    r := routes.SetupRouter()
    r.Run(":8080")
}

