# Texture Generation
[![Go Reference](https://pkg.go.dev/badge/github.com/jphsd/texture.svg)](https://pkg.go.dev/github.com/jphsd/texture)
[![Go Report Card](https://goreportcard.com/badge/github.com/jphsd/texture)](https://goreportcard.com/report/github.com/jphsd/texture)
[![Build Status](https://travis-ci.com/jphsd/texture.svg?branch=master)](https://travis-ci.com/jphsd/texture)

A package for the procedural generation of textures. Based on the ideas contained in the Bryce 3D deep texture editor.

![example](/doc/01.png?raw=true "Example")

More examples [here](/Examples1.md "Examples1"), [here](/Examples2.md "Examples2") and [here](/Examples3.md "Examples3").

The primary interfaces allow for the evaluation of a value, vector or color field at any point in the XY plane.

A subpackage covers the generation of surfaces based on lights illuminating a material.
