# MAS-regulated Financial Institutions

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

[![GitHub license](https://img.shields.io/badge/LICENSE-BSD--3--CLAUSE-GREEN?style=for-the-badge)](LICENSE)
[![scraper](https://img.shields.io/github/workflow/status/elliotwutingfeng/MASRegulatedFinancialInstitutions/scraper?label=SCRAPER&style=for-the-badge)](https://github.com/elliotwutingfeng/MASRegulatedFinancialInstitutions/actions/workflows/scraper.yml)
<img src="https://img.shields.io/tokei/lines/github/elliotwutingfeng/MASRegulatedFinancialInstitutions?label=Total%20Allowlist%20URLS&style=for-the-badge" alt="Total Allowlist URLs"/>

Machine-readable `.txt` allowlist of websites belonging to financial institutions regulated by the [Monetary Authority of Singapore (MAS)](https://eservices.mas.gov.sg/fid/institution), updated once a day.

The URLs in this allowlist are compiled by the **Monetary Authority of Singapore**.

**Disclaimer:** _This project is not sponsored, endorsed, or otherwise affiliated with the Monetary Authority of Singapore._

## Allowlist download

You may download the allowlist [here](mas-regulated-financial-institutions.txt?raw=1)

## Requirements

-   Go >= 1.17

## Setup instructions

`git clone` and `cd` into the project directory, then run the following

```bash
go mod tidy
```

## Usage

```bash
go run scraper.go
```

## Libraries/Frameworks used

-   [soup](https://github.com/anaskhan96/soup)
-   [tldextract](https://github.com/joeguo/tldextract)

&nbsp;

<sup>These files are provided "AS IS", without warranty of any kind, express or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose and noninfringement. In no event shall the authors or copyright holders be liable for any claim, damages or other liability, arising from, out of or in connection with the files or the use of the files.</sup>

<sub>Any and all trademarks are the property of their respective owners.</sub>
