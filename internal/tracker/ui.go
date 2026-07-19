package tracker

import "context"

type UI struct {
	In    Input
	Out   Output
	Store Store
}

func (u UI) Run(ctx context.Context) error {
	actions := map[string]UseCase{
		"add":    AddUseCase{},
		"get":    GetUseCase{},
		"delete": DeleteUseCase{},
		"find":   FindUseCase{},
		"update": UpdateUseCase{},
	}

	for {
		u.Out.Out("введите действие (add, get, delete, find, update, exit):")
		selected := u.In.Get()

		if selected == "exit" {
			break
		}

		action, ok := actions[selected]
		if !ok {
			u.Out.Out("действие не найдено")
			continue
		}

		if err := action.Done(ctx, u.In, u.Out, u.Store); err != nil {
			return err
		}
	}

	return nil
}
