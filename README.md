# Lumos ⚡️ ✨ 💫
### <em>"Light up your city, your country, your world."</em>

Lumos Maxima, an earth data provider. Lumos is (initially) a Proof of Concept which implements a Telegram Bot interface to make the knowledge and data of every place in the World accesible for everyone. Lumos aims to raise awaresness about the world's largest problems.

## Motivation

World is facing so many challenges and every place you visit has a story behind. Either for vacation or business everybody needs to know more about the context of a country/city. Their own challenges, projects ongoing, technology resources and so on:

- How climate change is being addressed?
- How polluted is the outdoor air?
- Is this country creating energy or just wasting? How?
- How's the economy?
- Is the water clean and fresh? 

...and thousands of questions we might have when we arrive to some place in this planet.

Lumos comes to give us concrete data at anytime anywhere.

**DISCLAIMER:** this started as a Proof of Concept and it aims to use all public data sources out there. The extra motivation in addition is to learn about technology, world data indicators and most importantly about Global Warming and how countries are progressing towards [Sustainable Development Goals](https://sdg-tracker.org/). 

## Sources

[OurWorldInData](https://ourworldindata.org/) is the main data source to feed Lumos. Initially data and topics are pre-selected and manually downloaded. Future iterations will bring up more robust and faster updated in order to integrate as much indicator as possible.
The main and final goal is to build a tool where every citizen can query any custom indicator (near-realtime) of every single country/city in the world at any time and also be aware about projects or news of the following topics:

- Demographic Change
- Health
- Food and Agriculture
- Energy and Environment
- Innovation and Technological Change
- Poverty and Economic Development
- Living conditions, Community and Wellbeing
- Human rights and Democracy
- Violence and War
- Education and Knowledge
- Sustainable Development Goals Tracker

(All these topics are published in [OurWorldInData](https://ourworldindata.org/))

You can see Lumos as a middleware and data aggregator. Initially it uses one data source but the idea is to include more.

## Lumos Bot (@lumosfy_bot)

The first interface of Lumos is a Telegram Bot. Every iteration is happening at any time, so you should not expect consistent behaviour of it.
As it goes Lumos will have more shape and every iteration will be more visible and it will follow a formal and pre-defined process.

![](documentation/images/logo.png#100x)


## Features

- [x] Only country Sweden some base indicators
- [x] Fixed input: **sweden** otherwise fallback quote
- [ ] Input pattern matching for multiple countries
- [ ] Rethink data model with: source, indicator description, indicator metadata...tbd...
- [ ] Data collector process to update info automatically
- [ ] Improve loader performance. Evaluate data growth.
- [ ] ...**TODO**...

## Architecture

// TODO

