Lori
=============

A practical functional library for Golang programmers .

![angry gopher](https://sdl-stickershop.line.naver.jp/products/0/0/1/1299646/android/stickers/12120436.png)

Lori's goal
-----------

Lori is a library that aims to offer pure functional functions that help make programming clear and easier ,
Golang is a great language that can be used in several tasks from microservices to systems programming ...
altough being an imperative language Golang offers several qualities of functional languages like functions 
being first class citizens, and great type inference , and it would benefit from a functional flavor .


Package Organization
---------------------

Lori is modular it follows the idea that you should only import what you need , it's organized into 
multiple subpackages that provide each a set of functions that operate on a datastructure or type .

* List : all functions in the list subpackage operate on golang slices .
* Math : any extra mathy things you need can be found here .
* Function : provides higher order function utils .
* Relation : provides functions that offer relation and order primitives .

Docs
--------------

Documentation will be up soon .


Philosophy
-----------

* Data structures are pure Golang structures with no mutations .
* A Function should never mutate an object or produce side effects .
* It should always be the programmer's error not the library .


What to expect
------------

First forgive me : 

* This is the creation part don't expect perfection or stability .

By the end of the few following months :

* A Stable API
* Concise and Clear documentation
* Great things 







