# fireblocks-sdk-go

A Golang wrapper for the [Fireblocks API](https://docs.fireblocks.com/api)

### build
go build

### Example
cd example  
// help  
go run main.go -h  
// on Fireblocks testnet: with secret file ../../../certs/fireblocks_secret.key and api key 98fe77cb-b459-6508-a945-b01cdda83576  
go run main.go -s ../../../certs/fireblocks_secret.key -k 98fe77cb-b459-6508-a945-b01cdda83576