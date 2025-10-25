# Simple Concurrent Web Crawler in Go 🕷️

This is a basic, concurrent web crawler written in **Go** that scrapes all unique, absolute **HTTPS URLs** from a set of initial seed URLs. It leverages Go's concurrency primitives like **goroutines** and **channels** to crawl multiple pages simultaneously.

---

## 💡 How It Works

1.  **Concurrency:** Each initial seed URL is crawled in a separate **goroutine**.
2.  **Communication:**
    * The `channelUrl` is used by the crawler goroutines to send found links back to the `main` function.
    * The `channelFinished` is used by the crawler goroutines to signal the `main` function when they have completed their job.
3.  **Parsing:** The `golang.org/x/net/html` package is used to tokenize the HTML content and extract the `href` attribute from anchor (`<a>`) tags.
4.  **Filtering:** Only URLs that begin with `"https"` (absolute URLs with the HTTPS protocol) are collected.
5.  **Uniqueness:** The `main` function stores all found URLs in a map (`foundUrl`) to ensure only unique links are displayed at the end.

---

## 🛠️ Prerequisites

* Go (version 1.18 or higher recommended)

---

## 🚀 Installation and Usage

### 1. Clone the repository

```bash
git clone web_scraper
cd web_scraper
