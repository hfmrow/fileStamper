// +build !gtk_3_6,!gtk_3_8,!gtk_3_10

// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// This file includes wrapers for symbols included since GTK 3.12, and
// and should not be included in a build intended to target any older GTK
// versions.  To target an older build, such as 3.10, use
// 'go build -tags gtk_3_10'.  Otherwise, if no build tags are used, GTK 3.12
// is assumed and this file is built.
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

// SetCenterWidget is a wrapper around gtk_box_set_center_widget().
func (a *Box) SetCenterWidget(child IWidget) {
	if child == nil {
		C.gtk_box_set_center_widget(a.native(), nil)
	} else {
		C.gtk_box_set_center_widget(a.native(), child.toWidget())
	}
}

// GetCenterWidget is a wrapper around gtk_box_get_center_widget().
func (a *Box) GetCenterWidget() (IWidget, error) {
	w := C.gtk_box_get_center_widget(a.native())
	if w == nil {
		return nil, nil
	}
	return castWidget(w)
}
