[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/abdollahpour/micro-image-manager)
[![Coverage Status](https://coveralls.io/repos/github/abdollahpour/micro-image-manager/badge.svg?branch=main)](https://coveralls.io/github/abdollahpour/micro-image-manager?branch=main)
![Build Status](https://github.com/abdollahpour/micro-image-manager/actions/workflows/release.yml/badge.svg)

# micro-image-manager

micro-image-manager is a super-fast microservice written in Go for optimizing, managing, and hosting images in your distributed applications.
You can process and host images in different formats (jpeg, gif, png & webp) and sizes. There's no need for any configuration and setting. The setup procedure could be as easy as running a single command (using Helm).

# Example

Use management endpoint to upload/add a new image. Ex:

    curl -F profile_small=400x300 -F profile_large=800x600 -F image=@my_image.jpg https://micro-image-manager.abdollahpour.com/api/v1/images

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
        <source srcset="//micro-image-manager.abdollahpour.com/image/5f4a459d28317a9f153c211d.webp?profile=large 800w
                        //micro-image-manager.abdollahpour.com/image/5f4a459d28317a9f153c211d.webp?profile=small 400w"
                        type="image/webp" />
        <source srcset="//micro-image-manager.abdollahpour.com/image/5f4a459d28317a9f153c211d.jpeg?profile=large 800w
                        //micro-image-manager.abdollahpour.com/image/5f4a459d28317a9f153c211d.jpeg?profile=small 400w"
                        type="image/webp" />
        <img src="//micro-image-manager.abdollahpour.com/image/5f4a459d28317a9f153c211d.jpeg" />
    </picture>

Please check HTML5 documentation for all possible combinations.

# More documents

- [Install on Kubernetes (Helm)](https://github.com/abdollahpour/helm-charts/tree/main/charts/micro-image-manager)
- [Install on Kubernetes (yaml)](docs/kubernetes.md)
- [Run using Docker](docs/docker.md)
- [Architecture](docs/architecture.md)
- [API documentations](docs/api.md)
