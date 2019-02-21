![The Cocktail logo][tcklogo] ❤ ![Karumi logo][karumilogo]
# Kata Maxibon for Go. [![Build Status](https://travis-ci.org/nando/MaxibonKataGo.svg?branch=master)](https://travis-ci.org/nando/MaxibonKataGo)

**DISCLAIMER: Work in progress.**

- We are here to practice [property based testing][property-based-testing].
- We are going to use [rapid][rapid] to write our tests.
- We are going to practice pair programming.

During this kata we will try to find how to use property based testing from different points of view:

* Using property based testing and [TDD][tdd] at the same time.
* Using property based testing to cover legacy code.
* Using the already tested code to check if the tests are working fine (manual mutation testing).

If you want to learn just property based testing using an already implemented application go to the branch ``write-tests``. If you prefer to practice TDD in the ``master`` branch you will find the kata already resolved. If you prefer to practice mutation testing stay in the ``master`` branch.

## Getting started

Karumi developers love ice cream. And one of our favorites ice cream is named Maxibon:

![Maxibon][maxibon]

Summer is comming and our small team sometime needs Maxibons to work better. But in the Karumi HQs finding a Maxibon is not always easy. We start every week with 10 Maxibons but once there are just 2 Maxibons or less we need to buy more.

Karumi developers can consume zero or a positive number of maxibons. The Karumi team is composed by five engineers and everytime some of these engineers go to the kitchen they grab some maxibons as follows:

* If the developer is Pedro, he grabs three maxibons.
* If the developer is Davide, he does not grab any maxibon.
* If the developer is Alberto, he grabs one maxibon.
* If the developer is Jorge, he grabs two maxibons.
* If the developer is Sergio, he grabs one maxibon.

When a Karumi engineer goes to the kitchen, they can go in group if needed, and there are just 2 maxibons or less left he has to send a message through the Slack API saying ``"Hi guys, I'm <NAME OF THE DEVELOPER>. We need more maxibons!"``. And the number of maxibons available will be automatically incremented by 10 :). If the number of maxibons left is lower than the number of maxibons the developer tries to get he will get just the number of maxibons available.

## Tasks

Your task as Go Developer is to resolve this problem or test an already implemented software, depending on the path you have choosen before to start.

* If you want to practice TDD write a little piece of software to implement the already described scenario following the TDD cycle.
* If you don not want to follow the TDD path create an empty project, add the [rapid][rapid] gem as a dependency and start writing code. 
* If you prefer to just learn how to use [property based testing][property-based-testing] you can use this repository just checkout to the branch ``write-tests``.

## Extra tasks:

* Change the initial number of maxibons to be configurable.
* Add the verbose modifier to the properties to review if the generation is working fine.
* Change your production code to check if your tests fail or not.
* Configure the number of generations used per property.

**This repository is ready to build the application, pass the checkstyle and your tests in Travis-CI environments.**

# License

Copyright (c) 2019 Fernando García Samblas. The Cocktail Experience, S.L.

This Source Code Form is subject to the terms of the Mozilla Public License,
v. 2.0. See LICENSE file for details. If a copy of the MPL was not distributed
with this file you can obtain one at https://mozilla.org/MPL/2.0/.

[karumilogo]: https://cloud.githubusercontent.com/assets/858090/11626547/e5a1dc66-9ce3-11e5-908d-537e07e82090.png
[tcklogo]: https://avatars0.githubusercontent.com/u/1177560?s=40
[property-based-testing]: http://es.slideshare.net/ScottWlaschin/an-introduction-to-property-based-testing
[rapid]: https://github.com/flyingmutant/rapid
[maxibon]: ./art/maxibon.jpg
[tdd]: https://en.wikipedia.org/wiki/Test-driven_development

