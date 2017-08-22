package mpd

const (
	ProfileFull        = "urn:mpeg:dash:profile:full:2011"
	ProfileISOOnDemand = "urn:mpeg:dash:profile:isoff-on-demand:2011"
	ProfileISOMain     = "urn:mpeg:dash:profile:isoff-main:2011"
	ProfileISOLive     = "urn:mpeg:dash:profile:isoff-live:2011"
	ProfileTSMain      = "urn:mpeg:dash:profile:mp2t-main:2011"
	ProfileTSSimple    = "urn:mpeg:dash:profile:mp2t-simple:2011"

	staticMPD  = "static"
	dynamicMPD = "dynamic"
	MPDXMLNS   = "urn:mpeg:dash:schema:mpd:2011"

	SchemeIdUri = "urn:mpeg:dash:23003:3:audio_channel_configuration:2011"
)

type MPD struct {
	Id                         string                  `xml:"id,attr,omitempty"`
	Profiles                   string                  `xml:"profiles,attr"`
	Type                       string                  `xml:"type,attr,omitempty"`
	AvailabilityStartTime      string                  `xml:"availabilityStartTime,attr,omitempty"`
	PublishTime                string                  `xml:"publishTime,attr,omitempty"`
	MediaPresentationDuration  string                  `xml:"mediaPresentationDuration,attr,omitempty"`
	AvailabilityEndTime        string                  `xml:"availabilityEndTime,attr,omitempty"`
	mediaPresentationDuration  string                  `xml:"mediaPresentationDuration,attr,omitempty"`
	MinimumUpdatePeriod        string                  `xml:"minimumUpdatePeriod,attr,omitempty"`
	MinBufferTime              string                  `xml:"minBufferTime,attr"`
	SuggestedPresentationDelay string                  `xml:"suggestedPresentationDelay,attr,omitempty"`
	MaxSegmentDuration         string                  `xml:"maxSegmentDuration,attr,omitempty"`
	MaxSubsegmentDuration      string                  `xml:"maxSubsegmentDuration,attr,omitempty"`
	Xmlns                      string                  `xml:"xmlns,attr"`
	ProgramInformation         []ProgramInformationXML `xml:"ProgramInformation,omitempty"`
	BaseURL                    []BaseURLXML            `xml:"BaseURL,omitempty"`
	Location                   []string                `xml:"Location,omitempty"`
	Period                     []PeriodXML             `xml:"Period"`
	Metrics                    []MetricsXML            `xml:"Metrics,omitempty"`
}

type ProgramInformationXML struct {
	Lang               string `xml:"lang,attr,omitempty"`
	MoreInformationURL string `xml:"moreInformationURL,attr,omitempty"`
	Title              string `xml:"Title,omitempty"`
	Source             string `xml:"Source,omitempty"`
	Copyright          string `xml:"Copyright"`
}

type BaseURLXML struct {
	ServiceLocation          string `xml:"serviceLocation,attr,omitempty"`
	ByteRange                string `xml:"byteRange,attr,omitempty"`
	AvailabilityTimeOffset   string `xml:"availabilityTimeOffset,attr,omitempty"`
	AvailabilityTimeComplete string `xml:"availabilityTimeComplete,attr,omitempty"`
}

type PeriodXML struct {
	Id                 string             `xml:"id,attr,omitempty"`
	Start              string             `xml:"start,attr,omitempty"`
	Duration           string             `xml:"duration,attr,omitempty"`
	BitstreamSwitching *bool              `xml:"bitstreamSwitching,attr,omitempty"`
	BaseURL            []BaseURLXML       `xml:"BaseURL,omitempty"`
	AdaptationSet      []AdaptationSetXML `xml:"AdaptationSet,omitempty"`
}

type MetricsXML struct {
	Metrics   string         `xml:"metrics,attr"`
	Range     []MetricsRange `xml:"Range,omitempty"`
	Reporting []string       `xml:"Reporting"`
}

type MetricsRange struct {
	Starttime string `xml:"starttime,attr,omitempty"`
	Duration  string `xml:"duration,attr,omitempty"`
}

type AdaptationSetXML struct {
	Xlinkhref                 string                        `xml:"xlink:href,attr,omitempty"`
	Xlinkactuate              string                        `xml:"xlink:actuate,attr,omitempty"`
	Id                        string                        `xml:"id,attr,omitempty"`
	Group                     string                        `xml:"group,attr,omitempty"`
	Lang                      string                        `xml:"lang,attr,omitempty"`
	ContentType               string                        `xml:"contentType,attr,omitempty"`
	Par                       string                        `xml:"par,attr,omitempty"`
	MinBandwidth              string                        `xml:"minBandwidth,attr,omitempty"`
	MaxBandwidth              string                        `xml:"maxBandwidth,attr,omitempty"`
	Width                     string                        `xml:"width,attr,omitempty"`
	Height                    string                        `xml:"height,attr,omitempty"`
	FrameRate                 string                        `xml:"frameRate,attr,omitempty"`
	SegmentAlignment          bool                          `xml:"segmentAlignment,attr,omitempty"`
	SubsegmentAlignment       bool                          `xml:"subsegmentAlignment,attr,omitempty"`
	SubsegmentStartsWithSAP   int                           `xml:"subsegmentStartsWithSAP,attr,omitempty"`
	MimeType                  string                        `xml:"mimeType,attr"`
	Codecs                    string                        `xml:"codecs,attr,omitempty"`
	AudioChannelConfiguration *AudioChannelConfigurationXML `xml:"AudioChannelConfiguration,omitempty"`
	SegmentTemplate           SegmentTemplateXML            `xml:"SegmentTemplate"`
	Representation            []RepresentationXML           `xml:"Representation,omitempty"`
}

type SegmentTemplateXML struct {
	Media           string              `xml:"media,attr"`
	Initialization  string              `xml:"initialization,attr"`
	Duration        *int                `xml:"duration,attr,omitempty"`
	StartNumber     string              `xml:"startNumber,attr"`
	TimeScale       string              `xml:"timescale,attr"`
	SegmentTimeline *SegmentTimelineXML `xml:"SegmentTimeline,omitempty"`
}

type SegmentTimelineXML struct {
	S []SegmentTimelineDesc `xml:"S"`
}

type SegmentTimelineDesc struct {
	T string `xml:"t,attr,omitempty"` //time
	D string `xml:"d,attr"`           //duration
	R string `xml:"r,attr,omitempty"` //repreat count default 0
}

type RepresentationXML struct {
	Id                string `xml:"id,attr"`
	Bandwidth         string `xml:"bandwidth,attr"`
	Width             string `xml:"width,attr,omitempty"`
	Height            string `xml:"height,attr,omitempty"`
	FrameRate         string `xml:"frameRate,attr,omitempty"`
	AudioSamplingRate string `xml:"audioSamplingRate,attr,omitempty"`
}

type AudioChannelConfigurationXML struct {
	SchemeIdUri string `xml:"schemeIdUri,attr"`
	Value       int    `xml:"value,attr"`
}