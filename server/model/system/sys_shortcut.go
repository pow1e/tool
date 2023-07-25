package system

import "errors"

type Shortcut struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"` // 描述
	Directives  string `json:"directives"`  // 指令
	Type        string `json:"type"`        // 0普通用户 1管理员
}

type ShortcutOption func(*Shortcut) error

func WithShorCutName() ShortcutOption {
	return func(shortcut *Shortcut) error {
		if shortcut.Name == "" {
			return errors.New("快捷方式的名称不能为空")
		}
		return nil
	}
}

func WithShortCutDescription() ShortcutOption {
	return func(shortcut *Shortcut) error {
		if shortcut.Description == "" {
			return errors.New("快捷方式的描述不能为空")
		}
		return nil
	}
}

func WithShortCutDirectives() ShortcutOption {
	return func(shortcut *Shortcut) error {
		if shortcut.Directives == "" {
			return errors.New("快捷方式的指令不能为空")
		}
		return nil
	}
}

func WithShortCutType() ShortcutOption {
	return func(shortcut *Shortcut) error {
		if shortcut.Type != "0" && shortcut.Type != "1" {
			return errors.New("快捷方式的类型只能为0或为1")
		}
		return nil
	}
}

func WithShortCutId() ShortcutOption {
	return func(shortcut *Shortcut) error {
		if shortcut.Id == "" {
			return errors.New("快捷方式的id不能为空")
		}
		return nil
	}
}

// Validate 校验参数，不传则校验全部参数(包括id)
func (s *Shortcut) Validate(options ...ShortcutOption) error {
	if len(options) == 0 {
		options = append(options, WithShortCutId(), WithShorCutName(), WithShortCutDescription(), WithShortCutDirectives(), WithShortCutType())
	}
	for _, option := range options {
		if err := option(s); err != nil {
			return err
		}
	}
	return nil
}

// ValidatedAll 参数校验
func (s *Shortcut) ValidatedAll() error {
	if s.Name == "" {
		return errors.New("快捷方式名称不能为空")
	}
	if s.Description == "" {
		return errors.New("快捷方式描述不能为空")
	}
	if s.Directives == "" {
		return errors.New("快捷方式的指令不能为空")
	}
	if s.Type != "0" && s.Type != "1" {
		return errors.New("快捷方式的类型只能为0或为1")
	}
	return nil
}
