// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package handlers

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

func easyjson3073ac56DecodeGithubComKihamoBoggartComponentsBoggartInternalHandlers(in *jlexer.Lexer, out *deviceHandlerResponseSuccess) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			if m, ok := out.Data.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Data.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Data = in.Interface()
			}
		case "result":
			out.Result = string(in.String())
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
func easyjson3073ac56EncodeGithubComKihamoBoggartComponentsBoggartInternalHandlers(out *jwriter.Writer, in deviceHandlerResponseSuccess) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Data != nil {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Data.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Data.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Data))
		}
	}
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Result))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v deviceHandlerResponseSuccess) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3073ac56EncodeGithubComKihamoBoggartComponentsBoggartInternalHandlers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v deviceHandlerResponseSuccess) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3073ac56EncodeGithubComKihamoBoggartComponentsBoggartInternalHandlers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *deviceHandlerResponseSuccess) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3073ac56DecodeGithubComKihamoBoggartComponentsBoggartInternalHandlers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *deviceHandlerResponseSuccess) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3073ac56DecodeGithubComKihamoBoggartComponentsBoggartInternalHandlers(l, v)
}
func easyjson3073ac56DecodeGithubComKihamoBoggartComponentsBoggartInternalHandlers1(in *jlexer.Lexer, out *deviceHandlerResponseFailed) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			out.Result = string(in.String())
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
func easyjson3073ac56EncodeGithubComKihamoBoggartComponentsBoggartInternalHandlers1(out *jwriter.Writer, in deviceHandlerResponseFailed) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Result))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v deviceHandlerResponseFailed) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3073ac56EncodeGithubComKihamoBoggartComponentsBoggartInternalHandlers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v deviceHandlerResponseFailed) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3073ac56EncodeGithubComKihamoBoggartComponentsBoggartInternalHandlers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *deviceHandlerResponseFailed) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3073ac56DecodeGithubComKihamoBoggartComponentsBoggartInternalHandlers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *deviceHandlerResponseFailed) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3073ac56DecodeGithubComKihamoBoggartComponentsBoggartInternalHandlers1(l, v)
}
