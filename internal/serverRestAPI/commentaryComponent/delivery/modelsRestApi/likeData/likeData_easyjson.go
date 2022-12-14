// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package likeData

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

func easyjson61671895Decode20222GoToTeamInternalServerRestAPICommentaryComponentDeliveryModelsRestApiLikeData(in *jlexer.Lexer, out *LikeData) {
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
			out.Id = int(in.Int())
		case "sign":
			out.Sign = int(in.Int())
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
func easyjson61671895Encode20222GoToTeamInternalServerRestAPICommentaryComponentDeliveryModelsRestApiLikeData(out *jwriter.Writer, in LikeData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"sign\":"
		out.RawString(prefix)
		out.Int(int(in.Sign))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LikeData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson61671895Encode20222GoToTeamInternalServerRestAPICommentaryComponentDeliveryModelsRestApiLikeData(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LikeData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson61671895Encode20222GoToTeamInternalServerRestAPICommentaryComponentDeliveryModelsRestApiLikeData(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LikeData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson61671895Decode20222GoToTeamInternalServerRestAPICommentaryComponentDeliveryModelsRestApiLikeData(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LikeData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson61671895Decode20222GoToTeamInternalServerRestAPICommentaryComponentDeliveryModelsRestApiLikeData(l, v)
}
