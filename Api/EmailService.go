package Api

import (
    "GoMailApp_Updated_Fixed/Common"
    "GoMailApp_Updated_Fixed/Proto"
    "context"
)

type EmailServiceServer struct {
    Proto.UnimplementedEmailServiceServer
}

func (s *EmailServiceServer) SendEmail(ctx context.Context, in *Proto.UserInput) (*Proto.EmailResponse, error) {
    err := Common.SendTemplatedEmail(in.Name, in.Email)
    if err != nil {
        return &Proto.EmailResponse{Success: false, Message: err.Error()}, nil
    }
    return &Proto.EmailResponse{Success: true, Message: "Email sent successfully"}, nil
}
