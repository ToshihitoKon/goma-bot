package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/slack-go/slack/slackevents"
)

func HandlerSlack(w http.ResponseWriter, r *http.Request) {
	verificationToken := os.Getenv("GOMABOT_SLACK_VERIFICATION_TOKEN")
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	eventsAPIEvent, e := slackevents.ParseEvent(
		json.RawMessage(body),
		slackevents.OptionVerifyToken(
			&slackevents.TokenComparator{
				VerificationToken: verificationToken,
			},
		),
	)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(e.Error()))
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}
}
