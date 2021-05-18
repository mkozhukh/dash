package main

import (
	"errors"
	remote "github.com/mkozhukh/go-remote"
	"log"
	"os/exec"
	"strings"
)

type AdminAPI struct{}

type ExecResponse struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type ConfigResponse struct {
	Type     string            `json:"type"`
	Commands []CommandResponse `json:"commands"`
	Server   string            `json:"server"`
}

type CommandResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Details string `json:"details"`
	Danger  bool   `json:"danger"`
}

type InfoValue struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type InfoResponse struct {
	Type string      `json:"type"`
	Info []InfoValue `json:"info"`
}

func (a AdminAPI) Login(key string) (string, error) {
	for _, x := range Config.Users {
		if x.Key == key {
			out, err := createUserToken(x.Groups)
			if err != nil {
				return "", err
			}

			return string(out), nil
		}
	}

	return "", errors.New("access denied")
}

func (a AdminAPI) GetInfo(key string, cid remote.ConnectionID, hub *remote.Hub) error {
	groups, err := verifyUserToken([]byte(key))
	if err != nil {
		return err
	}

	if a.getConfig(groups, cid, hub) {
		a.getInfo(cid, hub)
	}
	return nil
}

func (a AdminAPI) commandAllowed(c *CommandTarget, groups map[string]bool) bool {
	for _, cGroup := range c.Groups {
		if groups[cGroup] {
			return true
		}
	}

	return false
}

func (a AdminAPI) getConfig(groups map[string]bool, cid remote.ConnectionID, hub *remote.Hub) bool {
	out := ConfigResponse{Type: "config", Server: Config.Server}
	for _, x := range Config.Commands {
		if a.commandAllowed(&x, groups) {
			out.Commands = append(out.Commands, CommandResponse{
				ID:      x.ID,
				Name:    x.Name,
				Details: x.Details,
				Danger:  x.Danger,
			})
		}
	}

	if len(out.Commands) == 0 {
		return false
	}

	hub.Publish("update", out, cid)
	return true
}

func (a AdminAPI) getInfo(cid remote.ConnectionID, hub *remote.Hub) {
	hub.Publish("update", InfoResponse{
		Type: "info",
		Info: a.getInfoValues(),
	}, cid)
}

func (a AdminAPI) Exec(id string, uid string, key string, cid remote.ConnectionID, hub *remote.Hub) error {
	groups, err := verifyUserToken([]byte(key))
	if err != nil {
		return err
	}

	for _, c := range Config.Commands {
		if c.ID == id {
			for _, g := range c.Groups {
				if groups[g] {
					if !a.getExecValues(uid, c.Exec, cid, hub) {
						hub.Publish("update", ExecResponse{
							ID:      uid,
							Type:    "done",
							Status:  "error",
							Message: "Error during command execution",
						}, cid)
					}
					return nil
				}
			}
		}
	}

	hub.Publish("update", ExecResponse{
		ID:      uid,
		Type:    "exec",
		Status:  "error",
		Message: "command not found",
	}, cid)
	return nil
}

func (a AdminAPI) getUser(key string) *User {
	if Config.Users == nil {
		return nil
	}

	for i, u := range Config.Users {
		if u.Key == key {
			return &Config.Users[i]
		}
	}

	return nil
}

func (a AdminAPI) getExecValues(uid string, commands []string, cid remote.ConnectionID, hub *remote.Hub) bool {
	for _, x := range commands {
		out, err := exec.Command("/bin/bash", "-c", x).Output()
		if err != nil {
			log.Printf("ERROR:\n%s\n%s\n\n", x, err.Error())
			hub.Publish("update", ExecResponse{
				ID:      uid,
				Type:    "exec",
				Status:  "error",
				Message: err.Error(),
			}, cid)
			return false
		} else {
			hub.Publish("update", ExecResponse{
				ID:      uid,
				Type:    "exec",
				Status:  "ok",
				Message: string(out),
			}, cid)
		}
	}

	hub.Publish("update", ExecResponse{
		ID:     uid,
		Type:   "done",
		Status: "ok",
	}, cid)
	a.getInfo(cid, hub)

	return true
}

func (a AdminAPI) getInfoValues() []InfoValue {
	res := make([]InfoValue, len(Config.Info))
	for i, x := range Config.Info {
		out, err := exec.Command("/bin/bash", "-c", x.Exec).Output()
		if err != nil {
			log.Printf("ERROR(%s):\n%s\n%s\n\n", x.ID, x.Exec, err.Error())
		}

		res[i] = InfoValue{
			ID:    x.ID,
			Name:  x.Name,
			Value: strings.TrimSpace(string(out)),
		}
	}

	return res
}
