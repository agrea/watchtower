package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/agrea/watchtower/config"
	"github.com/sirupsen/logrus"
)

// render is a generic rendering method for all HTML templates.
func render(w http.ResponseWriter, templateName string) {
	tmpl, err := template.ParseFiles(
		filepath.Join(config.TemplatePath, "layout.html"),
		filepath.Join(config.TemplatePath, fmt.Sprintf("%s.html", templateName)),
	)
	if err != nil {
		logrus.WithError(err).Error("Could not render template")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Something went wrong")
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		logrus.WithError(err).Error("An error occurred when executing the template")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Something went wrong")
		return
	}
}
