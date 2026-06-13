# Go + Templ + HTMX "About Me" / CV Website

A premium, lightweight, highly interactive "About Me" and portfolio page designed to be used as an extended CV. Built with a server-side rendering stack: **Go**, **Templ** (for type-safe templates), and **HTMX** (for high-fidelity AJAX transitions and filtering) with **Vanilla CSS** styling.

## Features

- ⚡ **Zero-JS Feel**: Dynamic tab switching and form submission without full page refreshes, powered by HTMX.
- 🎨 **Minimalist Design**: Harmonious dark/light theme options, responsive timeline layouts, hover states, and glow effects.
- 💾 **Data-Driven**: Easily customize the entire website contents in a single JSON file without touching Go or HTML/CSS code.
- 🚀 **Serverless Ready**: Fully compatible with Vercel Go Serverless Functions.
- 🖼️ **CDN Profile Picture**: Natively pulls and serves profile pictures from GitHub's Camo CDN (`https://github.com/amaliebjorgen.png`) for high performance and zero data transfer bills.

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/) (1.20+)
- [Git](https://git-scm.com/)

### Running Locally

1. **Clone the repository:**
   ```bash
   git clone https://github.com/amaliebjorgen/about-me.git
   cd about-me
   ```

2. **Generate Go files from Templ templates:**
   If you make any changes to the template files (`.templ`), compile them using:
   ```bash
   go run github.com/a-h/templ/cmd/templ@latest generate
   ```

3. **Start the local development server:**
   ```bash
   go run local/main.go
   ```

4. **Open in browser:**
   Navigate to [http://localhost:8080](http://localhost:8080).

---

## How to Customize Your CV

To update the resume contents, biography, contacts, projects, or experience timeline, simply edit:

👉 **[api/models/data/cv.json](file:///home/amalie/Documents/GitHub/about-me/api/models/data/cv.json)**

Your data is embedded directly into the Go binary at build time via Go's `//go:embed` directive, meaning there is zero filesystem reading latency at runtime and perfect portability.

---

## Deployment to Vercel

This repository is pre-configured with a `vercel.json` file. 

1. Push this repository to your GitHub account (`github.com/amaliebjorgen/about-me`).
2. Go to [Vercel](https://vercel.com) and create a new project.
3. Import your GitHub repository.
4. Keep all build settings as default. Vercel automatically detects the Go handlers in `api/` and builds the serverless functions.
5. Hit **Deploy**!
