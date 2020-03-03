package gourd

import (
	"context"
	"github.com/andersfylling/disgord"
)

type CommandContext struct {
	prefix      string
	args        []string
	commandUsed string
	message     *disgord.Message
	client      *disgord.Client
	gourd       *Gourd
	command     *Command
}

// Prefix returns the prefix used in the command
func (ctx *CommandContext) Prefix() string {
	return ctx.prefix
}

// Args returns the args included with the command, if any
func (ctx *CommandContext) Args() []string {
	return ctx.args
}

func (ctx *CommandContext) CommandUsed() string {
	return ctx.commandUsed
}

// Message returns the message object for the command message.
// It is not recommended to use this directly, use the convenience methods instead, if possible
func (ctx *CommandContext) Message() *disgord.Message {
	return ctx.message
}

// Client returns the disgord client
// It is not recommended to use this directly, use the convenience methods instead, if possible
func (ctx *CommandContext) Client() *disgord.Client {
	return ctx.client
}

func (ctx *CommandContext) Command() *Command {
	return ctx.command
}

// Reply sends a message to the channel the command was used in.
// Input is any type, see https://github.com/andersfylling/disgord/blob/39ba986ca2e94602ce44f4bf7625063124bdc325/client.go#L705
func (ctx *CommandContext) Reply(data ...interface{}) (*disgord.Message, error) {
	msg, err := ctx.message.Reply(context.Background(), ctx.client, data...)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ctx *CommandContext) IsPrivate() bool {
	return ctx.message.IsDirectMessage()
}

func (ctx *CommandContext) Guild() (*disgord.Guild, error) {
	guild, err := ctx.client.GetGuild(context.Background(), ctx.message.GuildID)
	if err != nil {
		return nil, err
	}

	return guild, nil
}

func (ctx *CommandContext) Author() *disgord.User {
	return ctx.message.Author
}

func (ctx *CommandContext) AuthorMember() *disgord.Member {
	return ctx.message.Member
}

func (ctx *CommandContext) IsAuthorOwner() bool {
	return ctx.Author().ID.String() == ctx.gourd.ownerId
}
