package claude

import (
	"fmt"

	"github.com/tidwall/gjson"
)

// GetConversations is used to get conversations
func (c *Client) GetConversations() (*gjson.Result, error) {
	uri := fmt.Sprintf("/api/organizations/%s/chat_conversations", c.GetOrgid())

	return c.Get(uri)
}

// NewConversation is used to new conversation
func (c *Client) NewConversation(params MixMap) (*gjson.Result, error) {
	uri := fmt.Sprintf("/api/organizations/%s/chat_conversations", c.opts.Orgid)

	return c.Post(uri, params)
}
