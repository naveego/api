package errors

// Pipeline errors start at 2000
var (
	GetConnectorError         = 5002000
	GetActivityNodeError      = 5002001
	ConnectorCompilationError = 5002002
	TopicToInputMismatch      = 5002003
	InvalidInputStreamID      = 5002004
	PipelineActivityRunError  = 5002005
)
