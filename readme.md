# Toy Robot

The Toy Robot is a fun program that enchances the skills of a coder

## Overview

Toy Robot is an excercise in implementing a small program that moves a robot along a 5x5 block grid and ensures that it can move accross the board in a valid manner. It will warn the player if they are about to fall off of the grid and will also prevent them from doing so.

## Getting Started

To run the Toy Robot, clone the repo and simply run go run main.go (Do ensure that you have go installed on your os)

## Approach

The approach is to use methods for each command of the Robot. One of the main reasons for this is for keeping the state of the Robot.
Another reason would be for future implementations of possible interfaces, in which we could simply pass Robot as a parameter - This would ensure that the interface implements each part that it needs to.
There are some downsides to this, mainly that validating input commands can become tricky. However, since we care about the state of the robot, we deem methods the better option.

Once the methods are in place we can write the logic for calling them.

## Struggles

I did encounter some issues with implementing the tests around the methods, since the expected behaviour can be difficult to track. However,
once I implemented proper error handling this became a bit easier.

The other main issue I was facing is that my program would exit after handling the error, since the main funcion now came to its end. This was finally resolved by adding a loop to Instruct call in order to continue the program

## Improvements

There are still some major improvements to make here. The main improvement is going to center around creating a proper repository structure. This way you can separate each part to its own detail. We can also improve the overall functionality by improving things like switch cases, adding better error handling (perhaps a wrapper function) and improving readability.

I am also in need of adding a file reader for passing a file with commands.

## Future 

Toy Robot has much potential. We could make the course bigger. We can implement a better front-end to improve user friendliness - a grid that shows the robot's position, buttons for each command. We could add another player to the game to implement some sort of battle. We could add obstacles to the course. Etc.

## Experience

As it stands, this program took me about 2.5 hours to build as I ran into some unexpected errors. I learned about the main function closing once it has reached the end. I used github co-pilot to help me out with some improvements on my code as well as suggestions on best approaches to the program as a whole. 

