[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/abdollahpour/micro-image-manager)
[![Coverage Status](https://coveralls.io/repos/github/abdollahpour/micro-image-manager/badge.svg?branch=main)](https://coveralls.io/github/abdollahpour/micro-image-manager?branch=main)
![Build Status](https://github.com/abdollahpour/micro-image-manager/actions/workflows/release.yml/badge.svg)

# micro-image-manager

One of the biggest challenges to manage media in microservices is distributing them. It's hard to find out how is responsible to store, optimizing, and distribute them. Also, it's hard to implement and it brings a load of dependencies into your service.
micro-image-manager helps to do that simply in an elegant way. You can have images optimized in different format and sizes with a simple HTTP call.

Let's see and example. You need an image. Run the server and then:

    curl -F profile_small=400x300 -F profile_large=800x600 -F image=@my_image.jpg http://localhost:8080/api/v1/images

This will return something like (different ID):

    {
        "id": "5f4a459d28317a9f153c211d",
        "profiles": [
            {
                "name": "small",
                "width": 400,
                "height": 300
            },
            {
                "name": "large",
                "width": 800,
                "height": 600
            }
        ],
        "formats": ["jpeg", "webp"]
    }

You can use HTML5 to host your images, fast and efficiently:

    <picture>
        <source srcset="//localhost:8080/image/5f4a459d28317a9f153c211d.webp" type="image/webp" />
        <img src="//localhost:8080/image/5f4a459d28317a9f153c211d.jpeg" alt="Alt Text!" />
    </picture>

Or even optimize the image for different screen sizes that can boost up performance on mobile phones:

    <img src="//localhost:8080/image/5f4a459d28317a9f153c211d.webp?profile=small"
    srcset="//localhost:8080/image/5f4a459d28317a9f153c211d.webp?profile=large 800w"
    alt="Image description">

Please check HTML5 documentation for all possible combinations.

# More documents

- [Architecture](docs/architecture.md)
- [Run using Kubernetes](docs/kubernetes.md)
- [Test using Docker](docs/docker.md)
- [API documentations](docs/api.md)

# Use helm for kubernetes

    helm repo add micro-image-manager https://raw.githubusercontent.com/abdollahpour/micro-auth-request-chart/master/repository
    helm install micro-auth-request

For more information and customization please check the details [here](https://github.com/abdollahpour/micro-image-manager-chart).
