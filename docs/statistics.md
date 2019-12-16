# Statistics endpoint

```golang
//Get sessions statistics for one video
opts := &apivideosdk.SessionVideoOpts{
    CurrentPage: 1,
    PageSize: 25,
    Period: "2019-12-02",
}
s, err := client.Statistics.GetVideoSessions("videoID", opts)

//Get sessions statistics for one livestream
opts := &apivideosdk.SessionLivestreamOpts{
    CurrentPage: 1,
    PageSize: 25,
    Period: "2019-12-02",
}
s, err := client.Statistics.GetLivestreamSessions("livestreamID", opts)

//Get all events for one session
opts := &apivideosdk.SessionEventOpts{
    CurrentPage: 1,
    PageSize:    25,
}
s, err := client.Statistics.GetSessionEvents("sessionID", opts)

```