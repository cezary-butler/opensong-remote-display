# OpenSong Remote Display

## What is it?

This is an application that allows to display whatever is displayed by OpenSong on a screen of a computer that is connected to the same network (possibly wireless). It allows to create a setup where you have one computer which runs OpenSong and is used to mange projection, and another (possibly something like Raspberry Pi) connected to the projector using wired connection, running this application effectively having limited wireless projector capabilities.
Originally it was developed because of the problems we had in our church with projectors connected to the computer using very long cable. Because of repeated setup and teardown the cables often failed.

Currently, the project is in a POC state. It means that it was made to prove that it is possible to achieve goal of the project but without much deal of design and attention to testing. Problems are to be expected. I am going to improve that over time especially if you'll let me know you'd like to use it. 

## Requirements

The project can be run on any computer that is able to run golang applications and has some kind raster display capabilities. 
Application uses OpenSong API, this means it needs at least version 2 of OpenSong with API enabled. Also, the firewall on both machines need to be configured to allow communication between this application and OpenSong instance. 

## Usage

To run the application you need to run command

    opensong-remote-display --host _api-address:api-port_

Optionally you can add _width_, _height_ (both integer from 1 to 4096), and _quality_ (integer between 0 and 100, the 100 is the best quality but also more data sent over the network) parameters to specify your display resolution (by default 1280x720), and quality of image produced by OpenSong API.

Example command with all parameters may look like:

    opensong-remote-display --host 192.168.82.55:8082 --width 1920 --height 1080 --quality 100

## How to build

To build the project you will need go compiler suite installed in version 1.16 or newer.

You can build the project by running:

    go build -o opensong-remote-display ./cmd/main.go

Alternatively you can use GNU Make to build the project using:

    make clean build

## State of the project

This is a POC. Please don't expect seamless experience yet. I will update status as the project progresses. Since I have very limited time for the project it may take a few weeks to reach reliability. Bugs and feature requests are welcome.

## Known problems

1. There seems to be a memory leak within integration with fyne, which may cause problems especially for sessions with a lot of slide changes.
2. Connection process is not yet reliable, this means that it might happen that application will fail to receive events from OpenSong API. However, when it will connect, it will keep working until it is terminated or network connection is lost.
3. There is no reconnection yet so when connection is lost, you will need to restart the application when connection is restored. 

## Plans for the nearest future

1. Providing binaries for major platforms, possibly using goreleaser
2. Update dependencies (most importantly fyne) to the latest version   
3. Refactoring towards reliable design
4. Testing to remove remaining bugs proving reliability 
