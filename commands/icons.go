package commands

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const (
	inkscapeCmd  = "inkscape"
	convertCmd   = "convert"
	pngIcnsCmd   = "png2icns"
	pngExtension = ".png"
)

type IconsParams struct {
	InputFile       string
	OutputDirectory string
	Size16x16       bool
	Size32x32       bool
	Size64x64       bool
	Size128x128     bool
	Size256x256     bool
	Size512x512     bool
	GenerateIco     bool
	GenerateIcns    bool
	Verbose         bool
}

type resolution struct {
	width  int
	height int
	label  string
}

func (ip *IconsParams) getInputFileBaseName() string {
	ext := filepath.Ext(ip.InputFile)
	fileName := filepath.Base(ip.InputFile)
	if ext == "" {
		return fileName
	}
	return strings.Split(fileName, ext)[0]
}

func (ip *IconsParams) supportedResolutions() []resolution {
	resolutions := make([]resolution, 0)
	if ip.Size16x16 {
		resolutions = append(resolutions, resolution{
			width:  16,
			height: 16,
			label:  "16x16",
		})
	}
	if ip.Size32x32 {
		resolutions = append(resolutions, resolution{
			width:  32,
			height: 32,
			label:  "32x32",
		})
	}
	if ip.Size64x64 {
		resolutions = append(resolutions, resolution{
			width:  64,
			height: 64,
			label:  "64x64",
		})
	}
	if ip.Size128x128 {
		resolutions = append(resolutions, resolution{
			width:  128,
			height: 128,
			label:  "128x128",
		})
	}
	if ip.Size256x256 {
		resolutions = append(resolutions, resolution{
			width:  256,
			height: 256,
			label:  "256x256",
		})
	}
	if ip.Size512x512 {
		resolutions = append(resolutions, resolution{
			width:  512,
			height: 512,
			label:  "512x512",
		})
	}
	return resolutions
}

func (ip *IconsParams) pngOutputFile(outDir, label string) string {
	pngFileName := ip.getInputFileBaseName() + "_" + label + pngExtension
	return path.Join(outDir, pngFileName)
}

func GenerateIcons(params *IconsParams) error {
	inputFilePath, err := filepath.Abs(params.InputFile)
	if err != nil {
		return err
	}
	if params.Verbose {
		fmt.Printf("Input file: %s\n", inputFilePath)
	}
	outDir, err := filepath.Abs(params.OutputDirectory)
	if err != nil {
		return err
	}
	if params.Verbose {
		fmt.Printf("Output directory: %s\n", outDir)
	}
	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		log.Printf("Failed to create output directory")
		return err
	}
	err = generatePngIcons(params, inputFilePath, outDir)
	if err != nil {
		return err
	}
	err = generateIco(params, outDir)
	if err != nil {
		return err
	}
	err = generateIcns(params, outDir)
	return err
}

func generatePngIcons(params *IconsParams, inputFilePath, outDir string) error {
	for _, res := range params.supportedResolutions() {
		width := fmt.Sprintf("%d", res.width)
		height := fmt.Sprintf("%d", res.height)
		outFile := params.pngOutputFile(outDir, res.label)
		inkscape := exec.Command(inkscapeCmd, "-w", width, "-h", height, "-o", outFile, inputFilePath)
		out, err := inkscape.CombinedOutput()
		if err != nil {
			log.Printf("Failed to run inkscape: %s\n", err)
			return err
		}
		if params.Verbose {
			fmt.Println(string(out))
		}
	}
	return nil
}

func generateIco(params *IconsParams, outputDir string) error {
	if !params.GenerateIco {
		log.Println("Skip generating .ico file")
		return nil
	}
	pngFiles := getAllPngFilesInDir(outputDir, false)
	icoFile := path.Join(outputDir, params.getInputFileBaseName()+".ico")
	convertArgs := append(pngFiles, icoFile)
	convert := exec.Command(convertCmd, convertArgs...)
	out, err := convert.CombinedOutput()
	if err != nil {
		return err
	}
	if params.Verbose {
		fmt.Println(string(out))
	}
	return nil
}

func generateIcns(params *IconsParams, outputDir string) error {
	if !params.GenerateIcns {
		log.Println("Skip generating .icns file")
		return nil
	}
	pngFiles := getAllPngFilesInDir(outputDir, true)
	icnsFile := path.Join(outputDir, params.getInputFileBaseName()+".icns")
	args := []string{icnsFile}
	args = append(args, pngFiles...)
	png2icns := exec.Command(pngIcnsCmd, args...)
	out, err := png2icns.CombinedOutput()
	if err != nil {
		log.Printf("Failed to run png2icns: %s\n", err)
		return err
	}
	if params.Verbose {
		fmt.Println(string(out))
	}
	return nil
}

func getAllPngFilesInDir(directory string, isIcns bool) []string {
	var pngFiles []string
	filepath.WalkDir(directory, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		isPng := filepath.Ext(s) == pngExtension
		doAppend := (isPng && !isIcns) || (isPng && isIcns && !strings.HasSuffix(s, "64x64.png"))
		if doAppend {
			pngFiles = append(pngFiles, s)
		}
		return nil
	})
	return pngFiles
}
