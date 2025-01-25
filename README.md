# Golang Geekbot SDK

This is an unofficial SDK for the [Geekbot API](https://geekbot.com/developers/).

## Example Usage

```go
import (
	"fmt"
	"time"

	v1 "github.com/omissis/asmctl/internal/geekbot/v1"
)

var ErrCannotExecuteSuccessfully = fmt.Errorf("cannot execute successfully")

func main() {
	httpClient := &http.Client{}
	apiKey, ok := os.LookupEnv("GEEKBOT_API_KEY")
	if !ok {
		return "", fmt.Errorf("%w: GEEKBOT_API_KEY environment variable is required", ErrCannotExecuteSuccessfully)
	}

	geekbotSDK := v1.NewSDK(httpClient, v1.DefaultBaseURL, apiKey)

	oneWeekAgo := time.Now().AddDate(0, 0, -9).Truncate(24 * time.Hour).Unix()

	reports, err := geekbotSDK.ListReports(v1.ListReportsFilters{After: &oneWeekAgo})
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrCannotExecuteSuccessfully, err)
	}

	fmt.Printf("%+v\n", reports)
}
```

## Implemented Endpoints

- [x] Team > List
- [ ] Report > Create
- [x] Report > List
- [x] Standup > List
- [ ] Standup > Read
- [ ] Standup > Create
- [ ] Standup > Update
- [ ] Standup > Replace
- [ ] Standup > Duplicate
- [ ] Standup > Delete
- [ ] Standup > Start
