It's simple to use micro-image-manager. There are just two endpoints, one to add new images and another one for hosting images (fetching them).
Images are immutable. Since they get uploaded they are there and they never change. It also caches in browsers forever! Storage is cheap, you just upload a new image, get a new ID, update the ID in your entity's record (for example user's profile image field) and it's done. So, this strategy guarantees speed and performance for your users.

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
        ],
        "formats": ["jpeg", "webp"]
    }

and create images in two different size and two different formats for you (jpeg, webp).

## `/image/<IMAGE_ID>.<FORMAT>`

It fetches the biggest image with a format (in this case webp).

    /image/<IMAGE_ID>.webp

Or the size (profile):

    /image/<IMAGE_ID>.jpeg?profile=small

If you put a profile that does not exist you will still get the largest image.
