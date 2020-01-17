# Chapters endpoint

```golang
//List chapters for a video
c, err := client.Chapters.List("videoID")

//Get one chapter
c, err := client.Chapters.Get("videoID", "en")

//Upload a chapter file
c, err := client.Chapters.Upload("videoID", "en", "path/to/chapter.vtt")

//Delete a chapter
err := client.Chapters.Delete("videoID", "en")

```