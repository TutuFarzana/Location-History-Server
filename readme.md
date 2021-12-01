# Location History Server
It is a toy in-memory location history server.

## Execution Steps
#### 1. export envs 
```bash
export HISTORY_SERVER_LISTEN_ADDR=8080
```
#### 2. start the application
```bash
go run main.go
```

### Update Location History [PUT]

> http://localhost:8080/location/def456

Sample request 
```bash
{
	"lat": 12.34,
	"lng": 56.78
}
```

Sample response 
```bash
{
    "Status": "ok"
}
```

### Get Location History [GET]

> http://localhost:8080/location/def456?max=2

Sample response 
```bash
{
    "history": [
        {
            "lat": 12.34,
            "lng": 56.78
        },
    ],
    "order_id": "def456"
}
```

### Delete Location History [DELETE]

> http://localhost:8080/location/def456

Sample response 
```bash
{
    "Status": "ok"
}
```