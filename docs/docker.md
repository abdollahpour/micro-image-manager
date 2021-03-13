You can simply run micro-image-center using docker:

    docker run -d -p 8080:8080 abdollahpour/micro-image-manager

This will run micro-image-manager on port 8080. You can test it using curl as follow:

```bash
curl -F profile_small=400x300 -F profile_large=800x600 -F image=@P1030558.JPG http://localhost:8080/api/v1/images
```

The response would be something like:

```json
{
    "Id": "40406bec-1f82-4e2a-a9df-9209caaa2c5d",
    "Profiles": [
        {
            "Name": "small",
            "Width": 400,
            "Height": 300,
            "Default": false
        },
        {
            "Name": "large",
            "Width": 800,
            "Height": 600,
            "Default": false
        }
    ],
    "Formats": [
        "JPEG",
        "WEBP"
    ]
}
```

By default image will be stored in `/var/lib/micro-image-manager/data` directory and the data is gone when docker image is removed. You can use a named volume to persist data to you host machine:

    docker run -d -p 8080:8080 -v VOLUME_NAME:/var/lib/micro-image-manager/data abdollahpour/micro-image-manager