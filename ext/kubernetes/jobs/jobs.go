package jobs

import (
	"github.com/signavio/signa/pkg/bot"
)

type Job struct {
	*bot.Job
}

func NewJob(j *bot.Job) *Job {
	return &Job{j}
}

func (j *Job) createJob() (string, error) {
	cmd := NewCommand([]string{"create", "-f", j.Config, "-n", j.Namespace})
	output, err := cmd.Exec()
	if err != nil {
		return "", err
	}
	return output, nil
}

func (j *Job) getJobPods() (string, error) {
	cmd := NewCommand([]string{
		"get",
		"po",
		"--show-all",
		"--selector=job-name=" + j.Name,
		"--output=jsonpath={.items..metadata.name}",
		"-n",
		j.Namespace,
	})
	output, err := cmd.Exec()
	if err != nil {
		return "", err
	}
	return output, nil
}

func (j *Job) getJobLogs(pods string) (string, error) {
	cmd := NewCommand([]string{"logs", pods, "-n", j.Namespace})
	output, err := cmd.Exec()
	if err != nil {
		return "", err
	}
	return output, nil
}

func (j *Job) deleteJob() (string, error) {
	cmd := NewCommand([]string{"delete", "-f", j.Config, "-n", j.Namespace})
	output, err := cmd.Exec()
	if err != nil {
		return "", err
	}
	return output, nil
}
