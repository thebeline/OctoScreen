package uiWidgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/mcuadros/go-octoprint"
	"github.com/Z-Bolt/OctoScreen/utils"
)

type TemperatureIncreaseButton struct {
	*gtk.Button

	client							*octoprint.Client
	temperatureAmountStepButton		*TemperatureAmountStepButton
	selectToolStepButton			*SelectToolStepButton
	isIncrease						bool
}

func CreateTemperatureIncreaseButton(
	client							*octoprint.Client,
	temperatureAmountStepButton		*TemperatureAmountStepButton,
	selectToolStepButton			*SelectToolStepButton,
	isIncrease						bool,
) *TemperatureIncreaseButton {
	var base *gtk.Button
	if isIncrease {
		base = utils.MustButtonImageStyle("Increase", "increase.svg", "", nil)
	} else {
		base = utils.MustButtonImageStyle("Decrease", "decrease.svg", "", nil)
	}

	instance := &TemperatureIncreaseButton{
		Button:							base,
		client:							client,
		temperatureAmountStepButton:	temperatureAmountStepButton,
		selectToolStepButton:			selectToolStepButton,
		isIncrease:						isIncrease,
	}
	instance.Button.Connect("clicked", instance.handleClicked)

	return instance
}

func (this *TemperatureIncreaseButton) handleClicked() {
	value := this.temperatureAmountStepButton.Value()
	tool := this.selectToolStepButton.Value()
	target, err := utils.GetToolTarget(this.client, tool)
	if err != nil {
		utils.LogError("TemperatureIncreaseButton.handleClicked()", "GetToolTarget()", err)
		return
	}

	if this.isIncrease {
		target += value
	} else {
		target -= value
	}

	if target < 0 {
		target = 0
	}

	// TODO: should the target be checked for a max temp?
	// If so, how to calculate what the max should be?

	utils.Logger.Infof("TemperatureIncreaseButton.handleClicked() - setting target temperature for %s to %1.f°C.", tool, target)

	err = utils.SetToolTarget(this.client, tool, target)
	if err != nil {
		utils.LogError("TemperatureIncreaseButton.handleClicked()", "GetToolTarget()", err)
	}
}
