package headers

type Header struct {
	Key string
	Val string
}

var hdrs [4]Header = [4]Header{
	{
		Key: "Content-Type",
		Val: "text/event-stream",
	},
	{
		Key: "Cache-Control",
		Val: "no-cache",
	},
	{
		Key: "Connection",
		Val: "keep-alive",
	},
	{
		Key: "Transfer-Encoding",
		Val: "chunked",
	},
}

// Returns the necessary headers to add to responses to a subscription
func GetSSEHeaders() [4]Header {
	return hdrs
}
