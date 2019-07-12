// main.qml - qml definition for the gui application
// Copyright (C) 2018-2019  Jedrzej Stuczynski.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

import QtQuick 2.12
import QtQuick.Controls 2.5
import QtQuick.Layouts 1.12
// import QtQuick.Dialogs 1.3
import Qt.labs.platform 1.1 as QtLabs

Flickable {
    id: walletPage
    anchors.fill: parent

    ScrollIndicator.vertical: ScrollIndicator { }

    ColumnLayout {
        spacing: 5
        anchors.fill: parent
        anchors.margins: 30
        Label {
            id: label3
            text: qsTr("Nym Wallet Demo")
            horizontalAlignment: Text.AlignHCenter
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignHCenter | Qt.AlignTop
            font.weight: Font.DemiBold
            font.pointSize: 16
        }

        RowLayout {
            Layout.alignment: Qt.AlignTop
            Layout.preferredHeight: 40
            Layout.fillWidth: true
            TextField {
                id: path
                enabled: false
                text: "Please load Nym Client configuration file"
                Layout.fillWidth: true
            }
            Button {
                text: "open config"
                onClicked: fileDialog.open();
            }
        }

        ColumnLayout {
            property var config: ConfigBridge
            Layout.topMargin: 10
            Layout.alignment: Qt.AlignLeft | Qt.AlignTop
            spacing: 10

            id: configView
            Layout.fillWidth: true
            Layout.fillHeight: true

            // TODO: configView.visible = !configView.visible on getting config data from Go

            Label {
                id: label2
                text: qsTr("Configuration Details")
                Layout.rowSpan: 2
                Layout.columnSpan: 4
                font.pointSize: 14
                font.weight: Font.DemiBold
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            }

            GridLayout {
                id: gridLayout
                width: 100
                height: 100
                rowSpacing: 10
                columnSpacing: 5
                rows: 2
                columns: 5
                Layout.fillHeight: true
                Layout.fillWidth: true

                Label {
                    text: "Client id: "
                    font.weight: Font.DemiBold
                }

                TextField {
                    id: clientId
                    enabled: false
                    text: configView.config.identifier
                    Layout.fillWidth: false
                }

                ToolSeparator {
                    id: toolSeparator
                    opacity: 0
                }

                Label {
                    text: "Client address: "
                    font.weight: Font.DemiBold
                }

                TextField {
                    id: clientAddress
                    enabled: false
                    text: configView.config.address
                    Layout.fillWidth: true
                }

                Label {
                    text: "Keyfile: "
                    font.weight: Font.DemiBold
                }

                TextField {
                    id: clientKeyfile
                    enabled: false
                    text: configView.config.keyfile
                    Layout.columnSpan: 4
                    Layout.fillWidth: true
                }

            }






            GroupBox {
                id: groupBox
                width: 200
                height: 200
                enabled: true
                Layout.fillHeight: false
                Layout.fillWidth: true
                title: qsTr("Ethereum Configuration")

                GridLayout {
                    id: gridLayout1
                    x: -12
                    width: 100
                    anchors.horizontalCenter: parent.horizontalCenter
                    anchors.left: parent.left
                    anchors.bottom: parent.bottom
                    anchors.top: parent.top
                    anchors.bottomMargin: 10
                    anchors.topMargin: 10
                    rows: 3
                    columns: 2
                    Layout.fillHeight: true
                    Layout.fillWidth: true

                    Label {
                        text: "Client id: "
                        font.weight: Font.DemiBold
                    }

                    TextField {
                        id: clientId1
                        text: configView.config.identifier
                        enabled: false
                        Layout.fillWidth: true
                    }

                    Label {
                        text: "Client id: "
                        font.weight: Font.DemiBold
                    }

                    TextField {
                        id: clientId2
                        text: configView.config.identifier
                        enabled: false
                        Layout.fillWidth: true
                    }

                    Label {
                        text: "Client id: "
                        font.weight: Font.DemiBold
                    }

                    TextField {
                        id: clientId3
                        text: configView.config.identifier
                        enabled: false
                        Layout.fillWidth: true
                    }
                }
            }

            RowLayout {
                id: rowLayout
                Layout.rowSpan: 1
                Layout.fillWidth: true

                GroupBox {
                    id: groupBox1
                    width: 200
                    height: 200
                    Layout.preferredHeight: 200
                    Layout.minimumHeight: 200
                    Layout.maximumHeight: 300
                    Layout.fillHeight: false
                    Layout.fillWidth: false
                    Layout.preferredWidth: parent.width/2
                    title: qsTr("Nym Validator Nodes")

                    ListView {
                        id: listView
                        x: 0
                        width: 110
                        height: 160
                        anchors.top: parent.top
                        anchors.topMargin: 10
                        anchors.horizontalCenter: parent.horizontalCenter

                        model: ListModel {
                            ListElement {
                                name: "Grey"
                                colorCode: "grey"
                            }

                            ListElement {
                                name: "Red"
                                colorCode: "red"
                            }

                            ListElement {
                                name: "Blue"
                                colorCode: "blue"
                            }

                            ListElement {
                                name: "Green"
                                colorCode: "green"
                            }
                        }
                        delegate: Item {
                            x: 5
                            width: 80
                            height: 40
                            Row {
                                id: row1
                                spacing: 10
                                Rectangle {
                                    width: 40
                                    height: 40
                                    color: colorCode
                                }

                                Text {
                                    text: name
                                    font.bold: true
                                    anchors.verticalCenter: parent.verticalCenter
                                }
                            }
                        }
                    }
                }

                GroupBox {
                    id: groupBox2
                    width: 200
                    height: 200
                    Layout.preferredHeight: 200
                    Layout.minimumHeight: 200
                    Layout.maximumHeight: 300
                    Layout.fillHeight: false
                    Layout.fillWidth: true
                    title: qsTr("Tendermint Validator Nodes")

                    ListView {
                        id: listView1
                        x: 0
                        width: 110
                        height: 160
                        anchors.top: parent.top
                        anchors.topMargin: 10
                        anchors.horizontalCenterOffset: 0
                        anchors.horizontalCenter: parent.horizontalCenter
                        model: ListModel {
                            ListElement {
                                name: "Grey"
                                colorCode: "grey"
                            }

                            ListElement {
                                name: "Red"
                                colorCode: "red"
                            }

                            ListElement {
                                name: "Blue"
                                colorCode: "blue"
                            }

                            ListElement {
                                name: "Green"
                                colorCode: "green"
                            }
                        }
                        delegate: Item {
                            x: 5
                            width: 80
                            height: 40
                            Row {
                                id: row2
                                spacing: 10
                                Rectangle {
                                    width: 40
                                    height: 40
                                    color: colorCode
                                }

                                Text {
                                    text: name
                                    font.bold: true
                                    anchors.verticalCenter: parent.verticalCenter
                                }
                            }
                        }
                    }
                }

            }

            RowLayout {
                Layout.fillHeight: false
                Layout.alignment: Qt.AlignBottom | Qt.AlignRight

                Layout.preferredHeight: 40
                Layout.fillWidth: true

                Button {
                    Layout.alignment: Qt.AlignRight | Qt.AlignBottom
                    id: configConfirmBtn
                    text: "confirm"
                    Layout.fillHeight: false
                    Layout.fillWidth: false
                    onClicked: console.log("Confirmed config")
                }


            }

















        }

    }

    QtLabs.FileDialog {
        id: fileDialog
        folder: QtLabs.StandardPaths.standardLocations(QtLabs.StandardPaths.HomeLocation)[0]
        nameFilters: [ "Config files (*.toml)", "All files (*)" ]
        // onFolderChanged: {
        //     folderModel.folder = folder;
        // }
        onAccepted: {
            console.log("You chose: " + fileDialog.file)
            QmlBridge.loadConfig(fileDialog.file)
            path.text = fileDialog.file
        }
    }
    //   ProgressBar {
    //     id: loading
    //     anchors.horizontalCenter: parent.horizontalCenter
    //     visible: false
    //     indeterminate: true
    //   }
    // }



    Dialog {
        id: notificationDialog
        parent: ApplicationWindow.contentItem
        anchors.centerIn: ApplicationWindow.contentItem

        height: 200
        width: Math.min(ApplicationWindow.contentItem.width * 2/3, 800)

        modal: true

        closePolicy: Popup.CloseOnPressOutside | Popup.CloseOnEscape
        standardButtons: Dialog.Ok | Dialog.Cancel
        title: qsTr("Notification box")

        Label {
            id: notificationText
            // maximumWidth: notificationDialog.contentWidth
        }

        onAccepted: console.log("Ok clicked")
        onRejected: console.log("Cancel clicked")
    }

    Connections {
        target: QmlBridge
        onDisplayNotification: {
            notificationText.text = message
            // notificationDialog.implicitHeight = notificationText.height
            // notificationDialog.implicitWidth = notificationText.width
            notificationDialog.open()
        }
    }

}

















/*##^## Designer {
    D{i:0;autoSize:true;height:768;width:1024}
}
 ##^##*/
