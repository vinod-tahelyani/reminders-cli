package client

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type BackendHTTPClient interface {
	Create(title string, message string, duration time.Duration) ([]byte, error)
	Edit(id int, title string, message string, duration time.Duration) ([]byte, error)
	Fetch(id int) ([]byte, error)
	Delete(id int) (error)
	Healthy(host string) bool
}

type Switch struct {
	client BackendHTTPClient 
	backendAPIURL string
	commands map[string]func() func(string) error
}

func NewSwitch(uri string) Switch {
	httpClient := NewHTTPClient(uri)
	s :=  Switch {
		client: httpClient,
		backendAPIURL: uri,
	}
	s.commands = map[string]func() func(string) error {
		"create": s.create,
		"fetch": s.fetch,
		"edit": s.edit,
		"delete": s.delete,
		"health": s.health,
	}
	return s
}

func (s Switch) Help() {
	var help string
	for name := range s.commands {
		help += name + "\t --help\n"
	}
	fmt.Printf("Usage of %s:\n<command> [<args>]\n%s", os.Args[0], help)
}

func (s Switch) Switch() error {
	cmdName := os.Args[1]
	cmd, found := s.commands[cmdName]
	if !found {
		return fmt.Errorf("command '%s' not found, try --help", cmdName)
	}
	return cmd()(cmdName)
}

func (s Switch) create() func(string) error {
	return func(cmd string) error {
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(*createCmd)
		
		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(*createCmd); err != nil {
			return err
		}

		res, err := s.client.Create(t, m, d)
		if err != nil {
			return wrapError("error while creating reminder", err)
		}
		fmt.Printf("Reminder Created \n%s", string(res))
		return nil
	}
}

func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		id := s.idFlag(*fetchCmd)
		
		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(*fetchCmd); err != nil {
			return err
		}

		res, err := s.client.Fetch(id)
		if err != nil {
			return wrapError("error in fetching reminder with id '" + string(id) + "' ", err)
		}
		fmt.Printf("Reminder found: \n%s", res)
		return nil
	}
}

func (s Switch) edit() func(string) error {
	return func(cmd string) error {
		editCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(*editCmd)
		id := s.idFlag(*editCmd)
		
		if err := s.checkArgs(4); err != nil {
			return err
		}

		if err := s.parseCmd(*editCmd); err != nil {
			return err
		}

		res, err := s.client.Edit(id, t, m, d)
		if err != nil {
			return wrapError("error in editing reminder with id '" + string(id) + "' ", err)
		}
		fmt.Printf("Reminder updated: \n%s", res)
		return nil
	}
}

func (s Switch) delete() func(string) error {
	return func(cmd string) error {
		deleteCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		id := s.idFlag(*deleteCmd)
		
		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(*deleteCmd); err != nil {
			return err
		}

		res, err := s.client.Fetch(id)
		if err != nil {
			return wrapError("error in deleting reminder with id '" + string(id) + "' ", err)
		}
		fmt.Printf("Reminder deleted: \n%s", res)
		return nil
	}
}

func (s Switch) health() func(string) error {
	return func(cmd string) error {
		healthCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		
		if err := s.checkArgs(0); err != nil {
			return err
		}

		if err := s.parseCmd(*healthCmd); err != nil {
			return err
		}

		healthy := s.client.Healthy(s.backendAPIURL)
		if healthy {
			fmt.Printf("Server %s healthy and ok!\n", s.backendAPIURL)
		} else {
			fmt.Printf("Server %s is NOT healthy and ok!\n", s.backendAPIURL)
		}
		return nil
	}
}

func (s Switch) reminderFlags(f flag.FlagSet) (string, string, time.Duration) {
	t, m, d := "", "", time.Duration(0)

	f.StringVar(&t, "title", "", "Reminder title")
	f.StringVar(&t, "t", "", "Reminder title")
	f.StringVar(&m, "message", "", "Reminder message")
	f.StringVar(&m, "title", "", "Reminder message") 
	f.DurationVar(&d, "duration", 0, "Reminder duration")
	f.DurationVar(&d, "d", 0, "Reminder duration")
	return t, m, d
}

func (s Switch) idFlag(f flag.FlagSet) (int) {
	id := 0
	f.IntVar(&id, "id", 0, "Reminder id")
	f.IntVar(&id, "i", 0, "Reminder id")
	return id
}


func (s Switch) checkArgs(minArgs int) error {
	return nil
}

func (s Switch) parseCmd(cmd flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return wrapError("could not parse '" + cmd.Name() + "' flag", err)
	}
	return nil
}