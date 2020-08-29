[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Coverage Status](https://coveralls.io/repos/github/abdollahpour/micro-image-manager/badge.svg?branch=master)](https://coveralls.io/github/abdollahpour/micro-image-manager?branch=master)
[![Build Status](https://secure.travis-ci.org/abdollahpour/micro-image-manager.svg?branch=master)](http://travis-ci.org/abdollahpour/micro-image-manager)

# micro-image-manager

One of the big challenges to manage media in microservices is to how distributed them, manage and optimize them without bringing whole new infostructure and complexity in your app.

The solution is micro-image-manager. Simple, fast, and scalable solution to manage, distributed, and optimize your images. Some of the main features are:

* Convert images to different formats and sizes for different target browser and mobile
* Crop & resize images
* Optimize images (jpeg & webp)
* Distributed images resources over your cluster using MongoDB replica

# How to use it
micro-image-manager bring two endpoints:

## `/api/v1/images`
Is the one the you use to manage images. You DON'T open this endpoint to public.
You can simply post image to this endpoint:

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
It fetch the largest image with format that browser supports for you. For example it chrome it shows 800x600.webp.

You can explicity put the format:

    /image/<IMAGE_ID>.jpeg

Or size (profile):

    /image/<IMAGE_ID>.jpeg?profile=small


# Use it with HTML5
You can explicity check format support:

    <picture>
        <source srcset="/image/5f4a459d28317a9f153c211d.webp" type="image/webp" />
        <source srcset="/image/5f4a459d28317a9f153c211d.jpeg"  type="image/jpeg" /> 
        <img src="/image/5f4a459d28317a9f153c211d" alt="Alt Text!" />
    </picture>

Or even optimize image for different screen sizes that can boost up performance on mobile phones:

    <img src="/image/5f4a459d28317a9f153c211d?profile=small"
    srcset="/image/5f4a459d28317a9f153c211d?profile=large 800w"
    alt="Image description">