package repl

import (
	"charm/hub"
	"charm/readline"
	"charm/text"

	"io"
	"strings"
)

func Start(hub *hub.Hub, in io.Reader, out io.Writer) {
	rline := readline.NewInstance()
	for {
		rline.SetPrompt(makePrompt(hub))
		line, err := rline.Readline()
		if err != nil {
			hub.WriteError(err.Error())
			return
		}

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		quitCharm := hub.Do(line)
		if quitCharm {
			break
		}
	}
}

func makePrompt (hub *hub.Hub) string {
	if hub.GetCurrentServiceName() == "" {
		return text.PROMPT
	}
	return hub.GetCurrentServiceName() + " " + text.PROMPT
}