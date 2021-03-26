package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/mcuadros/go-octoprint"
	"github.com/Z-Bolt/OctoScreen/utils"
)

type TemperaturePresetButton struct {
	*gtk.Button

	client						*octoprint.Client
	selectToolStepButton		*SelectToolStepButton
	imageFileName				string
	temperaturePreset			*octoprint.TemperaturePreset
	callback					func()
}

func CreateTemperaturePresetButton(
	client						*octoprint.Client,
	selectToolStepButton		*SelectToolStepButton,
	imageFileName				string,
	temperaturePreset			*octoprint.TemperaturePreset,
	callback					func(),
) *TemperaturePresetButton {
	presetName := utils.StrEllipsisLen(temperaturePreset.Name, 10)
	base := utils.MustButtonImage(presetName, imageFileName, nil)

	instance := &TemperaturePresetButton{
		Button:						base,
		client:						client,
		selectToolStepButton:		selectToolStepButton,
		imageFileName:				imageFileName,
		temperaturePreset:			temperaturePreset,
		callback:					callback,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *TemperaturePresetButton) handleClicked() {
	utils.Logger.Infof("TemperaturePresetButton.handleClicked() - setting temperature to preset %s.", this.temperaturePreset.Name)
	utils.Logger.Infof("TemperaturePresetButton.handleClicked() - setting hotend temperature to %.0f.", this.temperaturePreset.Extruder)
	utils.Logger.Infof("TemperaturePresetButton.handleClicked() - setting bed temperature to %.0f.", this.temperaturePreset.Bed)

	currentTool := this.selectToolStepButton.Value()
	if currentTool == "" {
		utils.Logger.Error("TemperaturePresetButton.handleClicked() - currentTool is invalid (blank), defaulting to tool0")
		currentTool = "tool0"
	}

	/*
	CreateTemperaturePresetButton is used by TemperaturePresetsPanel.  Strictly speaking,
	CreateTemperaturePresetButton should only set the temperature of one device at at time,
	but that's a lousy UX.  Imagine being in the TemperaturePanel... with the tool set to
	the hotend, click the More button (and go to the TemperaturePresetsPanel), then
	clicking PLA (and get taken back to the TemperaturePanel), __THEN__ have to click the
	tool button to change to the bed, and then repeat the process over again.

	So, instead, the temperature of both the bed and the selected tool (or tool0 if the bed
	is selected) are set.
	*/

	// Set the bed's temp.
	bedTargetRequest := &octoprint.BedTargetRequest{Target: this.temperaturePreset.Bed}
	err := bedTargetRequest.Do(this.client)
	if err != nil {
		utils.LogError("TemperaturePresetButton.handleClicked()", "Do(BedTargetRequest)", err)
		return
	}

	// Set the hotend's temp.
	var toolTargetRequest *octoprint.ToolTargetRequest
	if currentTool == "bed" {
		// If current tool is set to "bed", use tool0.
		toolTargetRequest = &octoprint.ToolTargetRequest{Targets: map[string]float64{"tool0": this.temperaturePreset.Extruder}}
	} else {
		toolTargetRequest = &octoprint.ToolTargetRequest{Targets: map[string]float64{currentTool: this.temperaturePreset.Extruder}}
	}

	err = toolTargetRequest.Do(this.client)
	if err != nil {
		utils.LogError("TemperaturePresetButton.handleClicked()", "Do(ToolTargetRequest)", err)
	}

	if this.callback != nil {
		this.callback()
	}
}
