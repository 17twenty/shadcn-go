# Shadcn-go

Shadcn-go is a Go project that combines the power of Go programming language with Templ templating engine and the sleek UI components from shadcn/ui.

## Features

- Built with Go: Leverage the performance and simplicity of Go for backend development
- Templ Integration: Use Templ for efficient and type-safe HTML templating in Go
- shadcn/ui Components: Incorporate beautiful, accessible, and customizable UI components
- Modern Web Development: Combine server-side rendering with interactive client-side features

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   $ go mod tidy
   $ go install github.com/a-h/templ/cmd/templ@latest
   $ go install github.com/air-verse/air@latest
   ```

   Grab the static tailwind binary from [here](https://github.com/tailwindlabs/tailwindcss/releases) and store it in your path - i.e:
   ```bash
   $ # Note which arch you're using, this is for an older version and arm64
   $ mkdir -p ~/tailwind
   $ cd ~/tailwind
   $ curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.6/tailwindcss-macos-arm64
   $ chmod +x tailwindcss-macos-arm64
   $ mv tailwindcss-macos-arm64 tailwindcss
   $ PATH=$PATH:`pwd`/tools
   ...
   ```

3. Run the application:
   ```bash
   $ task run # or live etc. Check Taskfile.yml for options
   ```

## Project Structure

- `main.go`: Entry point of the application
- `shadcn/`: Reusable UI components built with Templ and shadcn/ui
- `static/`: Static assets (CSS, JavaScript, images)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
