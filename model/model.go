package model

import "time"

/*
	Pada file ini hanya terdiri model saja untuk data yang akan disajikan.
	Apabila project yang akan dibuat itu besar maka best practice untuk membuat model ini terpisah.

	struct ResponseMeta dan ErrResponse pada saat sekarang tidak digunakan.
*/

type ResponseMeta struct {
	AppStatusCode int    `json:"code"`
	Message       string `json:"statusType,omitempty"`
	ErrorDetail   string `json:"errorDetail,omitempty"`
	ErrorMessage  string `json:"errorMessage,omitempty"`
	DevMessage    string `json:"devErrorMessage,omitempty"`
}

type ErrResponse struct {
	HTTPStatusCode int          `json:"-"` // http response status code
	Status         ResponseMeta `json:"status"`
	AppCode        int64        `json:"code,omitempty"` // application-specific error code
}

type Blogs struct {
	ID              int       `json:"id"`
	BlogName        string    `json:"blog_name"`
	BlogDetails     string    `json:"blog_details,omitempty"`
	BlogDescription string    `json:"blog_description,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type BlogData struct {
	Blog    Blogs  `json:"blog"`
	Message string `json:"message"`
}
