/*
Sample code for vulnerable type: Improper Access Control: Email Content Injection
CWE : CWE-284
Description : Improper Access Control
*/
package main

import (
    "net/http"
    "net/smtp"
)

func mail(w http.ResponseWriter, r *http.Request) {
    // Assuming email is a query parameter in the request
    email := r.URL.Query().Get("email")
    
    // Replace this with your implementation to get the reset token for the given email
    token := getUserSecretResetToken(email)
    
    // Assuming the host is derived from the request header
    host := r.Host
    
    // Construct the email body
    body := "Click to reset password: " + host + "/reset/" + token   //source
    
    // Replace these placeholders with your SMTP server configurations
    smtpServer := "smtp.example.com"
    smtpPort := "587"
    smtpUsername := "your-smtp-username"
    smtpPassword := "your-smtp-password"
    smtpFrom := "from@example.com"
    
    // Set up authentication credentials
    auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
    
    // Send the email
    err := smtp.SendMail(smtpServer+":"+smtpPort, auth, smtpFrom, []string{email}, []byte(body))  //sink
    if err != nil {
        // Handle error
        http.Error(w, "Failed to send email", http.StatusInternalServerError)
        return
    }
    
    // Email sent successfully
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Email sent successfully"))
}

func getUserSecretResetToken(email string) string {
    // This is a placeholder function. Replace it with your implementation
    // to retrieve the reset token for the given email.
    // Example: Query the database or generate a random token
    return "reset-token-placeholder"
}

func main() {
    http.HandleFunc("/mail", mail)
    http.ListenAndServe(":8080", nil)
}
