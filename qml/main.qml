import QtQuick 2.8
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.3

ApplicationWindow {
    id: window
    width: 450
    height: 470
    color: "#ffffff"
    visible: true
    title: "NautiCalc by Adrian Reid"
    minimumWidth: 450
    minimumHeight: 470

    //  TAB BAR
    TabBar {
        id: tabBar
        position: TabBar.Header
        currentIndex: 0
        anchors.bottom: parent.top
        anchors.right: parent.right
        anchors.left: parent.left
        anchors.top: parent.top
        anchors.bottomMargin: -40

        TabButton {
            text: qsTr("About")
        }
        TabButton {
            text:qsTr("Compass Error")
        }
        TabButton {
            text:qsTr("Gyro Error")
        }
    }

    // ---- ACTUAL CONTENT AREA
    StackLayout {
        id: stackView
        currentIndex: tabBar.currentIndex
        anchors.topMargin: 10
        anchors.bottom: txt_copyright.top
        anchors.right: parent.right
        anchors.left: parent.left
        anchors.top: tabBar.bottom

        // ---- ABOUT
        Item {
            id: aboutTab
            anchors.fill: parent
            anchors.margins: 2
            Text {
                id: txt_about
                text: "This program calculates compass error and gyro error for you.\n\nTo get started, click on one of the tabs above.\n\nDisclaimer: This application has not been tested for every situation."
                anchors.fill: parent
                wrapMode: Text.WordWrap
                verticalAlignment: Text.AlignTop
                horizontalAlignment: Text.AlignLeft
            }
        }

        // ---- COMPASS ERROR
        Item {
            id: compassTab
            Layout.fillWidth: true
            anchors.fill: parent
            anchors.margins: 2
            Text {
                id: txt_compassAbout
                text: qsTr("Calculate compass error by entering the required values below.")
                wrapMode: Text.WordWrap
                anchors.right: parent.right
                anchors.left: parent.left
                anchors.top: parent.top
            }

            ToolSeparator {
                id: toolSeparator_compass
                orientation: Qt.Horizontal
                anchors.top: txt_compassAbout.bottom
                anchors.left: parent.left
                anchors.right: parent.right
            }

            GridLayout {
                id: grid_compass
                anchors.right: parent.right
                anchors.rightMargin: 10
                columns: 2
                anchors.top: toolSeparator_compass.bottom
                anchors.left: parent.left
                anchors.leftMargin: 10

                Text {
                    text: "True Heading/Bearing:"
                }
                TextField {
                    id: input_true
                }

                Text {
                    text: "Magnetic Heading/Bearing:"
                }
                TextField {
                    id: input_magnetic
                }

                Text {
                    text: "Variation:"
                }
                TextField {
                    id: input_variation
                }

                Text {
                    text: "Variation Direction:"
                }
                ComboBox {
                    id: combo_varDir
                    width: 200
                    currentIndex: 0
                    model: ["E", "W"]
                }
            }

            Button {
                id: btn_calcCompError
                text: qsTr("Calculate")
                anchors.bottom: parent.bottom
                anchors.bottomMargin: 5
                anchors.margins: 10
                anchors.left: parent.left
                anchors.right: parent.right
                onClicked: qmlBridge.calcCompError(input_true.text, input_magnetic.text, input_variation.text, combo_varDir.currentText)
            }
        }

        // ---- GYRO ERROR
        Item {
            id: gyroTab
            Layout.fillHeight: true
            Layout.fillWidth: true
            Text {
                id: txt_gyroAbout
                text: qsTr("Calculate gyro error by entering the required values below.\nWhen entering degrees and minutes use the below format:\n    285 degrees 39.3 minutes -> 285-39.3")
                wrapMode: Text.WordWrap
                anchors.right: parent.right
                anchors.left: parent.left
                anchors.top: parent.top
                anchors.margins: 2
            }

            ToolSeparator {
                id: toolSeparator_gyro
                orientation: Qt.Horizontal
                anchors.top: txt_gyroAbout.bottom
                anchors.left: parent.left
                anchors.right: parent.right
            }

            GridLayout {
                id: grid_gyro
                anchors.right: parent.right
                anchors.rightMargin: 10
                columns: 2
                anchors.top: toolSeparator_gyro.bottom
                anchors.left: parent.left
                anchors.leftMargin: 10

                Text {
                    id: txtGyro_gyro
                    text: "Gyro Bearing:"
                }
                TextField {
                    id: inputGyro_gyro
                }

                Text {
                    id: txtGyro_lat
                    text: "Latitude:"
                }
                TextField {
                    id: inputGyro_lat
                }

                Text {
                    id: txtGyro_latDir
                    text: "Latitude Direction"
                }
                ComboBox {
                    id: inputGyro_latDir
                    width: 200
                    model: ["N", "S"]
                    currentIndex: 0
                }

                Text {
                    id: txtGyro_LHA
                    text: "LHA"
                }
                TextField {
                    id: inputGyro_LHA
                }

                Text {
                    id: txtGyro_decl
                    text: "Declination:"
                }
                TextField {
                    id: inputGyro_decl
                }

                Text {
                    id: txtGyro_declDir
                    text: "Declination Direction"
                }
                ComboBox {
                    id: inputGyro_declDir
                    width: 200
                    model: ["N", "S"]
                    currentIndex: 0
                }
            }

            Button {
                id: btn_calcGyroError
                text: qsTr("Calculate")
                anchors.bottom: parent.bottom
                anchors.bottomMargin: 5
                anchors.margins: 10
                anchors.left: parent.left
                anchors.right: parent.right
                onClicked: qmlBridge.calcGyroError(inputGyro_gyro.text, inputGyro_lat.text, inputGyro_latDir.currentText, inputGyro_LHA.text, inputGyro_decl.text, inputGyro_declDir.currentText)
            }
        }

    }

    //  FOOTER
    Text {
        id: txt_copyright
        text: qsTr("Copyright 2018 - Adrian Reid")
        anchors.horizontalCenter: parent.horizontalCenter
        verticalAlignment: Text.AlignBottom
        anchors.bottom: parent.bottom
        font.pixelSize: 10
    }

	//	CUSTOM MESSAGE DIALOG
	Item {
        id: msgDialogWindow
        width: window.width / 1.5
        height: msgDialogText.height + 92
        anchors.horizontalCenter: stackView.horizontalCenter
        anchors.top: stackView.top
        visible: false

        Rectangle {
            id: rectangle
            radius: 10
            border.width: 2
            border.color: "#4990ff"
            anchors.fill: parent

            Text {
                id: msgDialogTitle
                anchors.top: parent.top
                width: parent.width
                text: "Title"
                font.bold: true
                horizontalAlignment: Text.AlignHCenter
                anchors.topMargin: 10
            }
            ToolSeparator
            {
                id: toolSeparator_MsgBox
                width: parent.width
                bottomPadding: 0
                topPadding: 0
                hoverEnabled: false
                anchors.top: msgDialogTitle.bottom
                padding: 5
                orientation: Qt.Horizontal

            }
            Text {
                id: msgDialogText
                anchors.top: toolSeparator_MsgBox.bottom
                anchors.left: parent.left
                anchors.right: parent.right
                text: "The quick brown fox jumps over the lazy dog."
                fontSizeMode: Text.Fit
                wrapMode: Text.WrapAtWordBoundaryOrAnywhere
                anchors.margins: 10
            }
            Button {
                id: btn_Okay
                anchors.top: msgDialogText.bottom
                anchors.horizontalCenter: parent.horizontalCenter
                anchors.margins: 5
                text: "Close"
                anchors.topMargin: 7
                onClicked: {
                    msgDialogWindow.visible = false
                    stackView.enabled = true
                    tabBar.enabled = true
                }
            }
        }
    }

    Connections {
		target: qmlBridge
		onShowMessageBox:
		{
            msgDialogTitle.text = msgTitle
            msgDialogText.text = msgText
            stackView.enabled = false
            tabBar.enabled = false
            msgDialogWindow.visible = true
		}
	}
}
