// NautiCalc (main.go)
// Copyright 11/05/2017
// Author: Adrian Reid
// Updated 02/05/2018

// A tool for nautical calculations

package main

import (
	"edtheloon/nauticalclib"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

func float(s string) float64 {
	var f float64
	_, _ = fmt.Sscanf(s, "%f", &f)
	return f
}

func splitDeg(deg string) float64 {
	split := strings.Split(deg, "-")
	degs := float(split[0])
	mins := float(split[1]) / 60
	newDegrees := degs + mins
	return newDegrees
}

func gyroError(gyro string, lat string, latDir string, lha string, decl string, declDir string) (string, error) {
	// 	NautiCalc library will handle the calculations
	var ge nauticalclib.GyroError

	// Regular Expression Error Checking then convert values
	rLat := regexp.MustCompile("[0-9]{1,2}-[0-9]{1,2}.[0-9]$")
	rLong := regexp.MustCompile("[0-9]{1,3}-[0-9]{1,2}.[0-9]$")

	if !rLong.MatchString(gyro) {
		return "", errors.New("Incorrect value (Gyro Bearing of Object): " + gyro)
	}
	if !rLat.MatchString(lat) {
		return "", errors.New("Incorrect value (Latitude): " + lat)
	}
	if !rLong.MatchString(lha) {
		return "", errors.New("Incorrect value (LHA): " + lha)
	}
	if !rLat.MatchString(decl) {
		return "", errors.New("Incorrect value (Declination): " + decl)
	}
	ge.Gyro = float(gyro)
	ge.Latitude = splitDeg(lat)
	ge.LatDir = latDir
	ge.LHA = splitDeg(lha)
	ge.Declination = splitDeg(decl)
	ge.DeclDir = declDir

	ge.Calculate()

	// Create slices to store the results to show
	//var res1, res2 []string
	//res1 = append(res1, "Latitude:")
	//res2 = append(res2, fmt.Sprintf("%3.01f %v", ge.Latitude, ge.LatDir))
	//res1 = append(res1, "LHA:")
	//res2 = append(res2, fmt.Sprintf("%3.01f", ge.LHA))
	//res1 = append(res1, "Declination:")
	//res2 = append(res2, fmt.Sprintf("%3.01f %v", ge.Declination, ge.DeclDir))
	//res1 = append(res1, "A3:")
	//res2 = append(res2, fmt.Sprintf("%v %3.01f %v", ge.CDir, ge.A3, ge.AzimuthDir))
	//res1 = append(res1, "Azimuth:")
	//res2 = append(res2, fmt.Sprintf("%3.01f", ge.Azimuth))
	//res1 = append(res1, "Gyro:")
	//res2 = append(res2, fmt.Sprintf("%3.01f", ge.Gyro))
	//res1 = append(res1, "Error:")
	//res2 = append(res2, fmt.Sprintf("%3.01f %v", ge.GyroErr, ge.ErrDir))

	var result string
	result = fmt.Sprintf("Latitude:\t%3.01f %v\n", ge.Latitude, ge.LatDir)
	result += fmt.Sprintf("LHA:\t%3.01f\n", ge.LHA)
	result += fmt.Sprintf("Declination:\t%3.01f %v\n", ge.Declination, ge.DeclDir)
	result += fmt.Sprintf("Azimuth:\t%v %3.01f %v\n", ge.CDir, ge.A3, ge.AzimuthDir)
	result += fmt.Sprintf("Gyro:\t%3.01f\n", ge.Gyro)
	result += fmt.Sprintf("Error:\t%3.01f %v", ge.GyroErr, ge.ErrDir)
	return result, nil
}

func compError(mag string, gyro string, variation string, varDir string) (string, error) {
	// NautiCalc library will handle the calculations
	var ce nauticalclib.CompassError

	// RegExp error checking
	reg := regexp.MustCompile("^[0-9]{1,3}$|^[0-9]{1,3}.[0-9]{1,3}$")
	if !reg.MatchString(mag) {
		return "", errors.New("Incorrect value (Magnetic): " + mag)
	}
	if !reg.MatchString(gyro) {
		return "", errors.New("Incorrect value (True): " + gyro)
	}
	if !reg.MatchString(variation) {
		return "", errors.New("Incorrect value (Variation): " + variation)
	}

	// Convert any strings to float64
	ce.Magnetic = float(mag)
	ce.Gyro = float(gyro)
	ce.Variation = float(variation)
	ce.VarDir = varDir

	ce.Calculate()

	// Create slices to store the results to show
	//var res1, res2 []string
	//res1 = append(res1, "Magnetic:")
	//res2 = append(res2, fmt.Sprintf("%3.01f", ce.Magnetic))
	//res1 = append(res1, "Deviation:")
	//res2 = append(res2, fmt.Sprintf("%3.01f %v", ce.Deviation, ce.DevDir))
	//res1 = append(res1, "Corrected:")
	//res2 = append(res2, fmt.Sprintf("%3.01f", ce.Corrected))
	//res1 = append(res1, "Variation:")
	//res2 = append(res2, fmt.Sprintf("%3.01f %v", ce.Variation, ce.VarDir))
	//res1 = append(res1, "Gyro/True:")
	//res2 = append(res2, fmt.Sprintf("%3.01f", ce.Gyro))
	//res1 = append(res1, "Error:")
	//res2 = append(res2, fmt.Sprintf("%3.01f %v", ce.ComErr, ce.ErrDir))
	
	var result string
	result = fmt.Sprintf("Compass:\t%3.01f\n", ce.Magnetic)
	result += fmt.Sprintf("Deviation:\t%3.01f %v\n", ce.Deviation, ce.DevDir)
	result += fmt.Sprintf("Magnetic:\t%3.01f\n", ce.Corrected)
	result += fmt.Sprintf("Variation:\t%3.01f %v\n", ce.Variation, ce.VarDir)
	result += fmt.Sprintf("True:\t%3.01f\n", ce.Gyro)
	result += fmt.Sprintf("Error:\t%3.01f %v", ce.ComErr, ce.ErrDir)
	return result, nil
}

//go:generate qtmoc
type QmlBridge struct {
	core.QObject

	// Signal for calcCompError
	_ func(trueBrg string, magBrg string,
		variation string, varDir string) `slot:"calcCompError"`
	_ func(trueBrg string, lat string, latDir string, LHA string,
		decl string, declDir string) `slot:"calcGyroError"`
	_ func(msgTitle string, msgText string) `signal:"showMessageBox"`
}

func messageBox(title string, text string) {
	qmlBridge.ShowMessageBox(title, text)
}

var qmlBridge = NewQmlBridge(nil)

func main() {
	// Create application
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	// Enable high DPI scaling
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Use the material style for qml
	quickcontrols2.QQuickStyle_SetStyle("material")

	// Create a QML application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	// BRIDGING
	qmlBridge.ConnectCalcCompError(
		func(trueBrg string, magBrg string, variation string,
			varDir string) {
				result, err := compError(magBrg, trueBrg, variation, varDir)
				if err != nil {
					messageBox("Error", err.Error())
				} else {
					messageBox("Compass Error Results", result)
				}
	})
	qmlBridge.ConnectCalcGyroError(
		func(trueBrg string, lat string, latDir string, LHA string,
			decl string, declDir string) {
				result, err := gyroError(trueBrg, lat, latDir, LHA, decl, declDir)
				if err != nil {
					messageBox("Error", err.Error())
				} else {
					messageBox("Gyro Error Results", result)
				}
	})
	qmlBridge.ConnectShowMessageBox(
		func(msgTitle string, msgText string) {
	})
	engine.RootContext().SetContextProperty("qmlBridge", qmlBridge)

	// Load the main qml file
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	// Execute app
	gui.QGuiApplication_Exec()
}
