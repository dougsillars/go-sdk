# Captions endpoint

```golang
//List captions for a video
c, err := client.Captions.List("videoID")

//Get one caption
c, err := client.Captions.Get("videoID", "en")

//Upload a caption file
c, err := client.Captions.Upload("videoID", "en", "path/to/caption.vtt")

//Update a caption default status
captionRequest := &apivideosdk.CaptionRequest{
    Default: true,
}
c, err := client.Captions.Update("videoID", "en", captionRequest)

//Delete a caption
err := client.Captions.Delete("videoID", "en")

```