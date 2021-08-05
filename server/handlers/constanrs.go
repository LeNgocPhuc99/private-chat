package handlers

const (
	// Request Paramyter Validation error/success messages
	UsernameAndPasswordCantBeEmpty = "username and password can't empty"
	UsernameIsNotAvailable         = "username is not available"
	UsernameIsAvailable            = "username is available"
	UserIdIsNotAvailable           = "userID is not available"
	LoginPasswordIsInCorrect       = "your Login Password is incorrect"
	UserRegistrationCompleted      = "user Registration Completed"
	UserLoginCompleted             = "user Login is Completed"
	YouAreNotLoggedIN              = "you are not logged in"
	YouAreLoggedIN                 = "you are logged in"
	UserIsNotRegisteredWithUs      = "this account does not exist in our system"
	UpdateStatusFail               = "update user's status fail"

	// Application response messages
	SuccessfulResponse   = "request completed successfully"
	ServerFailedResponse = "request failed to complete, we are working on it"
	APIWelcomeMessage    = "This is an API for Realtime Private chat application build in GoLang"
)
