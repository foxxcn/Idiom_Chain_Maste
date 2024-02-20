# Idiom Chain Master 🐉

[![Go version](https://img.shields.io/badge/Go-v1.22.0-blue)](https://golang.org/dl/)[![GitHub issues](https://img.shields.io/github/issues/foxxcn/Idiom_Chain_Maste)](https://github.com/foxxcn/Idiom_Chain_Maste/issues)[![GitHub forks](https://img.shields.io/github/forks/foxxcn/Idiom_Chain_Maste)](https://github.com/foxxcn/Idiom_Chain_Maste/network)[![GitHub stars](https://img.shields.io/github/stars/foxxcn/Idiom_Chain_Maste)](https://github.com/foxxcn/Idiom_Chain_Maste/stargazers)[![GitHub license](https://img.shields.io/github/license/foxxcn/Idiom_Chain_Maste)](https://github.com/foxxcn/Idiom_Chain_Maste/blob/main/LICENSE)

Welcome to the **Idiom Chain Master**! This Go-based game takes the classic word chain game to the next level, utilizing the richness of Chinese idioms. Dive into the linguistic beauty and challenge yourself with chains of idioms!

## 🎯 Objective

The game's goal is simple yet challenging: to create a chain of Chinese idioms, starting where the last idiom ended. Each player's input must start with the last character of the previous idiom, making each round a brain-teasing adventure.

## ✨ Features

- **Dictionary Integration**: Leverages an extensive dictionary of four-character Chinese idioms stored in a `.txt` file.
- **Smart Start**: Game initiation requires a valid idiom from the dictionary, ensuring every game kicks off on the right note.
- **Chain Mechanism**: The game checks the last character of the previous idiom and expects the next one to start with the same character. If direct character matching fails, it falls back to phonetic (Pinyin without tones) matching, adding a layer of complexity and learning.
- **Scoring System**: Points are awarded for successful chaining, with penalties for repetition or failure, keeping players on their toes.
- **Hints on Demand**: Stuck? Use the `tip` command to get a suggestion for your next move, though it comes with a slight score penalty.
- **Graceful Exit**: Players can end the game anytime with the `quit` command, triggering a final score calculation based on performance and efficiency.

## 🚀 How to Play

1. Clone the repository to your local machine.
2. Ensure you have Go installed and your GOPATH set up correctly.
3. Run the game using `go run .` from the project directory.
4. Follow the on-screen prompts to start chaining idioms!

## 📝 Implemented Features

- [x] Reading and parsing idioms from a `.txt` file.
- [x] Basic game loop accepting user inputs.
- [x] Scoring system based on successful chains, repetitions, and failures.
- [x] `tip` command for getting suggestions with a minor score penalty.
- [x] `quit` command for ending the game and calculating the final score.

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **External Libraries**:
  - `github.com/mozillazg/go-pinyin` for converting Chinese characters to Pinyin, aiding the phonetic fallback mechanism.

## 📦 Installation

1. Clone this repository.
2. Navigate to the cloned directory.
3. Run `go build` to compile the game.
4. Start the game with `./idiom-chain-master` (or `idiom-chain-master.exe` on Windows).

## 📈 Future Enhancements

- [ ] Improved phonetic matching to handle tones and multiple pronunciations.
- [ ] Enhanced AI for generating `tip` suggestions, potentially using machine learning to predict harder chains.
- [ ] Multiplayer support for online challenges.
- [ ] GUI for a more interactive experience.

## 🌐 A Friendly Nudge

For those intrigued by language games and programming, take a look at a similar project in Rust: [Idiom Solitaire](https://github.com/foxxcn/idiom_solitaire). While not fully fleshed out, it holds promise and awaits your contributions and curiosity.

## 📜 License

This project is open-sourced under the Apache License Version 2.0. Feel free to fork, modify, and use it in your projects. Contributions and suggestions are welcome!

## 🌟 Contributing

Your contributions are what make the **Idiom Chain Master** better. Any contributions you make are **greatly appreciated**. Check out our contribution guidelines for more information.

## 📚 Reference

This project uses the [go-pinyin](https://github.com/mozillazg/go-pinyin) library for handling Pinyin conversion, licensed under MIT.

---

Embark on a journey of words, wisdom, and wit with **Idiom Chain Master**. Let the chain begin!
