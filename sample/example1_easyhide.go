package sample

// Code generated by easyhide. DO NOT EDIT.

//go:generate easyjson

import (
	json "encoding/json"

	easyhide "github.com/denis96z/go-easyhide/pkg/easyhide"
)

//easyjson:json
type tp1f93603db53bfad5c92390f735d0cbb8617b4ab8214ae91c5664a3d1e9b009c8 T1

func (v T1) EasyHide() ([]byte, error) {
	xv := tp1f93603db53bfad5c92390f735d0cbb8617b4ab8214ae91c5664a3d1e9b009c8(v)
	if easyhide.HideData {
		xv.A2 = easyhide.HiddenMarker
		xv.A3 = easyhide.HiddenMarker + xv.A3[len(xv.A3)/2:]
		xv.A4 = xv.A4[:len(xv.A4)/2] + easyhide.HiddenMarker
		if xv.A5 != "" {
			xv.A5 = easyhide.HiddenMarker
		}
		if xv.A6 != "" {
			xv.A6 = easyhide.HiddenMarker + xv.A6[len(xv.A6)/2:]
		}
		if xv.A7 != "" {
			xv.A7 = xv.A7[:len(xv.A7)/2] + easyhide.HiddenMarker
		}
	}
	return json.Marshal(xv)
}

//easyjson:json
type tp5dd67f7fb9c529cb28245800137482c9c2fdff9b7d22f54cd3bfa90c59b78481 T3

func (v T3) EasyHide() ([]byte, error) {
	xv := tp5dd67f7fb9c529cb28245800137482c9c2fdff9b7d22f54cd3bfa90c59b78481(v)
	if easyhide.HideData {
		xv.C1.D1 = easyhide.HiddenMarker
	}
	return json.Marshal(xv)
}
