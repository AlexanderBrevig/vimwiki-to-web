<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/alexanderbrevig/vimwiki-to-web">
    <img src="assets/logo.png" alt="Logo" width="256" height="265">
  </a>

  <h3 align="center">vimwiki-to-web</h3>

  <p align="center">
    Publish your vimwiki go Gemini!
    <br />
    <blockquote>
        <p>Write what you know. That should leave you with a lot of free time.</p>
        <footer>Howard Nemerov</footer> 
    </blockquote>
    <br />
    <a href="https://github.com/alexanderbrevig/vimwiki-to-web"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    ·
    <a href="https://github.com/alexanderbrevig/vimwiki-to-web/issues">Report Features & Bug</a>
  </p>
</p>


<!-- ABOUT THE PROJECT -->
## About The Project

This project allows you to export your vimwiki (written in Markdown) to Gemini `.gmi` files for sharing on `gemini://`.

You can visit my site for an example (bare in mind, I just started my vimwiki):

gemini://gemini.alexanderbrevig.com


### Built With

* [Go](https://golang.org/)

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

* golang

### Installation and Use

1. Clone the repo
   ```sh
   git clone https://github.com/alexanderbrevig/vimwiki-to-web.git
   ```
2. Run the program
   ```sh
   # This will render ~/vimwiki to ~/gemini, recursively
   go run ./cmd/vimwiki-to-web -vimwiki ~/vimwiki -file / -gmiout ~/gemini
   ```
3. Install if you want
   ```sh
   GOBIN=/usr/local/bin/ go install ./cmd/vwtw
   ```

<!-- USAGE EXAMPLES -->
#### Usage

   ```
   > vwtw -h[elp]
   Usage of vwtw:
        -file string
            Path to your file (default "/index.md")
        -gmiout string
            Path to your output folder (default "/home/$(whoami)/vimwiki-gemini")
        -vimwiki string
            Path to your vimwiki (default "/home/$(whoami)/vimwiki")
   ```


<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/alexanderbrevig/vimwiki-to-web/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contribuions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Alexander Brevig - [@alexanderbrevig](https://twitter.com/alexanderbrevig) - alexanderbrevig@gmail.com

Project Link: [https://github.com/alexanderbrevig/vimwiki-to-web](https://github.com/alexanderbrevig/vimwiki-to-web)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [Amazing Go Markdown Library](github.com/gomarkdown/markdown)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/alexanderbrevig/vimwiki-to-web.svg?style=for-the-badge
[contributors-url]: https://github.com/alexanderbrevig/vimwiki-to-web/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/alexanderbrevig/vimwiki-to-web.svg?style=for-the-badge
[forks-url]: https://github.com/alexanderbrevig/vimwiki-to-web/network/members
[stars-shield]: https://img.shields.io/github/stars/alexanderbrevig/vimwiki-to-web.svg?style=for-the-badge
[stars-url]: https://github.com/alexanderbrevig/vimwiki-to-web/stargazers
[issues-shield]: https://img.shields.io/github/issues/alexanderbrevig/vimwiki-to-web.svg?style=for-the-badge
[issues-url]: https://github.com/alexanderbrevig/vimwiki-to-web/issues
[license-shield]: https://img.shields.io/github/license/alexanderbrevig/vimwiki-to-web.svg?style=for-the-badge
[license-url]: https://github.com/alexanderbrevig/vimwiki-to-web/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/alexanderbrevig

