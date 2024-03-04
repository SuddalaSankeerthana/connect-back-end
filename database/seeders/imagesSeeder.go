package seeders

import (
	"github.com/Tej-11/connect-backend-application/database/config"
	"github.com/Tej-11/connect-backend-application/database/models"
)

func InsertSeedsInImages() {
	var images = []models.Image{
		{
			ImageId:  "1",
			ImageURL: "https://bd.gaadicdn.com/processedimages/ducati/streetfighter-v2/640X309/streetfighter-v2630c515f9b05e.jpg",
			PostId:   "1",
		},
		{
			ImageId:  "2",
			ImageURL: "https://img.freepik.com/free-photo/fresh-coffee-steams-wooden-table-close-up-generative-ai_188544-8923.jpg",
			PostId:   "2",
		},
		{
			ImageId:  "3",
			ImageURL: "https://images.unsplash.com/photo-1546272192-c19942fa8b26?q=80&w=1000&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTF8fHNvbWV8ZW58MHx8MHx8fDA%3D",
			PostId:   "3",
		},
		{
			ImageId:  "4",
			ImageURL: "https://images.unsplash.com/photo-1607462109225-6b64ae2dd3cb?q=80&w=1000&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjB8fHBob3RvZ3JhcGh5fGVufDB8fDB8fHww",
			PostId:   "4",
		},
		{
			ImageId:  "5",
			ImageURL: "https://static.toiimg.com/photo/77630680.cms?imgsize=780095",
			PostId:   "4",
		},
		{
			ImageId:  "6",
			ImageURL: "https://images.unsplash.com/photo-1607462109225-6b64ae2dd3cb?q=80&w=1000&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjB8fHBob3RvZ3JhcGh5fGVufDB8fDB8fHww",
			PostId:   "5",
		},
		{
			ImageId:  "7",
			ImageURL: "https://images.unsplash.com/photo-1607462109225-6b64ae2dd3cb?q=80&w=1000&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjB8fHBob3RvZ3JhcGh5fGVufDB8fDB8fHww",
			PostId:   "6",
		},
	}
	config.DB.Create(&images)
}
