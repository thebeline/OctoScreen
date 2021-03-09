#!/usr/bin/env bash

REPO="Z-Bolt/OctoScreen"
MAIN="master"
RELEASES="https://api.github.com/repos/$REPO/releases/latest"
WGET_RAW="https://github.com/$REPO/raw/$MAIN"

source <(wget -qO- "$WGET_RAW/scripts/inquirer.sh")

exit

arch=$(uname -m)
#if [[ $arch == x86_64 ]]; then
#    releaseURL=$(curl -s "$RELEASES" | grep "browser_download_url.*amd64.deb" | cut -d '"' -f 4)
#elif [[ $arch == aarch64 ]]; then
#    releaseURL=$(curl -s "$RELEASES" | grep "browser_download_url.*arm64.deb" | cut -d '"' -f 4)
if  [[ $arch == arm* ]]; then
    releaseURL=$(curl -s "$RELEASES" | grep "browser_download_url.*armf.deb" | cut -d '"' -f 4)
fi
dependencies="libgtk-3-0 xserver-xorg xinit x11-xserver-utils"
IFS='/' read -ra version <<< "$releaseURL"

echo "Installing OctoScreen "${version[7]}, $arch""

echo "Installing Dependencies ..."
sudo apt -qq update
sudo apt -qq install $dependencies -y

if [ -d "/home/pi/OctoPrint/venv" ]; then
    DIRECTORY="/home/pi/OctoPrint/venv"
elif [ -d "/home/pi/oprint" ]; then
    DIRECTORY="/home/pi/oprint"
else
    echo "Neither /home/pi/OctoPrint/venv nor /home/pi/oprint can be found."
    echo "If your OctoPrint instance is running on a different machine just type - in the following prompt."
    text_input "Please specify OctoPrints full virtualenv path manually (no trailing slash)." DIRECTORY
fi;

if [ $DIRECTORY == "-" ]; then
    echo "Not installing any plugins for remote installation. Please make sure to have Display Layer Progress installed."
elif [ ! -d $DIRECTORY ]; then
    echo "Can't find OctoPrint Installation, please run the script again!"
    exit 1
fi;

#if [ $DIRECTORY != "-" ]; then
#  plugins=( 'Display Layer Progress (mandatory)' 'Filament Manager' 'Preheat Button' 'Enclosure' 'Print Time Genius' 'Ultimaker Format Package' 'PrusaSlicer Thumbnails' )
#  checkbox_input "Which plugins should I install (you can also install them via the Octoprint UI)?" plugins selected_plugins
#  echo "Installing Plugins..."
#
#  if [[ " ${selected_plugins[@]} " =~ "Display Layer Progress (mandatory)" ]]; then
#      "$DIRECTORY"/bin/pip install -q --disable-pip-version-check "https://github.com/OllisGit/OctoPrint-DisplayLayerProgress/releases/latest/download/master.zip"
#  fi;
#  if [[ " ${selected_plugins[@]} " =~ "Filament Manager" ]]; then
#      "$DIRECTORY"/bin/pip install -q --disable-pip-version-check "https://github.com/OllisGit/OctoPrint-FilamentManager/releases/latest/download/master.zip"
#  fi;
#  if [[ " ${selected_plugins[@]} " =~ "Preheat Button" ]]; then
#      "$DIRECTORY"/bin/pip install -q --disable-pip-version-check "https://github.com/marian42/octoprint-preheat/archive/master.zip"
#  fi;
#  if [[ " ${selected_plugins[@]} " =~ "Enclosure" ]]; then
#      "$DIRECTORY"/bin/pip install -q --disable-pip-version-check "https://github.com/vitormhenrique/OctoPrint-Enclosure/archive/master.zip"
#  fi;
#  if [[ " ${selected_plugins[@]} " =~ "Print Time Genius" ]]; then
#      "$DIRECTORY"/bin/pip install -q --disable-pip-version-check "https://github.com/eyal0/OctoPrint-PrintTimeGenius/archive/master.zip"
#  fi;
#  if [[ " ${selected_plugins[@]} " =~ "Ultimaker Format Package" ]]; then
#      "$DIRECTORY"/bin/pip install -q --disable-pip-version-check "https://github.com/jneilliii/OctoPrint-UltimakerFormatPackage/archive/master.zip"
#  fi;
#  if [[ " ${selected_plugins[@]} " =~ "PrusaSlicer Thumbnails" ]]; then
#      "$DIRECTORY"/bin/pip install -q --disable-pip-version-check "https://github.com/jneilliii/OctoPrint-PrusaSlicerThumbnails/archive/master.zip"
#  fi;
#fi;

echo "Installing OctoScreen "${version[7]}, $arch" ..."
cd ~
wget -O octoscreen.deb $releaseURL -q --show-progress

sudo dpkg -i octodash.deb

rm octodash.deb

yes_no=( 'yes' 'no' )

list_input "Shall I reboot your Pi now?" yes_no reboot
echo "OctoScreen has been successfully installed! :)"
if [ $reboot == 'yes' ]; then
    sudo reboot
fi