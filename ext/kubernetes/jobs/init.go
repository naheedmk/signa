package jobs

import (
	"fmt"

	"github.com/signavio/signa/pkg/bot"
)

const (
	invalidAmountOfParams = "Invalid amount of parameters"
	jobNotFound           = "Job not found."
	permissionDenied      = "You are not allowed to execute this operation. :sweat_smile:"
	errorMessage          = "Something went wrong"
	jobOutputNotFound     = "Job executed but output not found."
)

func init() {
	bot.RegisterCommand(
		"run",
		"Run kubernetes jobs.",
		"<job-name>",
		Run,
	)
}

func Run(c *bot.Cmd) (string, error) {
	if len(c.Args) < 1 {
		return invalidAmountOfParams, nil
	}

	job := bot.Cfg().FindJob(c.Args[0])
	if job == nil {
		return jobNotFound, nil
	}

	// TODO: Implement check in a global level.
	username := c.User.Nick
	if bot.Cfg().IsSuperuser(username) || job.IsExecUser(username) {
		var j *Job
		if len(c.Args) == 2 {
			j = NewJob(job, c.Args[1])
		} else {
			j = NewJob(job, "")
		}

		err := j.parseImageTag()
		if err != nil {
			return errorMessage, err
		}

		status := make(chan string)
		go j.exec(status)

		for {
			if current := <-status; current != "" {
				return fmt.Sprintf("```%s```", current), nil
			}
		}
	} else {
		return permissionDenied, nil
	}

	return jobOutputNotFound, nil
}
