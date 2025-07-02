package routes

import (
	"core_chat/application/person/usecase"
	"core_chat/web/rest"
	"core_chat/web/rest/dto"
	"core_chat/web/rest/mapper"
	"core_chat/web/rest/util"
	"os"
	"time"

	// "encoding/json"

	"net/http"
)

type PersonHandler struct {
	RegisterUC               *usecase.RegisterPersonUseCase
	EmailAvailabilityUC      *usecase.EmailAvailabilityUseCase
	IdentifierAvailabilityUC *usecase.IdentifierAvailabilityUseCase
	GetProfileUC             *usecase.GetProfileUseCase
	GetProfileImageUC        *usecase.GetProfileImageUseCase
}

func NewPersonHandler(emailAvailabilityUC *usecase.EmailAvailabilityUseCase, identifierAvailabilityUC *usecase.IdentifierAvailabilityUseCase, registerUC *usecase.RegisterPersonUseCase, getProfileUC *usecase.GetProfileUseCase, getProfileImageUC *usecase.GetProfileImageUseCase) *PersonHandler {
	return &PersonHandler{EmailAvailabilityUC: emailAvailabilityUC, IdentifierAvailabilityUC: identifierAvailabilityUC, RegisterUC: registerUC, GetProfileUC: getProfileUC, GetProfileImageUC: getProfileImageUC}
}

func (h *PersonHandler) CheckIdentifierAvailability(w http.ResponseWriter, r *http.Request) {
	identifier := r.FormValue("identifier")

	result := h.IdentifierAvailabilityUC.Execute(identifier)

	if !result {
		rest.SendResponse(w, 400, "Identifier already in use or not valid")
		return
	}

	rest.SendResponse(w, 200, "Identifier is available")
}

func (h *PersonHandler) CheckEmailAvailability(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	result := h.EmailAvailabilityUC.Execute(email)

	if !result {
		rest.SendResponse(w, 400, "Email already in use or not valid")
		return
	}

	rest.SendResponse(w, 200, "Email is available")
}

func (h *PersonHandler) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		rest.SendResponse(w, 400, "Invalid form data")
		return
	}

	uploadProfilePicture := true
	picturePath := ""

	file, _, err := r.FormFile("profile_picture")
	if err != nil {
		uploadProfilePicture = false
	} else {
		defer file.Close()
	}

	if uploadProfilePicture {
		isValid, contentType := util.IsValidImage(file)
		if !isValid {
			rest.SendResponse(w, 400, "Invalid image format: "+contentType)
			return
		}

		filename := r.FormValue("identifier")

		ext := map[string]string{
			"image/jpeg": ".jpg",
			"image/png":  ".png",
		}[contentType]

		imageTempProfileDir := os.Getenv("IMAGE_TEMP_PROFILE_DIR")

		picturePath, err = util.SaveImage(file, filename, imageTempProfileDir, ext)
		if err != nil {
			rest.SendResponse(w, 500, "Failed to save image")
			return
		}
	}

	// Decode other fields (non-file)
	var req dto.RegisterRequest
	req.Identifier = r.FormValue("identifier")
	req.Email = r.FormValue("email")
	req.Password = r.FormValue("password")

	req.Name = r.FormValue("name")
	if req.Name == "" {
		rest.SendResponse(w, 400, "Name is required")
		return
	}

	dobStr := r.FormValue("date_of_birth")
	if dobStr == "" {
		rest.SendResponse(w, 400, "Date of birth is required")
		return
	}

	parsedDOB, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		rest.SendResponse(w, 400, "Invalid date format, expected YYYY-MM-DD")
		return
	}

	req.DateOfBirth = parsedDOB
	req.Description = r.FormValue("description")

	// Map and call use case
	appRequest := mapper.ToRegisterRequest(req, picturePath)
	result := h.RegisterUC.Execute(appRequest)

	if !result.Success {
		rest.SendResponse(w, 400, result.Message)
		return
	}
	rest.SendResponse(w, 200, result.Message)
}

func (h *PersonHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	identifier := r.URL.Query().Get("identifier")
	if identifier == "" {
		rest.SendResponse(w, 400, "Identifier is required")
		return
	}

	result := h.GetProfileUC.Execute(identifier)
	if !result.Success {
		rest.SendResponse(w, 400, result.Message)
		return
	}

	imageProfileUrl := os.Getenv("IMAGE_PROFILE_URL")

	profileUrl := imageProfileUrl + "?identifier=" + result.Identifier

	profileResult := mapper.ToProfileResult(result, profileUrl, 200)

	rest.SendJSON(w, profileResult)
}

func (h *PersonHandler) GetProfileImage(w http.ResponseWriter, r *http.Request) {
	identifier := r.URL.Query().Get("identifier")
	if identifier == "" {
		rest.SendResponse(w, 400, "Identifier is required")
		return
	}

	result := h.GetProfileImageUC.Execute(identifier)
	if !result.Success {
		rest.SendResponse(w, 404, result.Message)
		return
	}

	http.ServeFile(w, r, result.PicturePath)
}
