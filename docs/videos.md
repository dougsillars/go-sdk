# Videos endpoint

```golang
//List videos
opts := &apivideosdk.VideoOpts{
    CurrentPage: 1,
    PageSize: 25,
    SortBy:    "publishedAt",
    SortOrder: "desc",
    Tags:     []string{"tag1", "tag2"},
    Metadata: map[string]string{"key": "value"},
}

r, err := client.Videos.List(opts)

//Get one video
r, err := client.Videos.Get("videoID")

//Create a video container
videoRequest := &apivideosdk.VideoRequest{
    Title: "My video title",
    Tags:     []string{"tag1", "tag2"},
    Metadata: map[string]string{"key": "value"},
}
v, err := client.Videos.Create(videoRequest)

//Upload a video to a container
//The upload will be automatically chuncked if the file is more than 128MB
v, err := c.Videos.Upload("videoID", "path/to/video.mp4")

//Update a video
videoRequest := &apivideosdk.VideoRequest{
    Title: "My updated video title",
}
v, err := client.Videos.Update(videoRequest)

//Delete a video
err := client.Videos.Delete("videoID")

//Get video encoding status
v, err := client.Videos.Status("videoID")

//Pick a thumnail with a timecode
v, err := c.Videos.PickThumbnail("videoID", "00:01:12:12")

//Upload a thumnail 
v, err := c.Videos.UploadThumbnail("videoID", "path/to/thumbnail.jpg")

```