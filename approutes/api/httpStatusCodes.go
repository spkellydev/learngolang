package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// RedirectError should serve as a temporary variable to carry error messages
var RedirectError struct {
	Message string
}

// Get404 is the default error handler for Bad Requests
func Get404(w http.ResponseWriter, r *http.Request) {
	if len(RedirectError.Message) == 0 {
		RedirectError.Message = "You have made an invalid request"
	}

	t, err := template.ParseFiles("views/httpStatus/404.html")
	if err != nil {
		fmt.Println("error reading template")
	}

	t.Execute(w, RedirectError)
}

// Get500 is the default error handler for Internal Server Errors
func Get500(w http.ResponseWriter, r *http.Request) {
	if len(RedirectError.Message) == 0 {
		RedirectError.Message = "You have made an invalid request"
	}

	t, err := template.ParseFiles("views/httpStatus/500.html")
	if err != nil {
		fmt.Println("error reading template")
	}

	t.Execute(w, RedirectError)
}

// ErrorRequestHandler handles status code redirects
func ErrorRequestHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	switch {
	case statusCode >= 400 && statusCode <= 499:
		RedirectError.Message = "Your request was not valid"
		http.Redirect(w, r, "/404", http.StatusTemporaryRedirect)
		break
	case statusCode >= 500:
		RedirectError.Message = "Your request literally broke the internet. Good job"
		http.Redirect(w, r, "/500", http.StatusTemporaryRedirect)
	default:
		return
	}
}
