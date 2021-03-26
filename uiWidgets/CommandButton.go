package uiWidgets

import (
	// "fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/Z-Bolt/OctoScreen/logger"
	"github.com/Z-Bolt/OctoScreen/octoprintApis"
	"github.com/Z-Bolt/OctoScreen/octoprintApis/dataModels"
	"github.com/Z-Bolt/OctoScreen/utils"
)

type CommandButton struct {
	*gtk.Button

	client				*octoprintApis.Client
	parentWindow		*gtk.Window
	commandDefinition	*dataModels.CommandDefinition
}

func CreateCommandButton(
	client				*octoprintApis.Client,
	parentWindow		*gtk.Window,
	commandDefinition	*dataModels.CommandDefinition,
	iconName			string,
) *CommandButton {
	base := utils.MustButtonImage(utils.StrEllipsisLen(commandDefinition.Name, 16), iconName + ".svg", nil)
	instance := &CommandButton {
		Button:				base,
		client:				client,
		parentWindow:		parentWindow,
		commandDefinition:	commandDefinition,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *CommandButton) handleClicked() {
	if len(this.commandDefinition.Confirm) != 0 {
		utils.MustConfirmDialogBox(this.parentWindow, this.commandDefinition.Confirm, this.sendCommand)
		return
	} else {
		this.sendCommand()
	}
}

func (this *CommandButton) sendCommand() {
	commandRequest := &octoprintApis.SystemExecuteCommandRequest{
		Source: dataModels.Custom,
		Action: this.commandDefinition.Action,
	}

	err := commandRequest.Do(this.client)
	if err != nil {
		logger.LogError("CommandButton.sendCommand()", "Do(SystemExecuteCommandRequest)", err)
	}
}
