package mdCos

import "fmt"

const (
	LocationMedia       = "media"
	LocationStreamVideo = LocationMedia + "/streams/test-video"
)

func GetStreamVideoLocation(name string) (location string) {
	location = fmt.Sprintf("%s/%s.m3u8", LocationStreamVideo, name)

	return location
}
