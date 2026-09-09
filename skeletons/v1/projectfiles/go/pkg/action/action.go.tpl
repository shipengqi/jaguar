package action

import (
	"context"
	"errors"
)

// Action represents a unit of work with an optional lifecycle.
type Action struct {
	Name              string
	PersistentPreRun  func(act *Action) error
	PreRun            func(act *Action) error
	Run               func(act *Action) error
	PostRun           func(act *Action) error
	PersistentPostRun func(act *Action) error
	Executable        func(act *Action) bool

	actions []*Action
	parent  *Action
	ctx     context.Context
}

func (a *Action) Context() context.Context { return a.ctx }
func (a *Action) HasParent() bool          { return a.parent != nil }
func (a *Action) Parent() *Action          { return a.parent }
func (a *Action) Runnable() bool           { return a.Run != nil }
func (a *Action) HasSubActions() bool      { return len(a.actions) > 0 }
func (a *Action) Actions() []*Action       { return a.actions }

func (a *Action) Root() *Action {
	if a.HasParent() {
		return a.Parent().Root()
	}
	return a
}

func (a *Action) AddAction(actions ...*Action) error {
	for _, x := range actions {
		if x == a {
			return errors.New("action can't be a child of itself")
		}
		x.parent = a
		a.actions = append(a.actions, x)
	}
	return nil
}

func (a *Action) RemoveAction(actions ...*Action) {
	var keep []*Action
outer:
	for _, act := range a.actions {
		for _, rem := range actions {
			if act == rem {
				act.parent = nil
				continue outer
			}
		}
		keep = append(keep, act)
	}
	a.actions = keep
}

func (a *Action) ExecuteContext(ctx context.Context) error {
	a.ctx = ctx
	return a.Execute()
}

func (a *Action) Execute() error {
	if a.ctx == nil {
		a.ctx = context.Background()
	}
	if a.HasParent() {
		return a.Root().Execute()
	}
	target := a.Find()
	if target == nil {
		target = a
	}
	if target.ctx == nil {
		target.ctx = a.ctx
	}
	return target.execute()
}

func (a *Action) Find() *Action {
	if a.Executable != nil && a.Executable(a) {
		return a
	}
	for _, v := range a.actions {
		if t := v.Find(); t != nil {
			return t
		}
	}
	return nil
}

func (a *Action) execute() error {
	if a == nil || !a.Runnable() {
		return nil
	}
	for p := a; p != nil; p = p.Parent() {
		if p.PersistentPreRun != nil {
			if err := p.PersistentPreRun(a); err != nil {
				return err
			}
			break
		}
	}
	if a.PreRun != nil {
		if err := a.PreRun(a); err != nil {
			return err
		}
	}
	if a.Run != nil {
		if err := a.Run(a); err != nil {
			return err
		}
	}
	if a.PostRun != nil {
		if err := a.PostRun(a); err != nil {
			return err
		}
	}
	for p := a; p != nil; p = p.Parent() {
		if p.PersistentPostRun != nil {
			if err := p.PersistentPostRun(a); err != nil {
				return err
			}
			break
		}
	}
	return nil
}
