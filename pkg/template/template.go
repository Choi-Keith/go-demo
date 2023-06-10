package template

import (
	"bytes"
	"html/template"
)

func generateTemplate() string {
	content := "Somebody asked to reset your password on [{{.SiteName}}].<br><br>\n\nIf it was not you, you can safely ignore this email.<br><br>\n\nClick the following link to choose a new password:<br>\n<a href='{{.PassResetUrl}}' target='_blank'>{{.PassResetUrl}}</a>\n"
	tmpl, err := template.New("update_password").Parse(content)
	if err != nil {
		return ""
	}
	bodyBuf := &bytes.Buffer{}
	templateData := struct {
		SiteName     string
		PassResetUrl string
	}{
		SiteName:     "墨情",
		PassResetUrl: "https://moqing.club",
	}
	// 利用io.Writer将数据传入到模板中
	tmpl.Execute(bodyBuf, templateData)
	return bodyBuf.String()
}
