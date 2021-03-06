// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type Composite struct {
	AssignTo           **walk.Composite
	Name               string
	Disabled           bool
	Hidden             bool
	Font               Font
	MinSize            Size
	MaxSize            Size
	StretchFactor      int
	Row                int
	RowSpan            int
	Column             int
	ColumnSpan         int
	ContextMenuActions []*walk.Action
	Layout             Layout
	Children           []Widget
}

func (c Composite) Create(parent walk.Container) error {
	w, err := walk.NewComposite(parent)
	if err != nil {
		return err
	}

	return InitWidget(c, w, func() error {
		if c.AssignTo != nil {
			*c.AssignTo = w
		}

		return nil
	})
}

func (c Composite) WidgetInfo() (name string, disabled, hidden bool, font *Font, minSize, maxSize Size, stretchFactor, row, rowSpan, column, columnSpan int, contextMenuActions []*walk.Action) {
	return c.Name, c.Disabled, c.Hidden, &c.Font, c.MinSize, c.MaxSize, c.StretchFactor, c.Row, c.RowSpan, c.Column, c.ColumnSpan, c.ContextMenuActions
}

func (c Composite) ContainerInfo() (Layout, []Widget) {
	return c.Layout, c.Children
}
