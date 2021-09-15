<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">GO Boilerplate</h3>
  <p align="center">
    GO boilerplate is an application that is developed to generate application by API specification and Database schema with the collaboration with opn-generator.  
    <br />
    <br />
    <a href="https://github.com/opn-ooo/opn-generator/issues/new">Report Bug</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#project-structure">Project Structure</a></li>
        <li><a href="#directory-explanation">Directory explanation</a></li>
        <li><a href="#app-structure">App Structure</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#requirements">Requirements</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

### Project Structure

```
📦go-boilerplate
 ┣ 📂cmd
 ┃ ┣ 📂admin
 ┃ ┣ 📂app
 ┃ ┣ 📂cli
 ┣ 📂config
 ┣ 📂database
 ┃ ┗ 📂migrations
 ┣ 📂deployments
 ┣ 📂docs
 ┣ 📂internal
 ┃ ┣ 📂http
 ┃ ┣ 📂models
 ┃ ┗ 📂repositories
 ┣ 📂pkg
 ┃ ┣ 📂database
 ┣ 📂templates
 ┃ ┣ 📂api
 ┃ ┣ 📂database
```

### Directory explanation

#### /cmd

Main applications for this project. The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).

#### /config

Configuration file templates or default configs.

#### /database

Contains database related file. (eg., /database/migrations)

#### /deployments

Contains deployment related file. (eg., aws codebuild, codepipeline, dockerfiles)

#### /docs

Contains project documents. (eg., app api docs, admin api docs, database schema etc.)

#### /internal

Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself.

#### /pkg

Library code that's ok to use by external applications (e.g., /pkg/mypubliclib). Other projects will import these libraries expecting them to work, so think twice before you put something here :-)

This tool is developed by following [Standard Go project layout](https://github.com/golang-standards/project-layout)

### App Structure

#### Repositories

Repositories are the gateway to get entities(model).

#### Services

Services are business logic layers.

#### Handlers

-   Handlers are entry points of all APIs and they should have the following things only
-   Validate request and build request objects
-   Call set of service methods
-   Build response object
-   You should not write any business logics on handler methods directly.

This project is using Uber’s [dig](https://github.com/uber-go/dig) as dependency injection toolkit.

<!-- USAGE EXAMPLES -->

## Usage

Please check [OPN generator](https://github.com/opn-ooo/opn-generator) readme section check the usage process of go boilerplate.

<!-- ROADMAP -->

## Roadmap

The roadmap is useful for planning large pieces of work several months in advance at the Epic level within a single project. Simple planning and dependency management features help teams visualize and manage work better together.

See the [open issues](https://github.com/opn-ooo/go-boilerplate/issues?q=is%3Aopen+is%3Aissue) for a list of proposed features (and known issues).

<!-- REQUIREMENTS -->

## Requirements

-   Go^1.16

<!-- ACKNOWLEDGEMENTS -->

## Acknowledgements

-   [Gin](https://gin-gonic.com)
-   [Dig](https://pkg.go.dev/go.uber.org/dig)
-   [Gorm](https://gorm.io)
    And many more
