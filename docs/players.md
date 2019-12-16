# Players endpoint

```golang
//List Players
opts := &apivideosdk.PlayerOpts{
    CurrentPage: 1,
    PageSize:    25,
}
p, err := client.Players.List(opts)

//Get one player
p, err := client.Players.Get("playerID")

//Create a player
playerRequest := &apivideosdk.PlayerRequest{
    ShapeMargin:           3,
    ShapeRadius:           10,
    ShapeAspect:           "flat",
    ShapeBackgroundBottom: "rgba(255, 0, 0, 0.95)",
    ShapeBackgroundTop:    "rgba(255, 0, 0, 0.95)",
    Text:                  "rgba(255, 0, 0, 0.95)",
    Link:                  "rgba(255, 0, 0, 0.95)",
    LinkHover:             "rgba(255, 0, 0, 0.95)",
    LinkActive:            "rgba(255, 0, 0, 0.95)",
    TrackPlayed:           "rgba(255, 0, 0, 0.95)",
    TrackUnplayed:         "rgba(255, 0, 0, 0.95)",
    TrackBackground:       "rgba(255, 0, 0, 0.95)",
    BackgroundTop:         "rgba(255, 0, 0, 0.95)",
    BackgroundBottom:      "rgba(255, 0, 0, 0.95)",
    BackgroundText:        "rgba(255, 0, 0, 0.95)",
    EnableAPI:             false,
    EnableControls:        true,
    ForceAutoplay:         false,
    HideTitle:             false,
}
p, err := client.Players.Create(playerRequest)

//Update a player
playerRequest := &apivideosdk.PlayerRequest{
    ShapeMargin:           3,
    ShapeRadius:           10,
    ShapeAspect:           "flat",
    ShapeBackgroundBottom: "rgba(255, 255, 0, 0.95)",
    ShapeBackgroundTop:    "rgba(255, 255, 0, 0.95)",
    Text:                  "rgba(255, 255, 0, 0.95)",
    Link:                  "rgba(255, 255, 0, 0.95)",
    LinkHover:             "rgba(255, 255, 0, 0.95)",
    LinkActive:            "rgba(255, 255, 0, 0.95)",
    TrackPlayed:           "rgba(255, 255, 0, 0.95)",
    TrackUnplayed:         "rgba(255, 255, 0, 0.95)",
    TrackBackground:       "rgba(255, 255, 0, 0.95)",
    BackgroundTop:         "rgba(255, 255, 0, 0.95)",
    BackgroundBottom:      "rgba(255, 255, 0, 0.95)",
    BackgroundText:        "rgba(255, 255, 0, 0.95)",
    EnableAPI:             false,
    EnableControls:        true,
    ForceAutoplay:         false,
    HideTitle:             false,
}
p, err := client.Players.Update("playerID", playerRequest)

//Upload a logo for a player
p, err := c.Players.UploadLogo(
    "playerID", 
    "logoLinkURL",
    "path/to/logo.jpg"
)

//Delete a player
err := client.Players.Delete("playerID")

```