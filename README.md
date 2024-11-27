# Go-ViGempad

<p>
    <a href="https://pkg.go.dev/github.com/CypherpunkSamurai/go-vigempad?tab=doc"><img src="https://godoc.org/github.com/CypherpunkSamurai/go-vigempad?status.svg" alt="GoDoc"></a>
</p>

A simple Go library to interact with the [ViGEmBus driver](https://vigembusdriver.com) and the [ViGEmClient](https://github.com/nefarius/ViGEmClient) library.

## Installation

Install the package with the following command:

```bash
go get github.com/CypherpunkSamurai/go-vigempad
```

Download and Install [ViGem Gamepad Driver](https://github.com/nefarius/ViGEmBus/releases)

Download ViGemClient dll from [Nefarius's Archive](https://buildbot.nefarius.at/builds/ViGEmClient/master/1.21.222.0/bin/release/) for your architecture (x86 or x64) and place it in your project directory.

Rename the dll to `ViGEmClient.dll` and place it in the same directory as your project, or your executable binary.

Compile your go project with the following command:

```bash
# Build the project
go build .
```

## Usage

A few examples are provided in the [examples](examples) directory. Use 

- [btn_test](examples/btn_test/main.go) - Presses the A button on the X360 controller.
- [trigger_test](examples/trigger_test/main.go) - Demonstrates how to use the X360 throttle function to control the vibration motors of the first connected controller.
- [axis_test](examples/axis_test/main.go) - Demonstrates how to use the X360 axis function to control the thumbsticks of the first connected controller.

## TODO

- [ ] Add support for more controller types.

# Credits

- [Nefarius](https://docs.nefarius.at/projects/ViGEm/) for ViGem Bus and ViGem Client.
