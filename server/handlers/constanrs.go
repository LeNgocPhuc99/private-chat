package handlers

const (
	// Request Paramyter Validation error/success messages
	UsernameCantBeEmpty            = "username can't be empty"
	UsernameIsAvailable            = "username is available."
	UsernameIsNotAvailable         = "username is not available"
	PasswordCantBeEmpty            = "password can't be empty"
	UsernameAndPasswordCantBeEmpty = "username and Password can't be empty"
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
