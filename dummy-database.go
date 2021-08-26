package main

import (
	pb "github.com/oyamoh-brian/tv-service-database/proto/database"
)



func GetDummyCategories() *[]*pb.Category{
	var categories []*pb.Category = []*pb.Category{
		&pb.Category{
			Id:                   "1",
			Name:                 "tech",
			VideosCount:          2,
			CategoryThumbnailUrl: "https://cdn.api.net/image/1/65x65",
			CategoryCoverUrl:     "https://cdn.api.net/image/1/1024x720",
		},&pb.Category{
			Id:                   "2",
			Name:                 "breaking-news",
			VideosCount:          3,
			CategoryThumbnailUrl: "https://cdn.api.net/image/1/65x65",
			CategoryCoverUrl:     "https://cdn.api.net/image/1/1024x720",
		},
	}

	return &categories
}

func GetDummyVideos(category *pb.Category) *[]* pb.Video {
	var allVideos [] *pb.Video = []* pb.Video {
		&pb.Video{
			Id:                "1",
			Description:       "20",
			Length:            0,
			StreamUrl:         "https://video.api.net/1",
			VideoThumbnailUrl: "https://video.api.net/1/thumb/65x65",
			VideoCoverUrl:     "https://video.api.net/1/cover/1024x720",
			VideoViews:        "20",
			VideoLikes:        "20",
		},
	}

	return &allVideos
}