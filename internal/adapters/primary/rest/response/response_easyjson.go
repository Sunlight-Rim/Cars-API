// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package response

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

func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse(in *jlexer.Lexer, out *User) {
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
		case "id":
			out.ID = uint64(in.Uint64())
		case "username":
			out.Username = string(in.String())
		case "email":
			out.Email = string(in.String())
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse1(in *jlexer.Lexer, out *UpdateCar) {
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
		case "car":
			(out.Car).UnmarshalEasyJSON(in)
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse1(out *jwriter.Writer, in UpdateCar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"car\":"
		out.RawString(prefix[1:])
		(in.Car).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateCar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateCar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateCar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateCar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse1(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse2(in *jlexer.Lexer, out *Signup) {
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
		case "id":
			out.ID = uint64(in.Uint64())
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse2(out *jwriter.Writer, in Signup) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Signup) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Signup) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Signup) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Signup) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse2(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse3(in *jlexer.Lexer, out *SignoutAll) {
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
		case "tokens":
			if in.IsNull() {
				in.Skip()
				out.Tokens = nil
			} else {
				in.Delim('[')
				if out.Tokens == nil {
					if !in.IsDelim(']') {
						out.Tokens = make([]string, 0, 4)
					} else {
						out.Tokens = []string{}
					}
				} else {
					out.Tokens = (out.Tokens)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Tokens = append(out.Tokens, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse3(out *jwriter.Writer, in SignoutAll) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tokens\":"
		out.RawString(prefix[1:])
		if in.Tokens == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Tokens {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SignoutAll) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SignoutAll) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SignoutAll) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SignoutAll) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse3(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse4(in *jlexer.Lexer, out *Signout) {
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse4(out *jwriter.Writer, in Signout) {
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
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Signout) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Signout) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Signout) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse4(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse5(in *jlexer.Lexer, out *Signin) {
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse5(out *jwriter.Writer, in Signin) {
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
func (v Signin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Signin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Signin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Signin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse5(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse6(in *jlexer.Lexer, out *Response) {
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
		case "response":
			if m, ok := out.Response.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Response.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Response = in.Interface()
			}
		case "error":
			if m, ok := out.Error.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Error.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Error = in.Interface()
			}
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse6(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"response\":"
		out.RawString(prefix[1:])
		if m, ok := in.Response.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Response.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Response))
		}
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		if m, ok := in.Error.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Error.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Error))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse6(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse7(in *jlexer.Lexer, out *Refresh) {
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse7(out *jwriter.Writer, in Refresh) {
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
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Refresh) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Refresh) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Refresh) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse7(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse8(in *jlexer.Lexer, out *GetMe) {
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
		case "user":
			(out.User).UnmarshalEasyJSON(in)
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse8(out *jwriter.Writer, in GetMe) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix[1:])
		(in.User).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetMe) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMe) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMe) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMe) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse8(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse9(in *jlexer.Lexer, out *GetCars) {
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
		case "cars":
			if in.IsNull() {
				in.Skip()
				out.Cars = nil
			} else {
				in.Delim('[')
				if out.Cars == nil {
					if !in.IsDelim(']') {
						out.Cars = make([]Car, 0, 1)
					} else {
						out.Cars = []Car{}
					}
				} else {
					out.Cars = (out.Cars)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Car
					(v4).UnmarshalEasyJSON(in)
					out.Cars = append(out.Cars, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse9(out *jwriter.Writer, in GetCars) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cars\":"
		out.RawString(prefix[1:])
		if in.Cars == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Cars {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCars) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCars) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCars) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCars) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse9(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse10(in *jlexer.Lexer, out *ErrorResponse) {
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
		case "code":
			out.Code = uint32(in.Uint32())
		case "message":
			out.Message = string(in.String())
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse10(out *jwriter.Writer, in ErrorResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.Code))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse10(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse11(in *jlexer.Lexer, out *DeleteCar) {
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
		case "car":
			(out.Car).UnmarshalEasyJSON(in)
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse11(out *jwriter.Writer, in DeleteCar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"car\":"
		out.RawString(prefix[1:])
		(in.Car).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeleteCar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteCar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteCar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteCar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse11(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse12(in *jlexer.Lexer, out *CreateCar) {
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
		case "id":
			out.ID = uint64(in.Uint64())
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse12(out *jwriter.Writer, in CreateCar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateCar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateCar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateCar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateCar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse12(l, v)
}
func easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse13(in *jlexer.Lexer, out *Car) {
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
		case "id":
			out.ID = uint64(in.Uint64())
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
func easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse13(out *jwriter.Writer, in Car) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"plate\":"
		out.RawString(prefix)
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
func (v Car) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Car) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeCarsInternalAdaptersPrimaryRestResponse13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Car) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Car) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeCarsInternalAdaptersPrimaryRestResponse13(l, v)
}
