// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package request

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest(in *jlexer.Lexer, out *UpdatePassword) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "new_password":
			out.NewPassword = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest(out *jwriter.Writer, in UpdatePassword) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"new_password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.NewPassword))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdatePassword) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdatePassword) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdatePassword) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdatePassword) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest1(in *jlexer.Lexer, out *UpdateInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "username":
			out.Username = string(in.String())
		case "phone":
			out.Phone = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest1(out *jwriter.Writer, in UpdateInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Phone))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest1(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest2(in *jlexer.Lexer, out *UpdateCar) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "car_id":
			out.CarID = uint64(in.Uint64())
		case "model":
			out.Model = string(in.String())
		case "color":
			out.Color = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest2(out *jwriter.Writer, in UpdateCar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"car_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.CarID))
	}
	{
		const prefix string = ",\"model\":"
		out.RawString(prefix)
		out.String(string(in.Model))
	}
	{
		const prefix string = ",\"color\":"
		out.RawString(prefix)
		out.String(string(in.Color))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateCar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateCar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateCar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateCar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest2(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest3(in *jlexer.Lexer, out *Signup) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "username":
			out.Username = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "phone":
			out.Phone = uint64(in.Uint64())
		case "password":
			out.Password = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest3(out *jwriter.Writer, in Signup) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix[1:])
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Phone))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Signup) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Signup) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Signup) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Signup) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest3(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest4(in *jlexer.Lexer, out *SignoutAll) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "token":
			out.Token = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest4(out *jwriter.Writer, in SignoutAll) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SignoutAll) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SignoutAll) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SignoutAll) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SignoutAll) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest4(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest5(in *jlexer.Lexer, out *Signout) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "token":
			out.Token = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest5(out *jwriter.Writer, in Signout) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Signout) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Signout) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Signout) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Signout) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest5(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest6(in *jlexer.Lexer, out *Signin) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest6(out *jwriter.Writer, in Signin) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Signin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Signin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Signin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Signin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest6(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest7(in *jlexer.Lexer, out *Sessions) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "token":
			out.Token = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest7(out *jwriter.Writer, in Sessions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Sessions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Sessions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Sessions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Sessions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest7(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest8(in *jlexer.Lexer, out *Refresh) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "token":
			out.Token = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest8(out *jwriter.Writer, in Refresh) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Refresh) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Refresh) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Refresh) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Refresh) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest8(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest9(in *jlexer.Lexer, out *GetMe) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest9(out *jwriter.Writer, in GetMe) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetMe) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMe) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMe) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMe) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest9(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest10(in *jlexer.Lexer, out *GetCars) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest10(out *jwriter.Writer, in GetCars) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCars) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCars) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCars) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCars) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest10(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest11(in *jlexer.Lexer, out *DeleteUser) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest11(out *jwriter.Writer, in DeleteUser) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeleteUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest11(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest12(in *jlexer.Lexer, out *DeleteCar) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "car_id":
			out.CarID = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest12(out *jwriter.Writer, in DeleteCar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"car_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.CarID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeleteCar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteCar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteCar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteCar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest12(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest13(in *jlexer.Lexer, out *CreateCar) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "plate":
			out.Plate = string(in.String())
		case "model":
			out.Model = string(in.String())
		case "color":
			out.Color = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest13(out *jwriter.Writer, in CreateCar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"plate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Plate))
	}
	{
		const prefix string = ",\"model\":"
		out.RawString(prefix)
		out.String(string(in.Model))
	}
	{
		const prefix string = ",\"color\":"
		out.RawString(prefix)
		out.String(string(in.Color))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateCar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateCar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateCar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateCar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest13(l, v)
}
