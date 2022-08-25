## Amazing terminal application

A copy from snake-life repo by some youtuber

go.sum is some generated file i think and go.mod is a dependency / package info file
i only changed the name of the package in the latter one.

# Rendering is stupid atm

It's stupid that the rectangle-generators are responsible of
rendering themselves to the screen. We should instead store the
renderable items to a que of sort, and then parce trough the rendering-que with each game-loop.

We can approach this in different ways, but I think we should
start by theorizing the optimal format of the final renderable object or map of sort.
