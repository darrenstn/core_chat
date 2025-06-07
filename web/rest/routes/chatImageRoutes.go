package routes

import (
	"core_chat/application/chat/usecase"
	"core_chat/web/rest"
	"core_chat/web/rest/dto"
	restutil "core_chat/web/rest/util"
	webutil "core_chat/web/util"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type ChatImageHandler struct {
	UploadChatImageUC *usecase.UploadChatImageUseCase
	GetChatImageUC    *usecase.GetChatImageUseCase
}

func NewChatImageHandler(uploadChatImgUC *usecase.UploadChatImageUseCase, getChatImgUC *usecase.GetChatImageUseCase) *ChatImageHandler {
	return &ChatImageHandler{
		UploadChatImageUC: uploadChatImgUC,
		GetChatImageUC:    getChatImgUC,
	}
}

func (h *ChatImageHandler) UploadChatImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		rest.SendResponse(w, 400, "Invalid form data")
		return
	}

	sender, ok := webutil.GetIdentifier(r)
	if !ok {
		rest.SendResponse(w, 401, "Unauthorized: sender not found")
		return
	}

	receiver := r.FormValue("receiver")
	if receiver == "" {
		rest.SendResponse(w, 400, "Receiver is required")
		return
	}

	file, _, err := r.FormFile("chat_image")
	if err != nil {
		rest.SendResponse(w, 400, "Failed to get image file")
		return
	}
	defer file.Close()

	// Validate image
	isValid, contentType := restutil.IsValidImage(file)
	if !isValid {
		rest.SendResponse(w, 400, "Invalid image format: "+contentType)
		return
	}

	// Determine extension
	ext := map[string]string{
		"image/jpeg": ".jpg",
		"image/png":  ".png",
	}[contentType]

	// Generate a unique filename
	filename := uuid.NewString()

	// Save the file to disk
	savedPath, err := restutil.SaveImage(file, filename, "./image/chat", ext)
	if err != nil {
		rest.SendResponse(w, 500, "Failed to save image on server")
		return
	}

	err = h.UploadChatImageUC.Execute(savedPath, sender, receiver)
	if err != nil {
		_ = os.Remove(savedPath) // remove image
		rest.SendResponse(w, 400, "Upload image failed: "+err.Error())
		return
	}

	// Construct a public URL (example assumes you serve /uploads from your server)
	//publicURL := fmt.Sprintf("/uploads/chat/image/%s%s", filename, ext) // need to create get image endpoint first
	imageChatUrl := os.Getenv("IMAGE_CHAT_URL")
	chatImgUrl := imageChatUrl + "?image_name=" + filename + ext

	// Respond with the image URL
	res := dto.ChatImageResult{
		Status:     200,
		Message:    "Upload image success",
		ChatImgURL: chatImgUrl,
	}

	rest.SendJSON(w, res)
}

func (h *ChatImageHandler) GetChatImage(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("image_name")
	if imageName == "" {
		rest.SendResponse(w, 400, "Chat image name is required")
		return
	}

	identifier, ok := webutil.GetIdentifier(r)
	if !ok {
		rest.SendResponse(w, 401, "Unauthorized: sender not found")
		return
	}

	imagePath := os.Getenv("DEFAULT_CHAT_IMAGE_DIR")

	result := h.GetChatImageUC.Execute(imageName, imagePath, identifier)
	if !result.Success {
		rest.SendResponse(w, 404, result.Message)
		return
	}

	http.ServeFile(w, r, result.PicturePath)
}
