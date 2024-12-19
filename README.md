# pixel-party
A project which enables users to create their own r/place
like canvas with up to multiple millions of pixels to fill
with their friends.

## Architecture
Every Pixel is stored inside a bitfield in redis.
Every Pixel has 16 bit. That means 4 bit for rgba.
That makes possible values of 0-15 for each color.
Having an canvas of 1000x1000 we have 1.000.000 * 16 bit,
So about 1.6 MB for the full canvas which is possible to send down
to the client.
To get a usable rgba value on the client we only need to multiply by 17.

## Local Dev
```shell
docker run --name redis-local -p 6379:6379 -d redis:latest
```
