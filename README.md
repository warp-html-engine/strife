# strife
A simple game framework that wraps around SDL2.

## notes
I pretty much copy/pasted this from one of my game projects into a separate repository, so it's all bundled into one module. Maybe I'll clean it up later, but hey ho, it works for now.

This is pretty feature-less right now. I'll be expanding it as my needs progress in a little project I'm working on.

## example
Here's a little example:

	package main

	import (
		"github.com/felixangell/strife"
	)

	func main() {
		window, err := strife.CreateRenderWindow(1280, 720, strife.DefaultConfig())
		if err != nil {
			panic(err)
		}

		for !window.CloseRequested() {
			ctx := window.GetRenderContext()
			ctx.Clear()

			{
				ctx.SetColor(strife.RGB(255, 0, 255))
				ctx.Rect(10, 10, 50, 50, strife.Fill)			
			}

			ctx.Display()
		}
	}

## license
[MIT](/LICENSE)