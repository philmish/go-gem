# go-gem - A genral project environment manager

When working on a project in Linux you often encounter commands which are specific to
a project but are very long or so cluttered with flags that remembering them is getting more complicated with every project.
But saving them to your .bash_aliases might seem unnecessary, because you will never use them outside of this project.
Another aspect of projects which is unique for each project are the upcoming todos.

go-gem centralizes the storage of environment specific commands and todos for a project in a single json file.
It also allows you to create a .gem_aliases file in the root directory of the project.
This aliases file can be sourced to use the saved commands by their assigned aliases while working on the project.

## Install

go-gem requires you to be on a linux machine with go installed.

Clone the repo
```
git clone https://github.com/philmish/go-gem
```

cd into the go-gem directory and build go-gem in your .local/bin by using the Makefile
```
cd go-gem
make build
```

You can now call go-gem with the gogem command from you command line.

## Default Environments

The init command offers a option to specify a default environment to create.
These templates exist for:


```
node
vue
go
python

```
These keys can be used to create as template names for the init command.

## Usage

```
gogem -c help [command]
gogem -c init [-n <template name>] [--alias]
gogem -c ls
gogem -c add -n <cmd alias> -a <shell cmd> [<args>...]
gogem -c do -n <cmd alias>
gogem -c rm -n <cmd alias>
gogem -c lstodo
gogem -c lsdone
gogem -c addtodo -n <todo content> -a <todo urgency>
gogem -c deltodo -n <todo id>
gogem -c churg -n <todo id> -a <new urgency>
```
