// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type ToolBar struct {
	AssignTo           **walk.ToolBar
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
	Actions            []*walk.Action
	MaxTextRows        int
	Orientation        Orientation
}

func (tb ToolBar) Create(parent walk.Container) (err error) {
	var w *walk.ToolBar
	if tb.Orientation == Vertical {
		w, err = walk.NewVerticalToolBar(parent)
	} else {
		w, err = walk.NewToolBar(parent)
	}
	if err != nil {
		return
	}

	return InitWidget(tb, w, func() error {
		imageList, err := walk.NewImageList(walk.Size{16, 16}, 0)
		if err != nil {
			return err
		}
		w.SetImageList(imageList)

		mtr := tb.MaxTextRows
		if mtr < 1 {
			mtr = 1
		}
		if err := w.SetMaxTextRows(mtr); err != nil {
			return err
		}

		if err := addToActionList(w.Actions(), tb.Actions); err != nil {
			return err
		}

		if tb.AssignTo != nil {
			*tb.AssignTo = w
		}

		return nil
	})
}

func (tb ToolBar) WidgetInfo() (name string, disabled, hidden bool, font *Font, minSize, maxSize Size, stretchFactor, row, rowSpan, column, columnSpan int, contextMenuActions []*walk.Action) {
	return tb.Name, tb.Disabled, tb.Hidden, &tb.Font, tb.MinSize, tb.MaxSize, tb.StretchFactor, tb.Row, tb.RowSpan, tb.Column, tb.ColumnSpan, tb.ContextMenuActions
}
