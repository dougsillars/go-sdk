# Livestreams endpoint

```golang
//List livestreams
opts := &apivideosdk.LivestreamOpts{
    CurrentPage: 1,
    PageSize:    25,
    StreamKey:   "stream-key",
}
l, err := client.Livestreams.List(opts)

//Get one livestream
l, err := client.Livestreams.Get("livestreamID")

//Create a livestream
livestreamRequest := &apivideosdk.LivestreamRequest{
    Name: "My livetream name",
    Record: false
}
l, err := client.Livestreams.Create(livestreamRequest)

//Update a livestream
livestreamRequest := &apivideosdk.LivestreamRequest{
    Name: "My updated livetream name"
}
l, err := client.Livestreams.Update(livestreamRequest)

//Delete a livestream
err := client.Livestreams.Delete("livestreamID")

//Upload a thumbnail
l, err := client.Livestreams.UploadThumbnail("livestreamID", "path/to/thumbnail.jpg")

//Delete a thumbnail
l, err := client.Livestreams.DeleteThumbnail("livestreamID")

```