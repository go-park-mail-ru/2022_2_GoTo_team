// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package modelsRestApi

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

func easyjson836cf2d1Decode20222GoToTeamInternalServerRestAPIUserComponentDeliveryModelsRestApi(in *jlexer.Lexer, out *UserInfo) {
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
		case "registration_date":
			out.RegistrationDate = string(in.String())
		case "subscribers_count":
			out.SubscribersCount = int(in.Int())
		case "rating":
			out.Rating = int(in.Int())
		case "subscribed":
			out.Subscribed = bool(in.Bool())
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
func easyjson836cf2d1Encode20222GoToTeamInternalServerRestAPIUserComponentDeliveryModelsRestApi(out *jwriter.Writer, in UserInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix[1:])
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"registration_date\":"
		out.RawString(prefix)
		out.String(string(in.RegistrationDate))
	}
	{
		const prefix string = ",\"subscribers_count\":"
		out.RawString(prefix)
		out.Int(int(in.SubscribersCount))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Int(int(in.Rating))
	}
	{
		const prefix string = ",\"subscribed\":"
		out.RawString(prefix)
		out.Bool(bool(in.Subscribed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson836cf2d1Encode20222GoToTeamInternalServerRestAPIUserComponentDeliveryModelsRestApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson836cf2d1Encode20222GoToTeamInternalServerRestAPIUserComponentDeliveryModelsRestApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson836cf2d1Decode20222GoToTeamInternalServerRestAPIUserComponentDeliveryModelsRestApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson836cf2d1Decode20222GoToTeamInternalServerRestAPIUserComponentDeliveryModelsRestApi(l, v)
}
