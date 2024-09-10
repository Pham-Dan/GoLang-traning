package domain

type CreatePostRequest struct {
	Title   string `form:"title" validate:"required"`
	Content string `form:"content" validate:"required"`
	// Image string `form:"image" validate:"image"`
}

type UpdatePostRequest struct {
	Title   string `form:"title" validate:"required"`
	Content string `form:"content" validate:"required"`
}

