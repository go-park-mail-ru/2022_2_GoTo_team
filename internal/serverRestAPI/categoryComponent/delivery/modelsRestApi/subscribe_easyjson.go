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

func easyjsonFc450ad2Decode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(in *jlexer.Lexer, out *Subscribe) {
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
		case "category_name":
			out.CategoryName = string(in.String())
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
func easyjsonFc450ad2Encode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(out *jwriter.Writer, in Subscribe) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"category_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.CategoryName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Subscribe) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFc450ad2Encode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Subscribe) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFc450ad2Encode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Subscribe) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFc450ad2Decode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Subscribe) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFc450ad2Decode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(l, v)
}
