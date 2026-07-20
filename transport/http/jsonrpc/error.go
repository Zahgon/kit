package jsonrpc

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e Error) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

const (
	ParseError int = -32700

	InvalidRequestError int = -32600

	MethodNotFoundError int = -32601

	InvalidParamsError int = -32602

	InternalError int = -32603
)

var errorMessage = map[int]string{
	ParseError:          "An error occurred on the server while parsing the JSON text.",
	InvalidRequestError: "The JSON sent is not a valid Request object.",
	MethodNotFoundError: "The method does not exist / is not available.",
	InvalidParamsError:  "Invalid method parameter(s).",
	InternalError:       "Internal JSON-RPC error.",
}

func ErrorMessage(code int) string { _ = "STUB: not implemented"; return "" }

type parseError string

func (e parseError) Error() string { _ = "STUB: not implemented"; return "" }

func (e parseError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

type invalidRequestError string

func (e invalidRequestError) Error() string { _ = "STUB: not implemented"; return "" }

func (e invalidRequestError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

type methodNotFoundError string

func (e methodNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e methodNotFoundError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

type invalidParamsError string

func (e invalidParamsError) Error() string { _ = "STUB: not implemented"; return "" }

func (e invalidParamsError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

type internalError string

func (e internalError) Error() string { _ = "STUB: not implemented"; return "" }

func (e internalError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }
