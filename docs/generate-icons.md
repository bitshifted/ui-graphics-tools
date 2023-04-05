# Generate icons from SVG file

Icons for Windows, Mac OS and Linux are generated from single SVG file. The tool will generate following icons:

* `.png` files for Linux with the following sizes: 16x16, 32x32, 64x64, 128x128,256x256,512x512
* `.ico` file for Windows, containing all sizes from above
* `.icns` file for Mac OS, containing all sizes from above, except 64x64

## Simple example

Let's say you have an SVG file called `my-icon.svg` in directory `my-dir`. Directory layout would look like this:

```
my-dir
  |- my-icon.svg
```

To generate icons from this SVG file, go into directory `my-dir` and run the following command:

```
docker run  -v ${PWD}:/workspace  ghcr.io/bitshifted/ui-graphics-tools:<version>  icons my-icon.svg
```

where `<version>` is the version of the tool you want to run. This will generate icons in directory `output/icons`. Directory layout will look like this:

```
my-dir
  |- output
     |- icons
       |- my-icon_16x16.png
       |- my-icon_32x32.png
       |- my-icon_64x64.png
       |- .....
       |- my-icon.ico
       |- my-icon.icns
  |- my-icon.svg
```

## Options

You can customize generation usiing the following options:

```
-v, --[no-]verbose  Verbose mode
--output-dir="output/icons"  
                Directory containing generated icons
--[no-]16x16    Generate 16x16 pixels PNG icon
--[no-]32x32    Generate 32x32 pixels PNG icon
--[no-]64x64    Generate 64x64 pixels PNG icon
--[no-]128x128  Generate 128x128 pixels PNG icon
--[no-]256x256  Generate 256x256 pixels PNG icon
--[no-]512x512  Generate 512x512 pixels PNG icon
--[no-]ico      Generate .ico type for Windows
--[no-]icns     Generate .icns type for Mac OS
```

* `--output-dir` - the directory where to generate icons
* `--no-[size]` - do not generate files of specified size
* `--no-ico` - do not generate .ico file
* `--no-icns` - do not generate .icns file

### Verbose mode

If you need to get more verbosse output, add `-v` options before `icons` command:

```
docker run  -v ${PWD}:/workspace  ghcr.io/bitshifted/ui-graphics-tools:<version> -v  icons my-icon.svg
```