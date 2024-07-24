package utils

import (
	"encoding/base64"
	"strings"

	"github.com/google/uuid"
)

func Base64Decoder(base64Images []string) (map[string][]byte, error) {

	base64DecodedImages := make(map[string][]byte)

	for i := 0; i < len(base64Images); i++ {

		base64Image := base64Images[i]
		newImageId := uuid.NewString()
		b64data := base64Image[strings.IndexByte(base64Image, ',')+1:]
		decodedData, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			return base64DecodedImages, err
		}
		newImageId = uuid.NewString()
		base64DecodedImages[newImageId] = decodedData

	}

	return base64DecodedImages, nil
}
