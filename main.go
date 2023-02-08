package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type video struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Licence    string `json:"licence"`
	URL        string `json:"url"`
	PreviewURL string `json:"preview_url"`
	Duration   string `json:"preview_url"`
}

var baseURI = "http://localhost:8081"

var videos = []video{
	{
		ID:         "1",
		Title:      "BigBuckBunny",
		Author:     "Blender",
		Licence:    "Creative Commons",
		URL:        baseURI + "/videos/bigbuckbunny/index.m3u8",
		PreviewURL: baseURI + "/assets/bigbuckbunny_preview.png",
		Duration:   "9min56",
	},
	{
		ID:         "2",
		Title:      "creative commons",
		Author:     "Creative Commons Organization",
		Licence:    "Creative Commons",
		URL:        baseURI + "/videos/creative_commons/index.m3u8",
		PreviewURL: baseURI + "/assets/creative_commons_preview.png",
		Duration:   "1min00",
	},
	{
		ID:         "3",
		Title:      "BigBuckBunny",
		Author:     "Blender",
		Licence:    "Creative Commons",
		URL:        baseURI + "/videos/bigbuckbunny/index.m3u8",
		PreviewURL: baseURI + "/assets/bigbuckbunny_preview.png",
		Duration:   "9min56",
	},
	{
		ID:         "4",
		Title:      "creative commons",
		Author:     "Creative Commons Organization",
		Licence:    "Creative Commons",
		URL:        baseURI + "/videos/creative_commons/index.m3u8",
		PreviewURL: baseURI + "/assets/creative_commons_preview.png",
		Duration:   "1min00",
	},
}

func main() {
	router := gin.Default()
	// serve static files
	router.Static("/assets", "./assets")
	router.Static("/videos", "./videos")

	router.LoadHTMLGlob("templates/*.html")

	// Json API
	router.GET("/api/videos", getVideos)
	router.GET("/api/videos/:id", getVideoByID)

	// Homepage with a list of videos
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"videos": videos,
		})
	})

	// Video player page
	router.GET("/video/:id", func(c *gin.Context) {
		id := c.Param("id")
		video := firstVideoByID(id)
		if video == nil {
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}
		c.HTML(http.StatusOK, "video.html", gin.H{
			"video": video,
		})
	})

	router.Run("localhost:8081")
}

// API endpoint
func getVideos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, videos)
}

func getVideoByID(c *gin.Context) {
	id := c.Param("id")

	// Find the first video with right ID
	for _, a := range videos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "video not found"})
}

// Helper function
func firstVideoByID(id string) *video {
	for i, a := range videos {
		if a.ID == id {
			return &videos[i]
		}
	}
	return nil
}
