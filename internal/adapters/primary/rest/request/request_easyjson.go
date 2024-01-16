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

func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest(in *jlexer.Lexer, out *Signup) {
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
		case "address":
			out.Address = string(in.String())
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
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest(out *jwriter.Writer, in Signup) {
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
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
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
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Signup) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Signup) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Signup) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest1(in *jlexer.Lexer, out *Signin) {
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
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest1(out *jwriter.Writer, in Signin) {
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
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Signin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Signin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Signin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest1(l, v)
}
func easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest2(in *jlexer.Lexer, out *GetMe) {
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
		case "Token":
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
func easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest2(out *jwriter.Writer, in GetMe) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetMe) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMe) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeCarsInternalAdaptersPrimaryRestRequest2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMe) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMe) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeCarsInternalAdaptersPrimaryRestRequest2(l, v)
}
