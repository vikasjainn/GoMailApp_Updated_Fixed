package Common

import (
    "bytes"
    "encoding/json"
    "fmt"
    "html/template"
    "io"
    "net/http"
)

// EmailPayload defines the structure expected by your API
type EmailPayload struct {
    ApiKey        string        `json:"ApiKey"`
    HtmlData      string        `json:"HtmlData"`
    SenderEmail   string        `json:"SenderEmail"`
    ReceiverEmail string        `json:"ReceiverEmail"`
    SubjectLine   string        `json:"SubjectLine"`
    AttachDocs    []interface{} `json:"AttachDocs"`
}

// parseTemplate loads and renders the HTML template with the given user name
func parseTemplate(name string) (string, error) {
    tmpl, err := template.ParseFiles(TemplatePath)
    if err != nil {
        return "", err
    }

    var tpl bytes.Buffer
    err = tmpl.Execute(&tpl, struct{ Name string }{Name: name})
    if err != nil {
        return "", err
    }

    return tpl.String(), nil
}

// SendTemplatedEmail constructs and sends the email using your vendor API
func SendTemplatedEmail(name, to string) error {
    htmlContent, err := parseTemplate(name)
    if err != nil {
        return fmt.Errorf("failed to parse HTML template: %w", err)
    }

    payload := EmailPayload{
        ApiKey:        ApiKey,
        HtmlData:      htmlContent,
        SenderEmail:   SenderEmail,
        ReceiverEmail: to,
        SubjectLine:   EmailSubject,
        AttachDocs:    []interface{}{},
    }

    jsonPayload, err := json.Marshal(payload)
    if err != nil {
        return fmt.Errorf("failed to marshal JSON: %w", err)
    }

    req, err := http.NewRequest("POST", ApiEndpoint, bytes.NewBuffer(jsonPayload))
    if err != nil {
        return fmt.Errorf("failed to create HTTP request: %w", err)
    }

    req.Header.Set("Content-Type", ContentType)

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("request failed: %w", err)
    }
    defer res.Body.Close()

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        body, _ := io.ReadAll(res.Body)
        return fmt.Errorf("failed to send email: %s", string(body))
    }

    return nil
}
