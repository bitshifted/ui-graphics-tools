package cli

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/bitshifted/ui-graphics-tools/commands"
)

const (
	defaultIconsOutputDir        = "output/icons"
	defaultSplashScreenOutputDir = "output/splash"
	defaultSplashScreenFileName  = "splash.png"
)

var (
	app     = kingpin.New("ui-graphics-tools", "Tools for UI graphics generation")
	verbose = app.Flag("verbose", "Verbose mode").Short('v').Bool()

	iconsCmd    = app.Command("icons", "Generates OS specific icons from SVG")
	inputFile   = iconsCmd.Arg("input-file", "Input file in SVG format").Required().String()
	outputDir   = iconsCmd.Flag("output-dir", "Directory containing generated icons").Default(defaultIconsOutputDir).String()
	size16x16   = iconsCmd.Flag("16x16", "Generate 16x16 pixels PNG icon").Default("true").Bool()
	size32x32   = iconsCmd.Flag("32x32", "Generate 32x32 pixels PNG icon").Default("true").Bool()
	size64x64   = iconsCmd.Flag("64x64", "Generate 64x64 pixels PNG icon").Default("true").Bool()
	size128x128 = iconsCmd.Flag("128x128", "Generate 128x128 pixels PNG icon").Default("true").Bool()
	size256x256 = iconsCmd.Flag("256x256", "Generate 256x256 pixels PNG icon").Default("true").Bool()
	size512x512 = iconsCmd.Flag("512x512", "Generate 512x512 pixels PNG icon").Default("true").Bool()
	icoType     = iconsCmd.Flag("ico", "Generate .ico type for Windows").Default("true").Bool()
	icnsType    = iconsCmd.Flag("icns", "Generate .icns type for Mac OS").Default("true").Bool()

	splashCmd         = app.Command("splash", "Generates splash screen in PNG format")
	splashConfigFile  = splashCmd.Flag("config-file", "Splash screen configuration file").Required().String()
	splashOutputDir   = splashCmd.Flag("output-dir", "Splash screen output directory").Default(defaultSplashScreenOutputDir).String()
	splashOutFileName = splashCmd.Flag("output-file", "Splash screen output file").Default(defaultSplashScreenFileName).String()
)

func Run() error {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case iconsCmd.FullCommand():
		params := commands.IconsParams{
			InputFile:       *inputFile,
			OutputDirectory: *outputDir,
			Size16x16:       *size16x16,
			Size32x32:       *size32x32,
			Size64x64:       *size64x64,
			Size128x128:     *size128x128,
			Size256x256:     *size256x256,
			Size512x512:     *size512x512,
			GenerateIco:     *icoType,
			GenerateIcns:    *icnsType,
			Verbose:         *verbose,
		}
		return commands.GenerateIcons(&params)
	case splashCmd.FullCommand():
		params := commands.SplashParams{
			ConfigFile:  *splashConfigFile,
			OutputDir:   *splashOutputDir,
			OutFileName: *splashOutFileName,
			Verbose:     *verbose,
		}
		return commands.GenerateSplashScreen(&params)
	default:
		fmt.Println("Unknown command")
		return nil
	}
}
