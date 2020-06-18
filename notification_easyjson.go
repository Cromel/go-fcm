// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package fcm

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

func easyjson9806e1DecodeGithubComHumansNetFcm(in *jlexer.Lexer, out *Message) {
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
		case "name":
			out.Name = string(in.String())
		case "data":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Data = make(map[string]string)
				} else {
					out.Data = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.Data)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "notification":
			if in.IsNull() {
				in.Skip()
				out.Notification = nil
			} else {
				if out.Notification == nil {
					out.Notification = new(Notification)
				}
				easyjson9806e1DecodeGithubComHumansNetFcm1(in, out.Notification)
			}
		case "android":
			if in.IsNull() {
				in.Skip()
				out.Android = nil
			} else {
				if out.Android == nil {
					out.Android = new(AndroidConfig)
				}
				easyjson9806e1DecodeGithubComHumansNetFcm2(in, out.Android)
			}
		case "token":
			out.Token = string(in.String())
		case "topic":
			out.Topic = string(in.String())
		case "condition":
			out.Condition = string(in.String())
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
func easyjson9806e1EncodeGithubComHumansNetFcm(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if len(in.Data) != 0 {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Data {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				out.String(string(v2Value))
			}
			out.RawByte('}')
		}
	}
	if in.Notification != nil {
		const prefix string = ",\"notification\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9806e1EncodeGithubComHumansNetFcm1(out, *in.Notification)
	}
	if in.Android != nil {
		const prefix string = ",\"android\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9806e1EncodeGithubComHumansNetFcm2(out, *in.Android)
	}
	if in.Token != "" {
		const prefix string = ",\"token\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Token))
	}
	if in.Topic != "" {
		const prefix string = ",\"topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Topic))
	}
	if in.Condition != "" {
		const prefix string = ",\"condition\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Condition))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9806e1EncodeGithubComHumansNetFcm(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9806e1EncodeGithubComHumansNetFcm(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9806e1DecodeGithubComHumansNetFcm(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9806e1DecodeGithubComHumansNetFcm(l, v)
}
func easyjson9806e1DecodeGithubComHumansNetFcm2(in *jlexer.Lexer, out *AndroidConfig) {
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
		case "collapse_key":
			out.CollapseKey = string(in.String())
		case "priority":
			out.Priority = AndroidMessagePriority(in.String())
		case "ttl":
			out.Ttl = string(in.String())
		case "restricted_package_name":
			out.RestrictedPackageName = string(in.String())
		case "data":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Data = make(map[string]string)
				} else {
					out.Data = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 string
					v3 = string(in.String())
					(out.Data)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "notification":
			if in.IsNull() {
				in.Skip()
				out.Notification = nil
			} else {
				if out.Notification == nil {
					out.Notification = new(AndroidNotification)
				}
				easyjson9806e1DecodeGithubComHumansNetFcm3(in, out.Notification)
			}
		case "fcm_options":
			if in.IsNull() {
				in.Skip()
				out.FCMOptions = nil
			} else {
				if out.FCMOptions == nil {
					out.FCMOptions = new(AndroidFCMOptions)
				}
				easyjson9806e1DecodeGithubComHumansNetFcm4(in, out.FCMOptions)
			}
		case "direct_boot_ok":
			out.DirectBootOk = bool(in.Bool())
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
func easyjson9806e1EncodeGithubComHumansNetFcm2(out *jwriter.Writer, in AndroidConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CollapseKey != "" {
		const prefix string = ",\"collapse_key\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.CollapseKey))
	}
	if in.Priority != "" {
		const prefix string = ",\"priority\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Priority))
	}
	if in.Ttl != "" {
		const prefix string = ",\"ttl\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ttl))
	}
	if in.RestrictedPackageName != "" {
		const prefix string = ",\"restricted_package_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RestrictedPackageName))
	}
	if len(in.Data) != 0 {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.Data {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				out.String(string(v4Value))
			}
			out.RawByte('}')
		}
	}
	if in.Notification != nil {
		const prefix string = ",\"notification\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9806e1EncodeGithubComHumansNetFcm3(out, *in.Notification)
	}
	if in.FCMOptions != nil {
		const prefix string = ",\"fcm_options\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9806e1EncodeGithubComHumansNetFcm4(out, *in.FCMOptions)
	}
	if in.DirectBootOk {
		const prefix string = ",\"direct_boot_ok\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.DirectBootOk))
	}
	out.RawByte('}')
}
func easyjson9806e1DecodeGithubComHumansNetFcm4(in *jlexer.Lexer, out *AndroidFCMOptions) {
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
		case "analytics_label":
			out.AnalyticsLabel = string(in.String())
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
func easyjson9806e1EncodeGithubComHumansNetFcm4(out *jwriter.Writer, in AndroidFCMOptions) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AnalyticsLabel != "" {
		const prefix string = ",\"analytics_label\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.AnalyticsLabel))
	}
	out.RawByte('}')
}
func easyjson9806e1DecodeGithubComHumansNetFcm3(in *jlexer.Lexer, out *AndroidNotification) {
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
		case "title":
			out.Title = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "icon":
			out.Icon = string(in.String())
		case "color":
			out.Color = string(in.String())
		case "sound":
			out.Sound = string(in.String())
		case "tag":
			out.Tag = string(in.String())
		case "click_action":
			out.ClickAction = string(in.String())
		case "body_loc_key":
			out.BodyLocKey = string(in.String())
		case "body_loc_args":
			if in.IsNull() {
				in.Skip()
				out.BodyLocArgs = nil
			} else {
				in.Delim('[')
				if out.BodyLocArgs == nil {
					if !in.IsDelim(']') {
						out.BodyLocArgs = make([]string, 0, 4)
					} else {
						out.BodyLocArgs = []string{}
					}
				} else {
					out.BodyLocArgs = (out.BodyLocArgs)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.BodyLocArgs = append(out.BodyLocArgs, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "title_loc_key":
			out.TitleLocKey = string(in.String())
		case "title_loc_args":
			if in.IsNull() {
				in.Skip()
				out.TitleLocArgs = nil
			} else {
				in.Delim('[')
				if out.TitleLocArgs == nil {
					if !in.IsDelim(']') {
						out.TitleLocArgs = make([]string, 0, 4)
					} else {
						out.TitleLocArgs = []string{}
					}
				} else {
					out.TitleLocArgs = (out.TitleLocArgs)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.TitleLocArgs = append(out.TitleLocArgs, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "channel_id":
			out.ChannelId = string(in.String())
		case "ticker":
			out.Ticker = string(in.String())
		case "sticky":
			out.Sticky = bool(in.Bool())
		case "event_name":
			out.EventName = string(in.String())
		case "local_only":
			out.LocalOnly = bool(in.Bool())
		case "notification_priority":
			out.NotificationPriority = NotificationPriority(in.String())
		case "default_sound":
			out.DefaultSound = bool(in.Bool())
		case "default_vibrate_timings":
			out.DefaultVibrateTimings = bool(in.Bool())
		case "default_light_settings":
			out.DefaultLightSettings = bool(in.Bool())
		case "vibrate_timings":
			if in.IsNull() {
				in.Skip()
				out.VibrateTimings = nil
			} else {
				in.Delim('[')
				if out.VibrateTimings == nil {
					if !in.IsDelim(']') {
						out.VibrateTimings = make([]string, 0, 4)
					} else {
						out.VibrateTimings = []string{}
					}
				} else {
					out.VibrateTimings = (out.VibrateTimings)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.VibrateTimings = append(out.VibrateTimings, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "visibility":
			out.Visibility = Visibility(in.String())
		case "notification_count":
			out.NotificationCount = int(in.Int())
		case "light_settings":
			if in.IsNull() {
				in.Skip()
				out.LightSettings = nil
			} else {
				if out.LightSettings == nil {
					out.LightSettings = new(LightSettings)
				}
				easyjson9806e1DecodeGithubComHumansNetFcm5(in, out.LightSettings)
			}
		case "image":
			out.Image = string(in.String())
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
func easyjson9806e1EncodeGithubComHumansNetFcm3(out *jwriter.Writer, in AndroidNotification) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Title != "" {
		const prefix string = ",\"title\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	if in.Body != "" {
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Body))
	}
	if in.Icon != "" {
		const prefix string = ",\"icon\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Icon))
	}
	if in.Color != "" {
		const prefix string = ",\"color\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Color))
	}
	if in.Sound != "" {
		const prefix string = ",\"sound\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sound))
	}
	if in.Tag != "" {
		const prefix string = ",\"tag\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Tag))
	}
	if in.ClickAction != "" {
		const prefix string = ",\"click_action\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClickAction))
	}
	if in.BodyLocKey != "" {
		const prefix string = ",\"body_loc_key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.BodyLocKey))
	}
	if len(in.BodyLocArgs) != 0 {
		const prefix string = ",\"body_loc_args\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.BodyLocArgs {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.TitleLocKey != "" {
		const prefix string = ",\"title_loc_key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TitleLocKey))
	}
	if len(in.TitleLocArgs) != 0 {
		const prefix string = ",\"title_loc_args\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.TitleLocArgs {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	if in.ChannelId != "" {
		const prefix string = ",\"channel_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ChannelId))
	}
	if in.Ticker != "" {
		const prefix string = ",\"ticker\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ticker))
	}
	if in.Sticky {
		const prefix string = ",\"sticky\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Sticky))
	}
	if in.EventName != "" {
		const prefix string = ",\"event_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EventName))
	}
	if in.LocalOnly {
		const prefix string = ",\"local_only\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.LocalOnly))
	}
	if in.NotificationPriority != "" {
		const prefix string = ",\"notification_priority\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.NotificationPriority))
	}
	if in.DefaultSound {
		const prefix string = ",\"default_sound\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.DefaultSound))
	}
	if in.DefaultVibrateTimings {
		const prefix string = ",\"default_vibrate_timings\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.DefaultVibrateTimings))
	}
	if in.DefaultLightSettings {
		const prefix string = ",\"default_light_settings\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.DefaultLightSettings))
	}
	if len(in.VibrateTimings) != 0 {
		const prefix string = ",\"vibrate_timings\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v12, v13 := range in.VibrateTimings {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	if in.Visibility != "" {
		const prefix string = ",\"visibility\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Visibility))
	}
	if in.NotificationCount != 0 {
		const prefix string = ",\"notification_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.NotificationCount))
	}
	if in.LightSettings != nil {
		const prefix string = ",\"light_settings\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9806e1EncodeGithubComHumansNetFcm5(out, *in.LightSettings)
	}
	if in.Image != "" {
		const prefix string = ",\"image\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Image))
	}
	out.RawByte('}')
}
func easyjson9806e1DecodeGithubComHumansNetFcm5(in *jlexer.Lexer, out *LightSettings) {
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
		case "color":
			easyjson9806e1DecodeGithubComHumansNetFcm6(in, &out.Color)
		case "light_on_duration":
			out.LightOnDuration = string(in.String())
		case "light_off_duration":
			out.LightOffDuration = string(in.String())
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
func easyjson9806e1EncodeGithubComHumansNetFcm5(out *jwriter.Writer, in LightSettings) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"color\":"
		out.RawString(prefix[1:])
		easyjson9806e1EncodeGithubComHumansNetFcm6(out, in.Color)
	}
	if in.LightOnDuration != "" {
		const prefix string = ",\"light_on_duration\":"
		out.RawString(prefix)
		out.String(string(in.LightOnDuration))
	}
	if in.LightOffDuration != "" {
		const prefix string = ",\"light_off_duration\":"
		out.RawString(prefix)
		out.String(string(in.LightOffDuration))
	}
	out.RawByte('}')
}
func easyjson9806e1DecodeGithubComHumansNetFcm6(in *jlexer.Lexer, out *Color) {
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
		case "red":
			out.Red = float32(in.Float32())
		case "green":
			out.Green = float32(in.Float32())
		case "blue":
			out.Blue = float32(in.Float32())
		case "alpha":
			out.Alpha = float32(in.Float32())
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
func easyjson9806e1EncodeGithubComHumansNetFcm6(out *jwriter.Writer, in Color) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Red != 0 {
		const prefix string = ",\"red\":"
		first = false
		out.RawString(prefix[1:])
		out.Float32(float32(in.Red))
	}
	if in.Green != 0 {
		const prefix string = ",\"green\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Green))
	}
	if in.Blue != 0 {
		const prefix string = ",\"blue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Blue))
	}
	if in.Alpha != 0 {
		const prefix string = ",\"alpha\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Alpha))
	}
	out.RawByte('}')
}
func easyjson9806e1DecodeGithubComHumansNetFcm1(in *jlexer.Lexer, out *Notification) {
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
		case "title":
			out.Title = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "image":
			out.Image = string(in.String())
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
func easyjson9806e1EncodeGithubComHumansNetFcm1(out *jwriter.Writer, in Notification) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Title != "" {
		const prefix string = ",\"title\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	if in.Body != "" {
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Body))
	}
	if in.Image != "" {
		const prefix string = ",\"image\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Image))
	}
	out.RawByte('}')
}
