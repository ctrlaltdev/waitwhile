package waitwhile

var (
	baseUrl = "https://api.waitwhile.com/v2"
	apiKey  *string
)

func Init(iApiKey *string) {
	apiKey = iApiKey
}
