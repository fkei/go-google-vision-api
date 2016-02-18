package main

import (
	transport "google.golang.org/api/googleapi/transport"
	vision "google.golang.org/api/vision/v1"
	"net/http"
	"log"
	"encoding/base64"
	"os"
	"fmt"
	"flag"
)

// VisionClient vision.Service embed structure
type VisionClient struct {
	*vision.Service
}

// VisionClient create NewVisionClient client
func NewVisionClient(apiKey string) (*VisionClient, error) {
	client := &VisionClient{}
	httpClient := &http.Client{}
	httpClient.Transport = &transport.APIKey{
		Key: apiKey,
	}
	service, err := vision.New(httpClient)
	if err != nil {
		return nil, err
	}
	client.Service = service
	return client, nil
}

func (v *VisionClient) Get(path string) (*vision.BatchAnnotateImagesResponse, error){
	return v.Images.Annotate(&vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{
			&vision.AnnotateImageRequest{
				Features: []*vision.Feature{
					&vision.Feature{
						Type: "FACE_DETECTION",
						MaxResults: 10,
					},
					&vision.Feature{
						Type: "LANDMARK_DETECTION",
						MaxResults: 10,
					},
					&vision.Feature{
						Type: "LOGO_DETECTION",
						MaxResults: 10,
					},
					&vision.Feature{
						Type: "LABEL_DETECTION",
						MaxResults: 10,
					},
					&vision.Feature{
						Type: "TEXT_DETECTION",
						MaxResults: 10,
					},
					&vision.Feature{
						Type: "SAFE_SEARCH_DETECTION",
						MaxResults: 10,
					},
					&vision.Feature{
						Type: "IMAGE_PROPERTIES",
						MaxResults: 10,
					},
				},
				Image: &vision.Image{
					Content: encode(path),
				},
			},
		},
	}).Do()

}

func encode(path string) string {

	file, _ := os.Open(path)
	defer file.Close()
	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

func main() {

	var apiKey string
	flag.StringVar(&apiKey, "api-key", "", "Google API KEY")

	var path string
	flag.StringVar(&path, "path", "", "Image file path")

	flag.Parse()

	if (apiKey == "" || path == "") {
		log.Println("[ERROR] Command Arguments not found.")
		os.Exit(2)
	}
	client, err := NewVisionClient(apiKey)
	res, err := client.Get(path)

	if err != nil {
		log.Panicln(err);
	}

	for i := 0; i < len(res.Responses); i++ {
		json, _ := res.Responses[i].MarshalJSON()
		fmt.Printf("----\n%d:\n%s\n\n", i+1, json);
	}

}
