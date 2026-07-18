package tracker

type UI struct {
	In      Input
	Out     Output
	Tracker *Tracker
}

func (u UI) Run() {
	actions := map[string]UseCase{
		"add":    AddUseCase{},
		"get":    GetUseCase{},
		"delete": DeleteUseCase{},
		"find":   FindUseCase{},
		"update": UpdateUseCase{},
	}

	for {
		u.Out.Out("введите действие:")
		for action := range actions {
			u.Out.Out(action)
		}

		selected := u.In.Get()
		if selected == "exit" {
			break
		}

		action, ok := actions[selected]
		if !ok {
			u.Out.Out("действие не найдено")
			continue
		}

		action.Done(u.In, u.Out, u.Tracker)
	}
}
