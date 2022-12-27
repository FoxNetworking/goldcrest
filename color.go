package main

import (
	tempest "github.com/Amatsagu/Tempest"
	"strconv"
	"strings"
)

var colorCmd = tempest.Command{
	Name:        "color",
	Description: "Sets the color of your role!",
	Options: []tempest.Option{
		{
			Name:        "color",
			Description: "Role color, in hex form",
			Type:        tempest.OPTION_STRING,
			Required:    true,
			// fff
			MinLength: 3,
			// #ffffff
			MaxLength: 7,
			Focused:   true,
		},
	},
	SlashCommandHandler: func(itx tempest.CommandInteraction) {
		// Validate that this is actually been 3 to 6, and hexadecimal.
		// Discord should validate the min/max lengths for us, but you never know...
		rawColor, exists := itx.GetOptionValue("color")
		if !exists {
			itx.SendReply(tempest.ResponseData{
				Content: "No color specified!",
			}, false)
			return
		}

		// Given:
		//   - #fff - valid
		//   - ffffff - valid
		//   - ff - invalid, not 3 or 6
		//   - fffff - invalid, not 3 or 6
		var colorString = rawColor.(string)
		// Remove the prefixed #, if necessary.
		colorString = strings.TrimPrefix(colorString, "#")

		if len(colorString) != 3 && len(colorString) != 6 {
			itx.SendReply(tempest.ResponseData{
				Content: "Invalid color specified!",
			}, true)
			return
		}

		// Decode from hexadecimal. This can and will fail,
		// as we don't validate [0-9a-fA-F].
		color, err := strconv.ParseUint(colorString, 16, 32)
		if err != nil {
			itx.SendReply(tempest.ResponseData{
				Content: "Invalid hexadecimal color specified!",
			}, true)
			return
		}
		if len(colorString) == 3 {
			// We should have a color similar to #fff if this is true.
			// 0xfff * 0x1001 = 0xffffff
			color *= 0x1001
		}

		itx.SendReply(tempest.ResponseData{
			Embeds: []*tempest.Embed{
				{
					Title: "Testing...",
					Color: uint32(color),
				},
			},
		}, false)
	},
}
