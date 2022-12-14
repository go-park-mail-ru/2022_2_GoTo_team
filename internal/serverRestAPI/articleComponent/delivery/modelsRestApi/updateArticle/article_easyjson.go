// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package updateArticle

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

func easyjson6de889b8Decode20222GoToTeamInternalServerRestAPIArticleComponentDeliveryModelsRestApiUpdateArticle(in *jlexer.Lexer, out *Article) {
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
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "category":
			out.Category = string(in.String())
		case "content":
			out.Content = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Tags = append(out.Tags, v1)
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
func easyjson6de889b8Encode20222GoToTeamInternalServerRestAPIArticleComponentDeliveryModelsRestApiUpdateArticle(out *jwriter.Writer, in Article) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Tags {
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
func (v Article) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6de889b8Encode20222GoToTeamInternalServerRestAPIArticleComponentDeliveryModelsRestApiUpdateArticle(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Article) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6de889b8Encode20222GoToTeamInternalServerRestAPIArticleComponentDeliveryModelsRestApiUpdateArticle(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Article) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6de889b8Decode20222GoToTeamInternalServerRestAPIArticleComponentDeliveryModelsRestApiUpdateArticle(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Article) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6de889b8Decode20222GoToTeamInternalServerRestAPIArticleComponentDeliveryModelsRestApiUpdateArticle(l, v)
}
