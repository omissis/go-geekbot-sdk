package v1

import "time"

type Team struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

type User struct {
	ID         string `json:"id"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Realname   string `json:"realname"`
	ProfileImg string `json:"profile_img"`
}

type Standup struct {
	ID                 int               `json:"id"`
	Name               string            `json:"name"`
	Time               string            `json:"time"`
	WaitTime           int               `json:"wait_time"`
	Timezone           string            `json:"timezone"`
	Days               []any             `json:"days"`
	Channel            string            `json:"channel"`
	ChannelReady       bool              `json:"channel_ready"`
	Questions          []StandupQuestion `json:"questions"`
	Users              []User            `json:"users"`
	UsersTotal         int               `json:"users_total"`
	Webhooks           []any             `json:"webhooks"`
	Master             string            `json:"master"`
	SyncChannelMembers bool              `json:"sync_channel_members"`
	SyncChannelReady   bool              `json:"sync_channel_ready"`
	SyncChannel        any               `json:"sync_channel"`
}

type StandupQuestion struct {
	ID            int        `json:"id"`
	Color         string     `json:"color"`
	Text          string     `json:"text"`
	Schedule      *time.Time `json:"schedule"`
	AnswerType    string     `json:"answer_type"`
	AnswerChoices []any      `json:"answer_choices"`
	HasAnswers    bool       `json:"hasAnswers"` //nolint:tagliatelle // the api is inconsistent, sucks to be me.
	IsRandom      bool       `json:"is_random"`
	RandomTexts   []any      `json:"random_texts"`
	PrefilledBy   *int       `json:"prefilled_by"`
	TextID        int        `json:"text_id"`
	Preconditions []any      `json:"preconditions"`
	Label         string     `json:"label"`
}

type Report struct {
	ID              int              `json:"id"`
	SlackTS         string           `json:"slack_ts"`
	StandupID       int              `json:"standup_id"`
	Timestamp       int              `json:"timestamp"`
	Channel         string           `json:"channel"`
	IsAnonymous     bool             `json:"is_anonymous"`
	BroadcastThread bool             `json:"broadcast_thread"`
	IsConfidential  bool             `json:"is_confidential"`
	Member          User             `json:"member"`
	Questions       []ReportQuestion `json:"questions"`
}

type ReportQuestion struct {
	ID         int           `json:"id"`
	Question   string        `json:"question"`
	QuestionID int           `json:"question_id"`
	Color      string        `json:"color"`
	Answer     string        `json:"answer"`
	Images     []interface{} `json:"images"`
}
