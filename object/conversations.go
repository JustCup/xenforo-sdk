package object // import "github.com/JustCup/xenforo-sdk/object"

// Conversation struct.
type Conversation struct {
	Username            string      `json:"username"`
	Recipients          interface{} `json:"recipients"`
	IsStarted           bool        `json:"is_starred"`
	CanEdit             bool        `json:"can_edit"`
	CanReply            bool        `json:"can_reply"`
	CanInvite           bool        `json:"can_invite"`
	CanUploadAttachment bool        `json:"can_upload_attachment"`
	ConversationID      int         `json:"conversation_id"`
	Title               string      `json:"title"`
	UserID              int         `json:"user_id"`
	StartDate           int         `json:"start_date"`
	OpenInvite          bool        `json:"open_invite"`
	ConversationOpen    bool        `json:"conversation_open"`
	ReplyCount          int         `json:"reply_count"`
	RecipientCount      int         `json:"recipient_count"`
	FirstMessageID      int         `json:"first_message_id"`
	LastMessageDate     int         `json:"last_message_date"`
	LastMessageID       int         `json:"last_message_id"`
	LastMessageUserID   int         `json:"last_message_user_id"`
	Starter             User        `json:"Starter"`
}

// ConversationMessage struct.
type ConversationMessage struct {
	Username          string        `json:"username"`
	MessageParsed     string        `json:"message_parsed"` // HTML parsed version of the message contents.
	CanEdit           bool          `json:"can_edit"`
	CanReact          bool          `json:"can_react"`
	ViewURL           string        `json:"view_url"`
	Conversation      Conversation  `json:"Conversation"`        // Conditionally returned. If requested by context, the conversation this message is part of.
	Attachments       []interface{} `json:"Attachments"`         // Conditionally returned. If there are attachments to this message, a list of attachments.
	IsReactedTo       bool          `json:"is_reacted_to"`       // True if the viewing user has reacted to this content.
	VisitorReactionID uint          `json:"visitor_reaction_id"` // If the viewer reacted, the ID of the reaction they used.
	MessageID         uint          `json:"message_id"`
	ConversationID    uint          `json:"conversation_id"`
	MessageDate       uint32        `json:"message_date"`
	UserID            uint          `json:"user_id"`
	Message           string        `json:"message"`
	AttachCount       uint          `json:"attach_count"`
	ReactionScore     uint          `json:"reaction_score"`
	User              User          `json:"User"`
}

// ConversationsResponse struct.
type ConversationsResponse struct {
	Success       bool                  `json:"success,omitempty"`
	Conversation  Conversation          `json:"conversation"`
	Conversations []Conversation        `json:"conversations"`
	Message       ConversationMessage   `json:"message"`
	Action        string                `json:"action"`
	Messages      []ConversationMessage `json:"messages"`
	Pagination    Pagination            `json:"pagination"`
}
