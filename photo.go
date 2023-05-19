// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    photos, err := UnmarshalPhotos(bytes)
//    bytes, err = photos.Marshal()

package main

import "encoding/json"

type Photos []Photo

func UnmarshalPhotos(data []byte) (Photos, error) {
	var r Photos
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Photos) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Photo struct {
	ID                     *string           `json:"id,omitempty"`
	Slug                   *string           `json:"slug,omitempty"`
	CreatedAt              *string           `json:"created_at,omitempty"`
	UpdatedAt              *string           `json:"updated_at,omitempty"`
	PromotedAt             *string           `json:"promoted_at,omitempty"`
	Width                  *int64            `json:"width,omitempty"`
	Height                 *int64            `json:"height,omitempty"`
	Color                  *string           `json:"color,omitempty"`
	BlurHash               *string           `json:"blur_hash,omitempty"`
	Description            *string           `json:"description"`
	AltDescription         *string           `json:"alt_description,omitempty"`
	Urls                   *Urls             `json:"urls,omitempty"`
	Links                  *PhotoLinks       `json:"links,omitempty"`
	Likes                  *int64            `json:"likes,omitempty"`
	LikedByUser            *bool             `json:"liked_by_user,omitempty"`
	CurrentUserCollections []interface{}     `json:"current_user_collections,omitempty"`
	Sponsorship            interface{}       `json:"sponsorship"`
	TopicSubmissions       *TopicSubmissions `json:"topic_submissions,omitempty"`
	User                   *User             `json:"user,omitempty"`
	Exif                   *Exif             `json:"exif,omitempty"`
	Location               *Location         `json:"location,omitempty"`
	Views                  *int64            `json:"views,omitempty"`
	Downloads              *int64            `json:"downloads,omitempty"`
}

type Exif struct {
	Make         *string `json:"make"`
	Model        *string `json:"model"`
	Name         *string `json:"name"`
	ExposureTime *string `json:"exposure_time"`
	Aperture     *string `json:"aperture"`
	FocalLength  *string `json:"focal_length"`
	ISO          *int64  `json:"iso"`
}

type PhotoLinks struct {
	Self             *string `json:"self,omitempty"`
	HTML             *string `json:"html,omitempty"`
	Download         *string `json:"download,omitempty"`
	DownloadLocation *string `json:"download_location,omitempty"`
}

type Location struct {
	Name     *string   `json:"name"`
	City     *string   `json:"city"`
	Country  *string   `json:"country"`
	Position *Position `json:"position,omitempty"`
}

type Position struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}

type TopicSubmissions struct {
	BusinessWork     *BusinessWork `json:"business-work,omitempty"`
	Film             *Animals      `json:"film,omitempty"`
	Nature           *Animals      `json:"nature,omitempty"`
	Animals          *Animals      `json:"animals,omitempty"`
	TexturesPatterns *Animals      `json:"textures-patterns,omitempty"`
	People           *BusinessWork `json:"people,omitempty"`
}

type Animals struct {
	Status     *string `json:"status,omitempty"`
	ApprovedOn *string `json:"approved_on,omitempty"`
}

type BusinessWork struct {
	Status *string `json:"status,omitempty"`
}

type Urls struct {
	Raw     *string `json:"raw,omitempty"`
	Full    *string `json:"full,omitempty"`
	Regular *string `json:"regular,omitempty"`
	Small   *string `json:"small,omitempty"`
	Thumb   *string `json:"thumb,omitempty"`
	SmallS3 *string `json:"small_s3,omitempty"`
}

type User struct {
	ID                *string       `json:"id,omitempty"`
	UpdatedAt         *string       `json:"updated_at,omitempty"`
	Username          *string       `json:"username,omitempty"`
	Name              *string       `json:"name,omitempty"`
	FirstName         *string       `json:"first_name,omitempty"`
	LastName          *string       `json:"last_name"`
	TwitterUsername   *string       `json:"twitter_username"`
	PortfolioURL      *string       `json:"portfolio_url"`
	Bio               *string       `json:"bio"`
	Location          *string       `json:"location"`
	Links             *UserLinks    `json:"links,omitempty"`
	ProfileImage      *ProfileImage `json:"profile_image,omitempty"`
	InstagramUsername *string       `json:"instagram_username"`
	TotalCollections  *int64        `json:"total_collections,omitempty"`
	TotalLikes        *int64        `json:"total_likes,omitempty"`
	TotalPhotos       *int64        `json:"total_photos,omitempty"`
	AcceptedTos       *bool         `json:"accepted_tos,omitempty"`
	ForHire           *bool         `json:"for_hire,omitempty"`
	Social            *Social       `json:"social,omitempty"`
}

type UserLinks struct {
	Self      *string `json:"self,omitempty"`
	HTML      *string `json:"html,omitempty"`
	Photos    *string `json:"photos,omitempty"`
	Likes     *string `json:"likes,omitempty"`
	Portfolio *string `json:"portfolio,omitempty"`
	Following *string `json:"following,omitempty"`
	Followers *string `json:"followers,omitempty"`
}

type ProfileImage struct {
	Small  *string `json:"small,omitempty"`
	Medium *string `json:"medium,omitempty"`
	Large  *string `json:"large,omitempty"`
}

type Social struct {
	InstagramUsername *string     `json:"instagram_username"`
	PortfolioURL      *string     `json:"portfolio_url"`
	TwitterUsername   *string     `json:"twitter_username"`
	PaypalEmail       interface{} `json:"paypal_email"`
}
