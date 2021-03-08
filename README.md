[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/abdollahpour/micro-image-manager)
[![Coverage Status](https://coveralls.io/repos/github/abdollahpour/micro-image-manager/badge.svg?branch=master)](https://coveralls.io/github/abdollahpour/micro-image-manager?branch=master)
[![Build Status](https://secure.travis-ci.org/abdollahpour/micro-image-manager.svg?branch=master)](http://travis-ci.org/abdollahpour/micro-image-manager)

# micro-image-manager

One of the biggest challenges to manage media in microservices is distributing them in the cluster. micro-image-manager help you to manage and optimize them without bringing the whole new info structure and complexity in your app.

micro-image-manager is a simple, fast, and scalable solution to manage, distributed, and optimize your images. Some of the main features are:

* Convert images to different formats and sizes for different target browser and mobile
* Crop & resize images
* Optimize images (jpeg & webp)
* Distributed images resources over your cluster using MongoDB replica

# How to use it
micro-image-manager bring two endpoints:

## `/api/v1/images`
Is the one you use to manage images. You DON'T open this endpoint to the public. You can simply post image to this endpoint:

    curl -F profile_small=400x300 -F profile_large=800x600 -F image=@my_image.jpg http://localhost:8700/api/v1/images

This will return:

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
        ]
    }

and create images in two different size and two different formats for you (jpeg, webp).

## `/image/<IMAGE_ID>`
It fetches the biggest image with a format that the browser supports for you. For example, in chrome, it shows 800x600.webp.

You can explicitly set the format:

    /image/<IMAGE_ID>.jpeg

Or the size (profile):

    /image/<IMAGE_ID>.jpeg?profile=small


## Use it with HTML5
You can use HTML5 to check the format support:

    <picture>
        <source srcset="/image/5f4a459d28317a9f153c211d.webp" type="image/webp" />
        <source srcset="/image/5f4a459d28317a9f153c211d.jpeg"  type="image/jpeg" /> 
        <img src="/image/5f4a459d28317a9f153c211d" alt="Alt Text!" />
    </picture>

Or even optimize the image for different screen sizes that can boost up performance on mobile phones:

    <img src="/image/5f4a459d28317a9f153c211d?profile=small"
    srcset="/image/5f4a459d28317a9f153c211d?profile=large 800w"
    alt="Image description">

# How to run it?

## Use npm
First, you need to have mongodb up and running then:

    npx micro-image-manger

If mongo has different configurations:

    MONGO_URL=mongodb://user@pass:mongo-host:27017/image-manager npx micro-image-manger

Application is ready on localhost:8700

## Use docker

    docker-compose up

Application is ready on `localhost:8700`

# Use helm for kubernetes

    helm repo add micro-image-manager https://raw.githubusercontent.com/abdollahpour/micro-auth-request-chart/master/repository
    helm install micro-auth-request

For more information and customization please check the details [here](https://github.com/abdollahpour/micro-image-manager-chart).